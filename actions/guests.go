package actions

import (
	"strconv"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/invitation/models"
	"github.com/pkg/errors"
)

// DeleteGuestFromUnsubscribe deletes a guest when he unsubscribes
func DeleteGuestFromUnsubscribe(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	guest := &models.Guest{}
	guests := models.Guests{}

	if err := tx.Find(guest, c.Param("guest_id")); err != nil {
		return c.Error(404, err)
	}

	tx.Where("email = ?", guest.Email).All(&guests)

	if err := tx.Destroy(&guests); err != nil {
		return errors.WithStack(err)
	}

	return c.Render(200, r.String("Your e-mail was deleted successfully"))
}

// StatusResponse serves the page where the guest can response to the invitation.
func StatusResponse(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	guest := &models.Guest{}

	if err := tx.Find(guest, c.Param("guest_id")); err != nil {
		return c.Error(404, err)
	}

	if guest.Status > 0 {
		c.Flash().Add("danger", "You already responded to this invitation!")
		return c.Redirect(302, "/")
	}

	c.Set("action_url", "/invitations/"+guest.InvitationID.String()+"/guests/"+guest.ID.String())

	return c.Render(200, r.HTML("guests/response.html"))
}

// SetStatusResponse puts the guest response data into the database.
func SetStatusResponse(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	guest := &models.Guest{}

	if err := tx.Find(guest, c.Param("guest_id")); err != nil {
		return c.Error(404, err)
	}

	guest.Status, _ = strconv.Atoi(getFormValue(c, "status"))
	guest.AdditionalComment = getFormValue(c, "additional_comment")

	c.Flash().Add("success", "Your response was successfully transmitted!")
	tx.Update(guest)

	return c.Redirect(302, "/")
}

func getFormValue(c buffalo.Context, s string) string {
	s1 := strings.ToLower(s)
	sret := c.Request().FormValue(s1)
	if sret == "" {
		s1 = strings.Join([]string{strings.ToUpper(string(s1[0])), s1[1:len(s1)]}, "")
		sret = c.Request().FormValue(s1)
	}
	return sret
}
