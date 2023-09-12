package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	BIP_project "github.com/nekitalek/bip_project/backend"
)

type InvitationPostgres struct {
	db *sqlx.DB
}

func NewInvitationPostgres(db *sqlx.DB) *InvitationPostgres {
	return &InvitationPostgres{db: db}
}

func (r *InvitationPostgres) CreateInvitation(user_id int, input *BIP_project.Event_invitations_input) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (event_id, user_id, status) SELECT $1, $2, $3 FROM %s e WHERE e.event_items_id = $1 RETURNING event_invitations_id", eventInvitationsTable, eventItemsTable)

	row := r.db.QueryRow(query, input.Event_id, input.User_id, input.Status)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (r *InvitationPostgres) GetInvitation(user_id int) ([]BIP_project.Event_invitations, error) {

	var invitations []BIP_project.Event_invitations

	query := fmt.Sprintf(`
	SELECT ei.*
	FROM %s AS ei
	LEFT JOIN %s AS ei_items ON ei.event_id = ei_items.event_items_id
	WHERE ei.user_id = $1 OR ei_items.admin = $1`, eventInvitationsTable, eventItemsTable)
	err := r.db.Select(&invitations, query, user_id)
	return invitations, err
}
func (r *InvitationPostgres) UpdateInvitation(user_id, event_invitations_id int, input *BIP_project.Event_invitations_input) error {

	query := fmt.Sprintf(`
	UPDATE %s
	SET status = $3
	WHERE event_invitations_id = $1 AND user_id = $2`, eventInvitationsTable)

	_, err := r.db.Exec(query, event_invitations_id, user_id, input.Status)

	return err
}
func (r *InvitationPostgres) DeleteInvitation(user_id int, input *BIP_project.Event_invitations_input) error {

	query := fmt.Sprintf(`
	DELETE FROM %s ei
	USING %s e_items
	WHERE (ei.user_id = $1 AND ei.event_id = $2) AND
		(ei.user_id = $3
   		OR 
		(e_items.admin = $3 AND ei.event_id = $2));`, eventInvitationsTable, eventItemsTable)
	res, err := r.db.Exec(query, input.User_id, input.Event_id, user_id)
	if num, _ := res.RowsAffected(); num == 0 {
		return errors.New("error when deleting to db")
	}
	return err
}
