package actions

import (
	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/invitation/mailers"
	"github.com/invitation/models"
	"github.com/pkg/errors"
)

// InvitMailSend sends the invitation mails to all guests.
func InvitMailSend(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	u := c.Value("current_user").(*models.User)
	if !ok {
		log.Println(errors.WithStack(errors.New("no transaction found")).Error())
		return c.Error(500, errors.New("Internal Server Error"))
	}

	invitation := &models.Invitation{}

	// To find the Invitation the parameter invitation_id is used.
	if err := tx.Eager().Find(invitation, c.Param("invitation_id")); err != nil {
		return c.Error(404, err)
	}

	if invitation.UserID != u.ID {
		c.Flash().Add("danger", "You are not allowed to visit this page!")
		return c.Redirect(302, "/invitations")
	}

	if err := mailers.SendInvitMail(invitation, u); err != nil {
		return errors.WithStack(err)
	}

	invitation.SentToGuests = true

	tx.Update(invitation)

	c.Flash().Add("primary", "Mails were successfully sent!")
	return c.Redirect(302, "/invitations/"+c.Param("invitation_id"))
}
