package models

import (
	"time"

	"github.com/gobuffalo/validate/validators"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"

	"github.com/gobuffalo/uuid"
)

// Guest is the declaration of a guest
type Guest struct {
	ID                uuid.UUID `json:"id" db:"id"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
	InvitationID      uuid.UUID `json:"invitationid" db:"invitationid"`
	Email             string    `json:"email" db:"email"`
	Gender            int       `json:"gender" db:"gender"`
	Name              string    `json:"name" db:"name"`
	Status            int       `json:"status" db:"status"`
	AdditionalComment string    `json:"additional_comment" db:"additional_comment"`
}

// Guests is an array of Guest
type Guests []Guest

// Validate checks if a guest is correctly entered
func (g *Guest) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.EmailIsPresent{Field: g.Email, Name: "Email"},
		&validators.StringIsPresent{Field: g.Name, Name: "Name"},
		&validators.IntIsGreaterThan{Field: g.Gender, Compared: 0, Name: "Gender"},
		&validators.IntIsLessThan{Field: g.Gender, Compared: 4, Name: "Gender"}), err
}
