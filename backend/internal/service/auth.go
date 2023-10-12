package service

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	BIP_project "github.com/nekitalek/bip_project/backend"
	"github.com/nekitalek/bip_project/backend/internal/repository"
	"github.com/spf13/viper"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

const (
	tokenTTL = 12 * time.Hour
)

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GenerateTokenJWT(user_id, code int) (string, error) {
	user, err := s.repo.GetUserById(user_id)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.User_id,
	})
	return token.SignedString([]byte(os.Getenv("JwT_SINGING_KEY")))

}

func (s *AuthService) CheckPass(login, password string) error {
	_, err := s.repo.GetUserByLogin(login)
	return err
}

func (s *AuthService) SingInByPass(login, password string) (int, int, error) {
	return s.AuthenticateAndSend2FAEmail(login, password, BIP_project.SecFactor)
}
func (s *AuthService) SignInSecondFactor(e_conf BIP_project.Email_confirmation) (string, error) {
	err := s.CheckEmailVerificationCode(e_conf, BIP_project.SecFactor)
	if err != nil {
		return "", err
	}
	token_JWT, err := s.GenerateTokenJWT(e_conf.User_id, e_conf.Code_email_confirmation)
	if err != nil {
		return "", err
	}
	err = s.SaveToken(token_JWT, e_conf)
	if err != nil {
		return "", err
	}

	return token_JWT, nil
}
func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("JwT_SINGING_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	// проверить в черном списке токен
	res, err := s.repo.CheckJWTBlacklist(claims.UserId, claims.IssuedAt)
	if err != nil {
		return 0, err
	}

	if !res {
		return 0, errors.New("token has been revoked")
	}
	return claims.UserId, nil
}

func (s *AuthService) SaveToken(token_JWT string, e_conf BIP_project.Email_confirmation) error {
	token := BIP_project.Auth_data{
		Token:    token_JWT,
		User_id:  e_conf.User_id,
		Time_end: time.Now().Add(tokenTTL),
		Device:   e_conf.Device,
	}
	//соханяем в бд
	_, err := s.repo.AddDataAuth(token)

	return err
}

func (s *AuthService) CreateUser(user BIP_project.User_auth) (int, int, error) {
	//если в бд уже есть почта то вернёт ошибку
	user.Password_hash = generatePasswordHash(user.Password)
	user.Email_confirmation = false
	user_id, err := s.repo.CreateUser(user)
	if err != nil {
		return 0, 0, err
	}

	//создаем подтвердение
	id, code, err := s.CreateEmailConfirmation(user_id, user.Login, BIP_project.Registration)
	if err != nil {
		return 0, 0, err
	}
	//отправляем подтверждение
	err = sendEmailWithCode(user.Login, viper.GetString("template.EmailConfirmation"), code)
	if err != nil {
		return 0, 0, err
	}
	if err != nil {
		return 0, 0, err
	}

	return user_id, id, nil
}

func (s *AuthService) SignUpSecondFactor(e_conf BIP_project.Email_confirmation) error {
	err := s.CheckEmailVerificationCode(e_conf, BIP_project.Registration)
	if err != nil {
		return err
	}
	return s.repo.UpdateUsersEmailConfirmation(e_conf.User_id, true)
}

func (s *AuthService) CheckLoginAttempt(login string, login_method BIP_project.Login_method) error {
	//попытка получить данные
	var data BIP_project.Login_attempt
	data, err := s.repo.GetDataLoginAttempt(login)

	//если данныех нет то создать
	if err != nil {
		//создаем запись
		data.Email = login
		data.Number_of_inputs = 1
		data.Unlock_time = time.Now().In(time.UTC)
		data.Login_method = login_method
		_, err = s.repo.CreateDataLoginAttempt(data)
		return err
	}

	//проверить данные
	if login_method != data.Login_method {
		return errors.New("wrong login_method")
	}

	if data.Number_of_inputs == 12 {
		return errors.New("entry limit exceeded")
	}
	if data.Unlock_time.After(time.Now().In(time.UTC)) {
		return errors.New(fmt.Sprintf("wait %v for the next move attempt", data.Unlock_time.Sub(time.Now().In(time.UTC)).Round(time.Second)))
	}

	data.Number_of_inputs += 1

	if data.Number_of_inputs%3 == 0 {
		data.Unlock_time = time.Now().In(time.UTC).Add(exponentialBackoff(data.Number_of_inputs))
	}
	err = s.repo.UpdateLoginAttempt(data)
	return err
}

func (s *AuthService) Authenticate(login, password string) (BIP_project.User_auth, error) {

	var user BIP_project.User_auth
	//проверить данные о попытках входа
	err := s.CheckLoginAttempt(login, BIP_project.Password)
	if err != nil {
		return user, err
	}
	//проверим login и pass
	user, err = s.repo.GetUserByLogin(login)
	if err != nil {
		return user, err
	}
	if user.Password_hash != generatePasswordHash(password) {
		return user, errors.New("wrong login or password")
	}

	//удаить данные о попытках входа
	s.repo.DeleteLoginAttempt(login)

	//если не подтверждена почта то запрещается вход
	if !user.Email_confirmation {
		return user, errors.New("mail not confirmed")
	}
	return user, nil
}
func (s *AuthService) Send2FAEmail(user_id int, login string, assignment BIP_project.Assignment) (int, int, error) {

	//создаем подтвердение
	id_e_conf, code, err := s.CreateEmailConfirmation(user_id, login, assignment)
	if err != nil {
		return 0, 0, err
	}
	//отправляем подтверждение
	err = sendEmailWithCode(login, viper.GetString("template.EmailConfirmation"), code)
	if err != nil {
		s.repo.DeleteEmailConfByUserId(user_id)
		return 0, 0, err
	}
	return user_id, id_e_conf, nil
}

func (s *AuthService) AuthenticateAndSend2FAEmail(login, password string, assignment BIP_project.Assignment) (int, int, error) {
	//тут костыль ибо в Authenticate и Send2FAEmail вызвают GetUserByLogin(login) можно вызывать всего 1 раз
	user, err := s.Authenticate(login, password)
	if err != nil {
		return 0, 0, err
	}

	return s.Send2FAEmail(user.User_id, login, assignment)
}

func (s *AuthService) CheckEmailVerificationCode(e_conf BIP_project.Email_confirmation, assignment BIP_project.Assignment) error {
	//получаяем по user_id данные Email_confirmation
	old_e_conf, err := s.repo.GetEmailConfByUserId(e_conf.User_id)
	if err != nil {
		return err
	}

	//проверить данные о попытках входа
	err = s.CheckLoginAttempt(old_e_conf.Email, BIP_project.Code2fa)
	if err != nil {
		return err
	}

	//сверяем assignment
	if assignment != old_e_conf.Assignment {
		return errors.New("invalid assignment")
	}
	//сверяем код
	if e_conf.Code_email_confirmation != old_e_conf.Code_email_confirmation {
		return errors.New("invalid Code")
	}

	//удаить данные о попытках входа
	s.repo.DeleteLoginAttempt(old_e_conf.Email)

	//удаяем данные о Verification Code
	err = s.repo.DeleteEmailConfByUserId(e_conf.User_id)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) ReSendCode(e_conf BIP_project.Email_confirmation) error {
	old_e_conf, err := s.repo.GetEmailConfByUserId(e_conf.User_id)
	if err != nil {
		return err
	}
	//проверим что у юзера было открыто подтверждение по почте
	if old_e_conf.User_id != e_conf.User_id {
		return errors.New("there is no Email_confirmation_id for this user_id")
	}

	//генерируем новый код
	new_code := generateCode()

	//обновяем данные в бд
	err = s.repo.UpdateCodeEmailConf(e_conf.Email_confirmation_id, new_code)
	if err != nil {
		return err
	}
	//отравяем письмо
	user, err := s.repo.GetUserById(e_conf.User_id)
	if err != nil {
		return err
	}
	err = sendEmailWithCode(user.Login, viper.GetString("template.EmailConfirmation"), new_code)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) CreateEmailConfirmation(user_id int, login string, assignment BIP_project.Assignment) (int, int, error) {
	code := generateCode()
	//code := 123456
	var email_confirmation BIP_project.Email_confirmation
	email_confirmation.User_id = user_id
	email_confirmation.Email = login
	email_confirmation.Code_email_confirmation = code
	email_confirmation.Time_end = time.Now().Add(emailConfirmationTTL)
	email_confirmation.Assignment = assignment
	email_confirmation.Device = "windows"

	//сохраняем данные в бд
	id, err := s.repo.CreateDataEmailConf(email_confirmation)
	if err != nil {
		return 0, 0, err
	}

	return id, code, nil
}

func (s *AuthService) ChangePassFirstFactor(login, password string) (int, int, error) {
	return s.AuthenticateAndSend2FAEmail(login, password, BIP_project.ChangePassword)
}

func (s *AuthService) ChangePassSecondFactor(e_conf BIP_project.Email_confirmation, new_password string) error {
	//проверяем 2fa
	err := s.CheckEmailVerificationCode(e_conf, BIP_project.ChangePassword)
	if err != nil {
		return err
	}
	//меняем пароль на новый
	user, err := s.repo.GetUserById(e_conf.User_id)
	if err != nil {
		return err
	}
	err = s.repo.UpdatePass(user.Login, user.Password_hash, generatePasswordHash(new_password))
	if err != nil {
		return err
	}
	//отзываем токены
	//err = s.repo.DeleteJwtTokens(db_e_conf.User_id)
	err = s.repo.CreateJWTBlacklist(e_conf.User_id, time.Now(), time.Now().Add(tokenTTL))
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) ChangeLoginFirstFactor(login, password string) (int, int, error) {
	return s.AuthenticateAndSend2FAEmail(login, password, BIP_project.ChangeEmail)
}

func (s *AuthService) ChangeLoginSecondFactor(e_conf BIP_project.Email_confirmation, new_login string) (int, int, error) {
	//проверяем 2fa
	err := s.CheckEmailVerificationCode(e_conf, BIP_project.ChangeEmail)
	if err != nil {
		return 0, 0, err
	}
	//проверим что такая почта уже не занята
	_, err = s.repo.GetUserByLogin(new_login)
	if err == nil {
		return 0, 0, errors.New("Email already registered")
	}
	//создаем запрос для верификации
	return s.Send2FAEmail(e_conf.User_id, new_login, BIP_project.ChangeEmail)
}

func (s *AuthService) VerificationNewEmail(e_conf BIP_project.Email_confirmation) error {
	//получим e_conf из бд что бы взять новую почту, так как после вызова CheckEmailVerificationCode запись удалится, если данные верны
	db_e_conf, err := s.repo.GetEmailConfByUserId(e_conf.User_id)
	if err != nil {
		return err
	}
	//проверяем 2fa
	err = s.CheckEmailVerificationCode(e_conf, BIP_project.ChangeEmail)
	if err != nil {
		return err
	}
	err = s.repo.UpdateLogin(db_e_conf.User_id, db_e_conf.Email)
	if err != nil {
		return err
	}
	//отзываем токены
	// err = s.repo.DeleteJwtTokens(db_e_conf.User_id)
	s.repo.CreateJWTBlacklist(e_conf.User_id, time.Now(), time.Now().Add(tokenTTL))
	if err != nil {
		return err
	}
	return nil
}
