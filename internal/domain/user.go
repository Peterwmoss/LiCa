package domain

type (
	User interface {
		Email() string
		Picture() string
	}

	user struct {
		email   string
		picture string
	}
)

func NewUser(email string, picture string) User {
	return &user{
		email:   email,
		picture: picture,
	}
}

func (u user) Email() string {
	return u.email
}

func (u user) Picture() string {
	return u.picture
}
