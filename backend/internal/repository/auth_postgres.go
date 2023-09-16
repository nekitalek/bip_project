package repository

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	BIP_project "github.com/nekitalek/bip_project/backend"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user BIP_project.User_auth) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, username, password_hash,email_confirmation) values ($1, $2, $3,$4) RETURNING user_id", usersTable)

	row := r.db.QueryRow(query, user.Login, user.Username, user.Password_hash, user.Email_confirmation)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (r *AuthPostgres) UpdateDataUser(user BIP_project.User_auth) error {

	query := fmt.Sprintf("UPDATE %s SET login = $2,password_hash = $3, username = $4, email_confirmation = $5, WHERE user_id=$1", usersTable)

	_ = r.db.QueryRow(query, user.User_id, user.Login, user.Password, user.Email_confirmation)

	return nil

}
func (r *AuthPostgres) UpdateUsersEmailConfirmation(user_id int, email_confirmation bool) error {
	query := fmt.Sprintf("UPDATE %s SET email_confirmation = $2 WHERE user_id=$1", usersTable)

	_ = r.db.QueryRow(query, user_id, email_confirmation)
	return nil
}

func (r *AuthPostgres) GetUserByLogin(login string) (BIP_project.User_auth, error) {
	var user BIP_project.User_auth

	query := fmt.Sprintf("SELECT * FROM %s WHERE login=$1", usersTable)
	err := r.db.Get(&user, query, login)
	return user, err
}
func (r *AuthPostgres) GetUserById(id int) (BIP_project.User_auth, error) {
	var user BIP_project.User_auth

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", usersTable)
	err := r.db.Get(&user, query, id)

	return user, err
}

func (r *AuthPostgres) AddDataAuth(data BIP_project.Auth_data) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (token,user_id,time_end,device) values ($1, $2, $3, $4) RETURNING auth_data_id", authDataTable)

	row := r.db.QueryRow(query, data.Token, data.User_id, data.Time_end, data.Device)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil

}

func (r *AuthPostgres) GetDataAuth(id int) (BIP_project.Auth_data, error) {
	var data BIP_project.Auth_data

	query := fmt.Sprintf("SELECT * FROM %s WHERE auth_data_id=$1", authDataTable)
	err := r.db.Get(&data, query, id)

	return data, err
}

func (r *AuthPostgres) UpdateDataAuth(data BIP_project.Auth_data) error {

	query := fmt.Sprintf("UPDATE %s SET token = $1,user_id = $2, time_end = $3, device = $4  WHERE auth_data_id=$1", authDataTable)

	_ = r.db.QueryRow(query, data.Auth_data_id, data.Token, data.User_id, data.Time_end, data.Device)

	return nil
}

func (r *AuthPostgres) CreateDataEmailConf(email_conf BIP_project.Email_confirmation) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, email,code_email_confirmation, time_end,assignment,device) values ($1, $2, $3,$4,$5,$6) RETURNING email_confirmation_id", emailConfirmationTable)

	row := r.db.QueryRow(query, email_conf.User_id, email_conf.Email, email_conf.Code_email_confirmation, email_conf.Time_end, email_conf.Assignment, email_conf.Device)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetEmailConfByUserId(user_id int) (BIP_project.Email_confirmation, error) {
	var e_conf BIP_project.Email_confirmation
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", emailConfirmationTable)
	err := r.db.Get(&e_conf, query, user_id)
	return e_conf, err
}

func (r *AuthPostgres) DeleteEmailConfByUserId(user_id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1", emailConfirmationTable)
	_ = r.db.QueryRow(query, user_id)
	return nil
}

func (r *AuthPostgres) UpdateCodeEmailConf(e_conf_id, new_code int) error {
	query := fmt.Sprintf("UPDATE %s SET code_email_confirmation = $2 WHERE email_confirmation_id=$1", emailConfirmationTable)

	row := r.db.QueryRow(query, e_conf_id, new_code)
	fmt.Println("UpdateCodeEmailConf, ", row)
	return nil
}

func (r *AuthPostgres) UpdatePass(login, password, new_password string) error {
	query := fmt.Sprintf("UPDATE %s SET password_hash = $3 WHERE login=$1 AND password_hash = $2", usersTable)

	r.db.QueryRow(query, login, password, new_password)

	return nil
}
func (r *AuthPostgres) UpdateLogin(user_id int, new_login string) error {
	fmt.Println(user_id, new_login)
	query := fmt.Sprintf("UPDATE %s SET login = $2 WHERE user_id = $1", usersTable)

	row := r.db.QueryRow(query, user_id, new_login)
	fmt.Println("UpdateLogin, ", row)

	return nil
}

func (r *AuthPostgres) DeleteJwtTokens(user_id int) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1", authDataTable)
	_ = r.db.QueryRow(query, user_id)
	return nil
}

func (r *AuthPostgres) GetDataLoginAttempt(login string) (BIP_project.Login_attempt, error) {
	var log_attempt BIP_project.Login_attempt

	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1", loginAttemptTable)
	err := r.db.Get(&log_attempt, query, login)
	return log_attempt, err
}

func (r *AuthPostgres) CreateDataLoginAttempt(log_attempt BIP_project.Login_attempt) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email,number_of_inputs,unlock_time,login_method) values ($1, $2, $3,$4) RETURNING login_attempt_id", loginAttemptTable)

	row := r.db.QueryRow(query, log_attempt.Email, log_attempt.Number_of_inputs, log_attempt.Unlock_time, log_attempt.Login_method)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) UpdateLoginAttempt(log_attempt BIP_project.Login_attempt) error {
	query := fmt.Sprintf("UPDATE %s SET email = $2,number_of_inputs = $3, unlock_time = $4,login_method = $5  WHERE login_attempt_id=$1", loginAttemptTable)
	_ = r.db.QueryRow(query, log_attempt.Login_attempt_id, log_attempt.Email, log_attempt.Number_of_inputs, log_attempt.Unlock_time, log_attempt.Login_method)
	return nil
}

func (r *AuthPostgres) DeleteLoginAttempt(login string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE email = $1", loginAttemptTable)
	_ = r.db.QueryRow(query, login)
	return nil
}

func (r *AuthPostgres) CheckJWTBlacklist(user_id int, token_valid_from int64) (bool, error) {
	var cnt int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE (token_valid_from>$1 AND user_id=$2)", jwtBlacklistTable)
	err := r.db.Get(&cnt, query, time.Unix(token_valid_from, 0), user_id)

	if err != nil {
		return false, err
	}
	if cnt == 0 {
		return true, nil
	}
	return false, nil
}

func (r *AuthPostgres) CreateJWTBlacklist(user_id int, token_valid_from, cleanup_time time.Time) error {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id,token_valid_from, cleanup_time) values ($1, $2, $3) RETURNING id", jwtBlacklistTable)

	row := r.db.QueryRow(query, user_id, token_valid_from, cleanup_time)
	if err := row.Scan(&id); err != nil {
		return err
	}

	return nil
}
