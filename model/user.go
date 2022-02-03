package model

import "time"

type (
	Password string
	User     struct {
		Username     string     `json:"username"`
		Email        string     `json:"email"`
		Password     Password   `json:"password"`
		LastLoggedIn *time.Time `json:"last_logged_in"`
		General
	}

	UserLoginReq struct {
		Email    string   `json:"email"`
		Password Password `json:"password"`
	}

	UserSignUpReq struct {
		Username string   `json:"username"`
		Email    string   `json:"email"`
		Password Password `json:"password"`
	}

	UserAuthResponse struct {
		User User `json:"user"`

		AccessToken string    `json:"access_token"`
		AccessExp   time.Time `json:"access_exp"`

		RefreshToken string    `json:"refresh_token"`
		RefreshExp   time.Time `json:"refresh_exp"`
	}
)

func (User) TableName() string {
	return "user"
}

//TODO
func (p Password) Hash() Password {
	return p
}

func (p Password) Compare(password string) bool {
	return p == Password(password)
}

func (p Password) String() string {
	return string(p)
}
