package models

import (
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"

	"github.com/gobuffalo/uuid"
)

// Invitation is the structure
type Invitation struct {
	ID           uuid.UUID `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	UserID       uuid.UUID `json:"userid" db:"userid"`
	Salutation   int       `json:"salutation" db:"salutation"`
	Mailtext     string    `json:"mailtext" db:"mailtext"`
	SentToGuests bool      `json:"senttoguests" db:"senttoguests"`
	Guests       Guests    `has_many:"guests" fk_id:"invitationid"`
}

// Invitations is the invitation slice.
type Invitations []Invitation

// ContainsGuests is a validator which checks if an invitation has at least one guest.
type ContainsGuests struct {
	Field   Guests
	Message string
}

// IsValid checks if an invitation has at least one guest.
func (c *ContainsGuests) IsValid(errors *validate.Errors) {
	if len(c.Field) < 1 {
		errors.Add("guests", "At least one guest is neccessary!")
	}
	return
}

// Validate validates an invitation.
func (i *Invitation) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: i.Mailtext, Name: "Text body"},
		&ContainsGuests{
			Field: i.Guests,
		},
		&validators.IntIsGreaterThan{Field: i.Salutation, Compared: -1, Name: "Salutation"},
		&validators.IntIsLessThan{Field: i.Salutation, Compared: 4, Name: "Salutation"}), err
}
