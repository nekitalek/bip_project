package BIP_project

import "time"

type User_auth struct {
	User_id            int    `json:"user_id" db:"user_id"`
	Login              string `json:"login" binding:"required"`
	Username           string `json:"username" binding:"required"`
	Password           string `json:"password" binding:"required"`
	Password_hash      string
	Email_confirmation bool
}

type User_data struct {
	User_id  int    `json:"user_id" db:"user_id"`
	Login    string `json:"login" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type User_data_Participant struct {
	User_id  int    `json:"user_id" db:"user_id"`
	Username string `json:"username" binding:"required"`
}

type Login_method string

const (
	Password Login_method = "Password"
	Code2fa  Login_method = "Code2fa"
)

type Login_attempt struct {
	Login_attempt_id int
	Email            string
	Number_of_inputs int
	Unlock_time      time.Time
	Login_method     Login_method
}

//для Email_confirmation
type Assignment string

const (
	Registration   Assignment = "Registration"
	SecFactor      Assignment = "SecFactor"
	ChangeEmail    Assignment = "ChangeEmail"
	ChangePassword Assignment = "ChangePassword"
)

type Email_confirmation struct {
	Email_confirmation_id   int        `json:"email_confirmation_id"`
	Email                   string     `json:"email"`
	User_id                 int        `json:"user_id"` // binding:"required"`
	Code_email_confirmation int        `json:"code"`
	Time_end                time.Time  `json:"time_end"`
	Assignment              Assignment `json:"assignment"`
	Device                  string     `json:"device"` // binding:"required"`
}

type Auth_data struct {
	Auth_data_id int       `db:"Auth_data_id"`
	Token        string    `db:"token"`
	User_id      int       `db:"user_id"`
	Time_end     time.Time `db:"time_end"`
	Device       string    `db:"device"`
}

//____________________________________________________________________________________________________________________________________

type Event_items struct {
	Event_items_id int                     `json:"event_items_id" form:"event_items_id"`
	Admin          int                     `json:"admin" form:"admin"`
	Participant    []User_data_Participant `json:"participant" form:"participant"`
	Time_start     time.Time               `json:"time_start" form:"time_start"`
	Time_end       time.Time               `json:"time_end" form:"time_end"`
	Place          string                  `json:"place" form:"place"`
	Game           string                  `json:"game" form:"game"`
	Description    string                  `json:"description" form:"description"`
	Public         bool                    `json:"public" form:"public"`
}

type Event_items_input struct {
	Event_items_id *int       `json:"event_items_id" form:"event_items_id"`
	Admin          *int       `json:"admin" form:"admin"`
	Time_start     *time.Time `json:"time_start" form:"time_start"`
	Time_end       *time.Time `json:"time_end" form:"time_end"`
	Place          *string    `json:"place" form:"place"`
	Game           *string    `json:"game" form:"game"`
	Description    *string    `json:"description" form:"description"`
	Public         *bool      `json:"public" form:"public"`
}

type Event_invitations struct {
	Event_invitations_id int    `json:"event_invitations_id"`
	Event_id             int    `json:"event_id"`
	User_id              int    `json:"user_id"`
	Status               string `json:"status"`
}

type Status string

const (
	Pending   Status = "Pending"
	Confirmed Status = "Confirmed"
	Rejected  Status = "Rejected"
)

type Event_invitations_input struct {
	Event_invitations_id *int    `json:"event_invitations_id" form:"event_invitations_id"`
	Event_id             *int    `json:"event_id" form:"event_id"`
	User_id              *int    `json:"user_id" form:"user_id"`
	Status               *Status `json:"status" form:"status"`
}

type Push_notification_input struct {
	Token  string `json:"token"`
	Device string `json:"device"`
}
