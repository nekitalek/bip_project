package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	BIP_project "github.com/nekitalek/bip_project/backend"
)

type EventItemPostgres struct {
	db *sqlx.DB
}

func NewEventItemPostgres(db *sqlx.DB) *EventItemPostgres {
	return &EventItemPostgres{db: db}
}

func (r *EventItemPostgres) CreateEvent(item *BIP_project.Event_items) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (admin, time_start,time_end,place,game,description,public) values ($1, $2, $3, $4, $5, $6, $7) RETURNING event_items_id", eventItemsTable)

	row := tx.QueryRow(createItemQuery, item.Admin, item.Time_start, item.Time_end, item.Place, item.Game, item.Description, item.Public)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createEventInvitationsQuery := fmt.Sprintf("INSERT INTO %s (event_id, user_id,status) values ($1, $2, $3)", eventInvitationsTable)
	_, err = tx.Exec(createEventInvitationsQuery, itemId, item.Admin, BIP_project.Confirmed)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

type Event_items_test struct {
	Event_items_id int            `json:"event_items_id"`
	Admin          int            `json:"admin"`
	Participant    sql.NullString `json:"Participant"`
	Time_start     time.Time      `json:"time_start"`
	Time_end       time.Time      `json:"time_end"`
	Place          string         `json:"place"`
	Game           string         `json:"game"`
	Description    string         `json:"description"`
	Public         bool           `json:"public"`
}

func (r *EventItemPostgres) GetEvents(input *BIP_project.Event_items_input) ([]BIP_project.Event_items, error) {
	var events []BIP_project.Event_items
	var events_test []Event_items_test

	setValues := make([]string, 0)

	if input.Event_items_id != nil {
		setValues = append(setValues, fmt.Sprintf("e.event_items_id=%d", *input.Event_items_id))
	}

	if input.Admin != nil {
		setValues = append(setValues, fmt.Sprintf("e.admin = %d", *input.Admin))
	}

	if input.Time_start != nil && input.Time_end != nil {
		setValues = append(setValues, fmt.Sprintf("(e.time_start BETWEEN DATE '%s' AND DATE '%s')", input.Time_start.Format(time.RFC3339), input.Time_end.Format(time.RFC3339)))
	}

	if input.Game != nil {
		setValues = append(setValues, fmt.Sprintf("e.game = '%s'", *input.Game))
	}

	if input.Place != nil {
		setValues = append(setValues, fmt.Sprintf("e.place = '%s'", *input.Place))
	}
	if input.Public != nil {
		setValues = append(setValues, fmt.Sprintf("e.public = %t", *input.Public))
	}

	var setQuery string
	if len(setValues) != 0 {
		setQuery = "WHERE " + strings.Join(setValues, " AND ")
	}

	query := fmt.Sprintf(`SELECT
		e.event_items_id,
		e.admin,
		e.time_start,
		e.time_end,
		e.place,
		e.game,
		e.description,
		e.public,
		json_agg(json_build_object('user_id', ei.user_id, 'username', u.username)) AS participant
	FROM
		%s e
	LEFT JOIN
		%s ei ON e.event_items_id = ei.event_id
	LEFT JOIN
		users u ON ei.user_id = u.user_id
		%s
	GROUP BY
		e.event_items_id, e.admin, e.time_start, e.time_end, e.place, e.game, e.description, e.public`, eventItemsTable, eventInvitationsTable, setQuery)

	err := r.db.Select(&events_test, query)
	if err != nil {
		return events, err
	}
	for i := range events_test {
		new_item := BIP_project.Event_items{
			Event_items_id: events_test[i].Event_items_id,
			Admin:          events_test[i].Admin,
			Time_start:     events_test[i].Time_start,
			Time_end:       events_test[i].Time_end,
			Place:          events_test[i].Place,
			Game:           events_test[i].Game,
			Description:    events_test[i].Description,
			Public:         events_test[i].Public,
		}
		json.Unmarshal([]byte(events_test[i].Participant.String), &new_item.Participant)

		events = append(events, new_item)
	}

	return events, err
}

func (r *EventItemPostgres) UpdateEvent(userId, eventId int, input *BIP_project.Event_items_input) error {

	setValues := make([]string, 0)
	if input.Admin != nil {
		setValues = append(setValues, fmt.Sprintf("admin = %d", *input.Admin))
	}

	if input.Time_start != nil {
		setValues = append(setValues, fmt.Sprintf("time_start = DATE '%s'", input.Time_start.Format(time.RFC3339)))
	}
	if input.Time_end != nil {
		setValues = append(setValues, fmt.Sprintf("Time_end = DATE '%s'", input.Time_end.Format(time.RFC3339)))
	}
	if input.Place != nil {
		setValues = append(setValues, fmt.Sprintf("place = '%s'", *input.Place))
	}
	if input.Game != nil {
		setValues = append(setValues, fmt.Sprintf("game = '%s'", *input.Game))
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description = '%s'", *input.Description))
	}
	if input.Public != nil {
		setValues = append(setValues, fmt.Sprintf("public = %t", *input.Public))
	}

	if len(setValues) == 0 {
		return errors.New("update structure has no values")
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE %s e SET %s WHERE e.event_items_id = $1 AND e.admin = $2`,
		eventItemsTable, setQuery)

	_, err := r.db.Exec(query, eventId, userId)
	return err
}

func (r *EventItemPostgres) DeleteEvent(userId, eventId int) error {
	query := fmt.Sprintf(`DELETE FROM %s e WHERE e.event_items_id = $1 AND e.admin = $2`,
		eventItemsTable)
	res, err := r.db.Exec(query, eventId, userId)

	if num, _ := res.RowsAffected(); num == 0 {
		return errors.New("error when deleting to db")
	}
	return err
}
