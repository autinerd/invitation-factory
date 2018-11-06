package models_test

import (
	"github.com/gobuffalo/pop"
	"github.com/invitation/models"
	"golang.org/x/crypto/bcrypt"
)

func (ms *ModelSuite) Test_User_Validate_Correct() {
	hash, _ := bcrypt.GenerateFromPassword([]byte("password"), 10)
	u := &models.User{
		Email:        "mark@example.com",
		Password:     "password",
		PasswordHash: string(hash),
	}
	verrs, err := u.Validate(&pop.Connection{Dialect: nil})
	ms.Equal(false, verrs.HasAny())
	ms.NoError(err)
}

func (ms *ModelSuite) Test_User_Validate_NoHash() {
	u := &models.User{
		Email:    "mark@example.com",
		Password: "password",
	}
	verrs, err := u.Validate(&pop.Connection{Dialect: nil})
	ms.Equal(true, verrs.HasAny())
	ms.NoError(err)
}

func (ms *ModelSuite) Test_User_Validate_WrongEmail() {
	hash, _ := bcrypt.GenerateFromPassword([]byte("password"), 10)
	u := &models.User{
		Email:        "sapodj",
		Password:     "password",
		PasswordHash: string(hash),
	}
	verrs, err := u.Validate(&pop.Connection{Dialect: nil})
	ms.Equal(true, verrs.HasAny())
	ms.NoError(err)
}

func (ms *ModelSuite) Test_User_Create() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.User{
		Email:                "mark@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
	}
	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(u.PasswordHash)

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_User_Create_ValidationErrors() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.User{
		Password: "password",
	}
	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)
}

func (ms *ModelSuite) Test_User_Create_UserExists() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &models.User{
		Email:                "mark@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
	}
	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(u.PasswordHash)

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)

	u = &models.User{
		Email:    "mark@example.com",
		Password: "password",
	}
	verrs, err = u.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}
