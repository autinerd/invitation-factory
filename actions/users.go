package actions

import (
	"log"

	"github.com/invitation/mailers"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/invitation/models"
	"github.com/pkg/errors"
)

// UsersNew opens the register form.
func UsersNew(c buffalo.Context) error {
	u := models.User{}
	c.Set("user", u)
	return c.Render(200, r.HTML("users/new.html"))
}

// UsersCreate registers a new user with the application.
func UsersCreate(c buffalo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return c.Error(500, err)
	}
	u.Verified = false
	tx := c.Value("tx").(*pop.Connection)
	verrs, err := u.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("user", u)
		c.Set("errors", verrs)
		return c.Render(422, r.HTML("users/new.html"))
	}

	//c.Session().Set("current_user_id", u.ID)
	c.Flash().Add("success", "Welcome to the invitation factory! Please verify your mail address")
	mailers.SendVerifyMail(u)
	return c.Redirect(302, "/")
}

// SetCurrentUser attempts to find a user based on the current_user_id
// in the session. If one is found it is set on the context.
func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Eager("Invitations.Guests").Find(u, uid)
			if err != nil || u.Verified == false {
				c.Session().Clear()
				return c.Redirect(302, "/")
			}
			c.Set("current_user", u)
		}
		return next(c)
	}
}

// Authorize require a user be logged in before accessing a route
func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {
			c.Flash().Add("danger", "You have to be logged in to see this page!")
			return c.Redirect(302, "/")
		}
		return next(c)
	}
}

// VerifyUser verifies a user when he clicked on the link in verify mail.
func VerifyUser(c buffalo.Context) error {
	u := &models.User{}
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		log.Println("no transactions found")
		c.Flash().Add("danger", "Error while verifying!")
		return c.Redirect(302, "/")
	}

	if err := tx.Find(u, c.Param("user_id")); err != nil {
		log.Println(err)
		return err
	}

	// Verifying user
	u.Verified = true
	tx.Update(u)

	c.Flash().Add("success", "Successfully verified! You can now log in.")
	return c.Redirect(302, "/signin")
}
