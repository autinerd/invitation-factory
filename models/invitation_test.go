package models_test

import (
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/invitation/models"
)

func (ms *ModelSuite) Test_Inv_Validate_Correct() {
	i := models.Invitation{
		Mailtext:     "Dies ist ein Test",
		ID:           uuid.Must(uuid.NewV4()),
		Salutation:   1,
		SentToGuests: false,
		UserID:       uuid.Must(uuid.NewV4()),
		Guests: models.Guests{
			models.Guest{
				Email:  "test@bla.de",
				Gender: 2,
				ID:     uuid.Must(uuid.NewV4()),
			},
		},
	}
	verrs, err := i.Validate(&pop.Connection{})
	ms.NoError(err)
	ms.Equal(false, verrs.HasAny())
}

func (ms *ModelSuite) Test_Inv_Validate_NoMailtext() {
	i := models.Invitation{
		Mailtext:     "",
		ID:           uuid.Must(uuid.NewV4()),
		Salutation:   1,
		SentToGuests: false,
		UserID:       uuid.Must(uuid.NewV4()),
		Guests: models.Guests{
			models.Guest{
				Email:  "test@bla.de",
				Gender: 2,
				ID:     uuid.Must(uuid.NewV4()),
			},
		},
	}
	verrs, err := i.Validate(&pop.Connection{})
	ms.NoError(err)
	ms.Equal(true, verrs.HasAny())
}

func (ms *ModelSuite) Test_Inv_Validate_NoGuests() {
	i := models.Invitation{
		Mailtext:     "Test",
		ID:           uuid.Must(uuid.NewV4()),
		Salutation:   1,
		SentToGuests: false,
		UserID:       uuid.Must(uuid.NewV4()),
	}
	verrs, err := i.Validate(&pop.Connection{})
	ms.NoError(err)
	ms.Equal(true, verrs.HasAny())
}

func (ms *ModelSuite) Test_Inv_Validate_InvalidSalutation() {
	i := models.Invitation{
		Mailtext:     "Test",
		ID:           uuid.Must(uuid.NewV4()),
		Salutation:   20,
		SentToGuests: false,
		UserID:       uuid.Must(uuid.NewV4()),
		Guests: models.Guests{
			models.Guest{
				Email:  "test@bla.de",
				Gender: 2,
				ID:     uuid.Must(uuid.NewV4()),
			},
		},
	}
	verrs, err := i.Validate(&pop.Connection{})
	ms.NoError(err)
	ms.Equal(true, verrs.HasAny())
}
