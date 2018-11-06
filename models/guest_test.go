package models_test

import (
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/invitation/models"
)

func (ms *ModelSuite) Test_Guests_Validate_Correct() {
	g := models.Guest{
		AdditionalComment: "",
		Email:             "test@berber.de",
		Gender:            1,
		ID:                uuid.Must(uuid.NewV4()),
		InvitationID:      uuid.Must(uuid.NewV4()),
		Name:              "Sylvia",
		Status:            0,
	}
	verrs, err := g.Validate(&pop.Connection{})
	ms.NoError(err)
	ms.Equal(false, verrs.HasAny())
}

func (ms *ModelSuite) Test_Guests_Validate_WrongGender() {
	g := models.Guest{
		AdditionalComment: "",
		Email:             "test@berber.de",
		Gender:            12,
		ID:                uuid.Must(uuid.NewV4()),
		InvitationID:      uuid.Must(uuid.NewV4()),
		Name:              "Sylvia",
		Status:            0,
	}
	verrs, err := g.Validate(&pop.Connection{})
	ms.NoError(err)
	ms.Equal(true, verrs.HasAny())
}

func (ms *ModelSuite) Test_Guests_Validate_NoName() {
	g := models.Guest{
		AdditionalComment: "",
		Email:             "test@berber.de",
		Gender:            1,
		ID:                uuid.Must(uuid.NewV4()),
		InvitationID:      uuid.Must(uuid.NewV4()),
		Name:              "",
		Status:            0,
	}
	verrs, err := g.Validate(&pop.Connection{})
	ms.NoError(err)
	ms.Equal(true, verrs.HasAny())
}

func (ms *ModelSuite) Test_Guests_Validate_WrongMail() {
	g := models.Guest{
		AdditionalComment: "",
		Email:             "testberber.de",
		Gender:            1,
		ID:                uuid.Must(uuid.NewV4()),
		InvitationID:      uuid.Must(uuid.NewV4()),
		Name:              "Sylvia",
		Status:            0,
	}
	verrs, err := g.Validate(&pop.Connection{})
	ms.NoError(err)
	ms.Equal(true, verrs.HasAny())
}
