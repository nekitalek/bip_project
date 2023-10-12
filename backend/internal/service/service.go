package service

import (
	BIP_project "github.com/nekitalek/bip_project/backend"
	"github.com/nekitalek/bip_project/backend/internal/repository"
)

type Authorization interface {
	CreateUser(user BIP_project.User_auth) (int, int, error)
	SignUpSecondFactor(e_conf BIP_project.Email_confirmation) error

	// SendVerificationCodeToEmail(user_id int, login string, assignment BIP_project.Assignment) (int, error)
	// SendVerificationCodeToEmailByLogin(login string, assignment BIP_project.Assignment) (int, error)
	// SendVerificationCodeToEmailByUserId(user_id int, assignment BIP_project.Assignment) (int, error)
	CheckLoginAttempt(login string, login_method BIP_project.Login_method) error
	Authenticate(login, password string) (BIP_project.User_auth, error)
	Send2FAEmail(user_id int, login string, assignment BIP_project.Assignment) (int, int, error)

	AuthenticateAndSend2FAEmail(login, password string, assignment BIP_project.Assignment) (int, int, error)
	CheckEmailVerificationCode(e_conf BIP_project.Email_confirmation, assignment BIP_project.Assignment) error

	// ReSendVerificationCodeToEmail(e_conf_id int, login string) error
	// ReSendVerificationCodeToEmailByUserId(e_conf_id, user_id int) error
	ReSendCode(e_conf BIP_project.Email_confirmation) error

	SingInByPass(login, password string) (int, int, error)
	SignInSecondFactor(e_conf BIP_project.Email_confirmation) (string, error)

	GenerateTokenJWT(user_id, code int) (string, error)
	CheckPass(login, password string) error

	SaveToken(token_JWT string, e_conf BIP_project.Email_confirmation) error

	CreateEmailConfirmation(user_id int, login string, assignment BIP_project.Assignment) (int, int, error)

	ChangePassFirstFactor(login, password string) (int, int, error)
	ChangePassSecondFactor(e_conf BIP_project.Email_confirmation, new_password string) error

	ChangeLoginFirstFactor(login, password string) (int, int, error)
	ChangeLoginSecondFactor(e_conf BIP_project.Email_confirmation, new_login string) (int, int, error)
	VerificationNewEmail(e_conf BIP_project.Email_confirmation) error

	ParseToken(accessToken string) (int, error)
}

type EventItem interface {
	CreateEvent(user_id int, item *BIP_project.Event_items) (int, error)
	GetEvents(input *BIP_project.Event_items_input) ([]BIP_project.Event_items, error)
	UpdateEvent(userId, itemId int, input *BIP_project.Event_items_input) error
	DeleteEvent(userId, itemId int) error
}
type Invitation interface {
	CreateInvitation(user_id int, input *BIP_project.Event_invitations_input) (int, error)
	GetInvitation(user_id int) ([]BIP_project.Event_invitations, error)
	UpdateInvitation(user_id, event_invitations_id int, input *BIP_project.Event_invitations_input) error
	DeleteInvitation(user_id int, input *BIP_project.Event_invitations_input) error
}

type User interface {
	GetUser(user_id int) (BIP_project.User_data, error)
}
type PushNotification interface {
	CreatePushNotification(user_id int, token, device string) error
	DeletePushNotification(user_id int, token string) error
	SendPushNotification(event_id int) error
}

type Service struct {
	Authorization
	Invitation
	EventItem
	User
	PushNotification
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:    NewAuthService(repos.Authorization),
		Invitation:       NewInvitationService(repos.Invitation),
		EventItem:        NewEventItemService(repos.EventItem),
		User:             NewUserService(repos.Authorization),
		PushNotification: NewPushNotificationService(repos.PushNotification),
	}
}
