package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	BIP_project "github.com/nekitalek/bip_project/backend"
)

type PushNotificationPostgres struct {
	db *sqlx.DB
}

func NewPushNotificationPostgres(db *sqlx.DB) *PushNotificationPostgres {
	return &PushNotificationPostgres{db: db}
}

func (r *PushNotificationPostgres) CreatePushNotification(user_id int, token,device string ) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, push_token, device) values ($1, $2, $3) RETURNING user_id", pushNotificationTable)

	row := r.db.QueryRow(query, user_id, token, device)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PushNotificationPostgres) DeletePushNotification(user_id int, token string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE (user_id = $1 AND push_token = $2)", pushNotificationTable)
	_ = r.db.QueryRow(query, user_id, token)
	return nil
}

func (r *PushNotificationPostgres) GetPushNotification(event_id int, status BIP_project.Status) ([]string, error) {
	tokens := make([]string, 0)

	query := fmt.Sprintf(`
	SELECT pn.push_token
	FROM %s pn
	JOIN %s ei ON pn.user_id = ei.user_id
	WHERE ei.event_id = $1 AND ei.status = $2;`, pushNotificationTable, eventInvitationsTable)

	err := r.db.Select(&tokens, query, event_id, status)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
