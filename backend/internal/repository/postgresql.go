package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	usersTable             = "users"
	loginAttemptTable      = "login_attempt"
	authDataTable          = "auth_data"
	emailConfirmationTable = "email_confirmation"
	eventItemsTable        = "event_items"
	eventInvitationsTable  = "event_invitations"
	jwtBlacklistTable      = "JWT_blacklist"
	pushNotificationTable  = "push_notification"
)

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
