package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"time"
)

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

	UserLogOutReq struct {
		RefreshToken string `json:"refresh_token"`
		AccessToken  string `json:"access_token"`
	}
)

func (User) TableName() string {
	return "user"
}

func (u User) CreateTable(tx *gorm.DB) error {
	return tx.AutoMigrate(u)
}

//TODO
func (p Password) Hash() Password {
	//return if password is already hashed
	if len(p.String()) > 50 {
		return p
	}

	hashedPassword, er := bcrypt.GenerateFromPassword([]byte(p.String()), bcrypt.DefaultCost)
	if er != nil {
		log.Println(er.Error())
	}

	return Password(hashedPassword)
}

func (p Password) Compare(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p.String()), []byte(password)) == nil

}

func (p Password) String() string {
	return string(p)
}

func (u UserSignUpReq) Validate() error {
	return validation.ValidateStruct(&u,
		// Username cannot be empty, and the length must between 1 and 50
		validation.Field(&u.Username, validation.Required, validation.Length(1, 50)),
		// Password cannot be empty, and the length must between 5 and 50
		validation.Field(&u.Password, validation.Required, validation.Length(5, 50)),
		// Email cannot be empty and should be in a valid email format.
		validation.Field(&u.Email, validation.Required, is.Email),
	)
}

func (u UserLoginReq) Validate() error {
	return validation.ValidateStruct(&u,
		// Password cannot be empty, and the length must between 5 and 50
		validation.Field(&u.Password, validation.Required, validation.Length(5, 50)),
		// Email cannot be empty and should be in a valid email format.
		validation.Field(&u.Email, validation.Required, is.Email),
	)
}

func (u UserLogOutReq) Validate() error {
	return validation.ValidateStruct(&u,
		// Password cannot be empty, and the length must between 5 and 50
		validation.Field(&u.RefreshToken, validation.Required),
		// Email cannot be empty and should be in a valid email format.
		validation.Field(&u.RefreshToken, validation.Required),
	)
}
