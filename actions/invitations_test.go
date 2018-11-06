package actions

import (
	"github.com/gobuffalo/uuid"

	"github.com/invitation/models"
)

func (as *ActionSuite) Test_InvitationsResource_List() {
	as.LoadFixture("Test data")
	u := &models.User{}
	err := as.DB.Eager().Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)

	res := as.HTML("/invitations").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Show")
}

func (as *ActionSuite) Test_InvitationsResource_Show() {
	as.LoadFixture("Test data")
	u := &models.User{}
	err := as.DB.Eager().Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)
	i := u.Invitations[0].ID

	res := as.HTML("/invitations/" + i.String()).Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Sie sind herzlich eingeladen!")
}

func (as *ActionSuite) Test_InvitationsResource_Show_WrongID() {
	as.LoadFixture("Test data")
	u := &models.User{}
	err := as.DB.Eager().Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)
	i, _ := uuid.NewV4()

	res := as.HTML("/invitations/" + i.String()).Get()
	as.Equal(404, res.Code)
}

func (as *ActionSuite) Test_InvitationsResource_Show_WrongUser() {
	as.LoadFixture("Test data")
	u1 := &models.User{}
	u2 := &models.User{}
	err := as.DB.Eager().Where("email = ?", "marco@example.com").First(u1)
	err = as.DB.Eager().Where("email = ?", "sonja@example.com").First(u2)
	as.Session.Set("current_user_id", u1.ID)
	as.NoError(err)
	i := u2.Invitations[0].ID

	res := as.HTML("/invitations/" + i.String()).Get()
	as.Equal(404, res.Code)
}

func (as *ActionSuite) Test_InvitationsResource_New() {
	as.LoadFixture("Test data")
	u := &models.User{}
	err := as.DB.Eager().Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)

	res := as.HTML("/invitations/new").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Salutation")
}

type invitationTest struct {
	Mailtext   string
	Salutation int
	Name0      string
	Gender0    int
	Mail0      string
	Mail1      string
	Name1      string
	Gender1    int
	Mail2      string
	Name2      string
	Gender2    int
}

type invitationTest2 struct {
	Mailtext   string
	Salutation int
	Name0      string
	Gender0    int
	Mail0      string
	Mail1      string
	Name1      string
	Gender1    int
	Mail3      string
	Name3      string
	Gender3    int
}

func (as *ActionSuite) Test_InvitationsResource_Create() {
	as.LoadFixture("Test data")
	u := &models.User{}
	err := as.DB.Eager().Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)

	i := &invitationTest{
		Mailtext:   "Sie sind herzlich eingeladen! Mit freundlichen Gruessen",
		Salutation: 2,
		Name0:      "Alfred",
		Name1:      "Harald",
		Name2:      "Alex",
		Gender0:    1,
		Gender1:    2,
		Gender2:    3,
		Mail0:      "alfred@example.com",
		Mail1:      "harald@example.com",
		Mail2:      "alex@example.com",
	}

	res := as.HTML("/invitations").Post(i)
	as.Equal(302, res.Code)
	as.Contains(res.Location(), "/invitations/")
	count, err := as.DB.Count("invitations")
	as.NoError(err)
	as.Equal(3, count)

	count, err = as.DB.Count("guests")
	as.NoError(err)
	as.Equal(5, count)
}

func (as *ActionSuite) Test_InvitationsResource_Create_NoGuestsInBetween() {
	as.LoadFixture("Test data")
	u := &models.User{}
	err := as.DB.Eager().Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)

	i := &invitationTest2{
		Mailtext:   "Sie sind herzlich eingeladen! Mit freundlichen Gruessen",
		Salutation: 2,
		Name0:      "Alfred",
		Name1:      "Harald",
		Name3:      "Alex",
		Gender0:    1,
		Gender1:    2,
		Gender3:    3,
		Mail0:      "alfred@example.com",
		Mail1:      "harald@example.com",
		Mail3:      "alex@example.com",
	}

	res := as.HTML("/invitations").Post(i)
	as.Equal(302, res.Code)
	as.Contains(res.Header().Get("Location"), "/invitations/")
	count, err := as.DB.Count("invitations")
	as.NoError(err)
	as.Equal(3, count)

	count, err = as.DB.Count("guests")
	as.NoError(err)
	as.Equal(5, count)
}

func (as *ActionSuite) Test_InvitationsResource_Create_NoMailtext() {
	as.LoadFixture("Test data")
	u := &models.User{}
	err := as.DB.Eager().Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)

	i := &invitationTest{
		Mailtext:   "",
		Salutation: 2,
		Name0:      "Alfred",
		Name1:      "Harald",
		Name2:      "Alex",
		Gender0:    1,
		Gender1:    2,
		Gender2:    3,
		Mail0:      "alfred@example.com",
		Mail1:      "harald@example.com",
		Mail2:      "alex@example.com",
	}

	res := as.HTML("/invitations").Post(i)
	as.Equal(422, res.Code)
	count, err := as.DB.Count("invitations")
	as.NoError(err)
	as.Equal(2, count)

	count, err = as.DB.Count("guests")
	as.NoError(err)
	as.Equal(2, count)
}

func (as *ActionSuite) Test_InvitationsResource_Edit() {
	as.LoadFixture("Test data")
	u := &models.User{}
	err := as.DB.Eager().Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)
	i := u.Invitations[0].ID

	res := as.HTML("/invitations/" + i.String() + "/edit").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Edit")
}

type updateInvitationTest struct {
	Mailtext   string
	Salutation int
	Name0      string
	Gender0    int
	Mail0      string
	Mail1      string
	Name1      string
	Gender1    int
	Mail2      string
	Name2      string
	Gender2    int
	Mail3      string
	Name3      string
	Gender3    int
}

func (as *ActionSuite) Test_InvitationsResource_Update() {
	as.LoadFixture("Test data")
	u := &models.User{}
	err := as.DB.Eager().Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)

	i := &updateInvitationTest{
		Mailtext:   "Sie sind herzlich eingeladen! Mit freundlichen Gruessen",
		Salutation: 2,
		Name0:      "Alfred",
		Name1:      "Harald",
		Name2:      "Alex",
		Name3:      "Manfred",
		Gender0:    3,
		Gender1:    2,
		Gender2:    3,
		Gender3:    1,
		Mail0:      "alfred@example.com",
		Mail1:      "harald@example.com",
		Mail2:      "alex@example.com",
		Mail3:      "manfred@example.com",
	}

	count, err := as.DB.Where("invitationid = ?", u.Invitations[0].ID).Count("guests")
	as.NoError(err)
	as.Equal(2, count)

	res := as.HTML("/invitations/" + u.Invitations[0].ID.String()).Put(i)
	as.Equal(302, res.Code)
	as.Contains(res.Header().Get("Location"), "/invitations/")
	count, err = as.DB.Count("invitations")
	as.NoError(err)
	as.Equal(2, count)

	count, err = as.DB.Where("invitationid = ?", u.Invitations[0].ID).Count("guests")
	as.NoError(err)
	as.Equal(4, count)
}

func (as *ActionSuite) Test_InvitationsResource_Update_WrongID() {
	as.LoadFixture("Test data")
	u := &models.User{}
	wrongID, err1 := uuid.NewV4()
	err2 := as.DB.Eager().Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err1, err2)

	i := &updateInvitationTest{
		Mailtext:   "Sie sind herzlich eingeladen! Mit freundlichen Gruessen",
		Salutation: 2,
		Name0:      "Alfred",
		Name1:      "Harald",
		Name2:      "Alex",
		Name3:      "Manfred",
		Gender0:    3,
		Gender1:    2,
		Gender2:    3,
		Gender3:    1,
		Mail0:      "alfred@example.com",
		Mail1:      "harald@example.com",
		Mail2:      "alex@example.com",
		Mail3:      "manfred@example.com",
	}

	res := as.HTML("/invitations/" + wrongID.String()).Put(i)
	as.Equal(404, res.Code)
	count, err := as.DB.Count("invitations")
	as.NoError(err)
	as.Equal(2, count)

	count, err = as.DB.Count("guests")
	as.NoError(err)
	as.Equal(2, count)
}

func (as *ActionSuite) Test_InvitationsResource_Destroy() {
	as.LoadFixture("Test data")
	u := &models.User{}
	err := as.DB.Eager().Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)
	i := u.Invitations[0].ID
	res := as.HTML("/invitations/" + i.String()).Delete()
	as.Equal(302, res.Code)
	count, err := as.DB.Count("invitations")
	as.NoError(err)
	as.Equal(1, count)
}

func (as *ActionSuite) Test_InvitationsResource_Destroy_WrongID() {
	as.LoadFixture("Test data")
	u := &models.User{}
	err := as.DB.Eager().Where("email = ?", "sonja@example.com").First(u)
	as.Session.Set("current_user_id", u.ID)
	as.NoError(err)

	res := as.HTML("/invitations/" + "abcdefgh").Delete()
	as.Equal(404, res.Code)
	count, err := as.DB.Count("invitations")
	as.NoError(err)
	as.Equal(2, count)
}

func (as *ActionSuite) Test_formParser() {
	m := map[string][]string{
		"Mailtext":   []string{"Tester"},
		"Salutation": []string{"2"},
		"Name0":      []string{"Alfred"},
		"Name1":      []string{"Harald"},
		"Name2":      []string{"Alex"},
		"Gender0":    []string{"1"},
		"Gender1":    []string{"2"},
		"Gender2":    []string{"3"},
		"Mail0":      []string{"alfred@example.com"},
		"Mail1":      []string{"harald@example.com"},
		"Mail2":      []string{"alex@example.com"},
	}
	inv, err := formParser(m)
	as.NoError(err)
	as.Equal("Tester", inv.Mailtext)
}

func (as *ActionSuite) Test_formParser_NoMailtext() {
	m := map[string][]string{
		"Mailtext":   []string{""},
		"Salutation": []string{"2"},
		"Name0":      []string{"Alfred"},
		"Name1":      []string{"Harald"},
		"Name2":      []string{"Alex"},
		"Gender0":    []string{"1"},
		"Gender1":    []string{"2"},
		"Gender2":    []string{"3"},
		"Mail0":      []string{"alfred@example.com"},
		"Mail1":      []string{"harald@example.com"},
		"Mail2":      []string{"alex@example.com"},
	}
	_, err := formParser(m)
	as.Error(err)
}

func (as *ActionSuite) Test_formParser_IndexTooBig() {
	m := map[string][]string{
		"Gender101": []string{"2"},
	}
	_, err := formParser(m)
	as.Error(err)
}
