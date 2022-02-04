package model

import (
	"conference/testdata"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestUserSignUpReq_Validate(t *testing.T) {

	type tData struct {
		Req   UserSignUpReq `json:"req"`
		Error string        `json:"error"`
	}

	SignUpReqs := make([]tData, 0)
	if er := testdata.Load("../testdata/model/signup.json", &SignUpReqs); er != nil {
		t.Error(er.Error())
	}

	for i, v := range SignUpReqs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			//check for validation errors
			if er := v.Req.Validate(); er != nil {
				assert.Equal(t, v.Error, er.Error())
			}
		})
	}

}

func TestUserLoginReq_Validate(t *testing.T) {

	type tData struct {
		Req   UserLoginReq `json:"req"`
		Error string       `json:"error"`
	}

	LoginReqs := make([]tData, 0)
	if er := testdata.Load("../testdata/model/login.json", &LoginReqs); er != nil {
		t.Error(er.Error())
	}

	for i, v := range LoginReqs {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			//check for validation errors
			if er := v.Req.Validate(); er != nil {
				assert.Equal(t, v.Error, er.Error())
			}
		})
	}
}

func TestPassword_Compare(t *testing.T) {

	tt := []struct {
		Password     Password `json:"password"`
		HashPassword Password `json:"hash_password"`
		Error        bool     `json:"error"`
	}{
		{
			Password:     "pass",
			HashPassword: Password("pass").Hash(),
			Error:        false,
		},
		{
			Password:     "pass",
			HashPassword: Password("passd").Hash(),
			Error:        true,
		},
		{
			Password:     "pass",
			HashPassword: Password("pass").Hash(),
			Error:        false,
		},
		{
			Password:     "password",
			HashPassword: Password("password").Hash(),
			Error:        false,
		},
		{
			Password:     "passworrr",
			HashPassword: Password("pass").Hash(),
			Error:        true,
		},
	}

	for i, v := range tt {
		t.Run(fmt.Sprintf("test: %d", i), func(t *testing.T) {
			assert.Equal(t, true, len(v.HashPassword) > 50)
			assert.Equal(t, !v.Error, v.HashPassword.Compare(v.Password.String()))
			assert.Equal(t, !v.Error, v.HashPassword.Hash().Compare(v.Password.String()))
			assert.NotEqual(t, v.HashPassword, v.Password)
		})
	}

}

func TestPassword_String(t *testing.T) {
	inputtedPassword := "password"
	actualPassword := Password("password")
	assert.Equal(t, inputtedPassword, actualPassword.String())
}

func TestUser_TableName(t *testing.T) {
	var user User
	assert.Equal(t, user.TableName(), "user")
}

func TestPassword_Hash(t *testing.T) {
	actualPassword := Password("password")
	hashedPassword := actualPassword.Hash()
	assert.NotEqual(t, actualPassword, hashedPassword)
}
