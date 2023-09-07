package repository

import (
	"github.com/jmoiron/sqlx"
	BIP_project "github.com/nekitalek/bip_project/backend"
)

type Authorization interface {
	CreateUser(user BIP_project.User_auth) (int, error)
	UpdateDataUser(user BIP_project.User_auth) error
	GetUserByLogin(login string) (BIP_project.User_auth, error)
	GetUserById(id int) (BIP_project.User_auth, error)
	AddDataAuth(data BIP_project.Auth_data) (int, error)
	GetDataAuth(id int) (BIP_project.Auth_data, error)
	UpdateDataAuth(data BIP_project.Auth_data) error
	CreateDataEmailConf(email_conf BIP_project.Email_confirmation) (int, error)
	GetEmailConfByUserId(user_id int) (BIP_project.Email_confirmation, error)
	DeleteEmailConfByUserId(user_id int) error
	UpdateUsersEmailConfirmation(user_id int, email_confirmation bool) error
	UpdateCodeEmailConf(e_conf_id, new_code int) error

	UpdatePass(login, password, new_password string) error
	UpdateLogin(user_id int, new_login string) error

	DeleteJwtTokens(user_id int) error

	CreateDataLoginAttempt(log_attempt BIP_project.Login_attempt) (int, error)
	GetDataLoginAttempt(login string) (BIP_project.Login_attempt, error)
	UpdateLoginAttempt(log_attempt BIP_project.Login_attempt) error
	DeleteLoginAttempt(login string) error
}

type EventItem interface {
	CreateEvent(item *BIP_project.Event_items) (int, error)
	GetEvents(input *BIP_project.Event_items_input) ([]BIP_project.Event_items, error)
	UpdateEvent(userId, itemId int, input *BIP_project.Event_items_input) error
	DeleteEvent(userId, eventId int) error
}
type Invitation interface {
	CreateInvitation(user_id int, input *BIP_project.Event_invitations_input) (int, error)
	GetInvitation(user_id int) ([]BIP_project.Event_invitations, error)
	UpdateInvitation(user_id, event_invitations_id int, input *BIP_project.Event_invitations_input) error
	DeleteInvitation(user_id int, input *BIP_project.Event_invitations_input) error
}

type Repository struct {
	Authorization
	Invitation
	EventItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Invitation:    NewInvitationPostgres(db),
		EventItem:     NewEventItemPostgres(db),
	}
}
