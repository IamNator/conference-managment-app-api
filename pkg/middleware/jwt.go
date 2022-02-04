package middleware

import (
	"conference/model"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("my_secret_key")

type Middleware struct {
	AccessTokenDuration time.Duration
	JWTSecretKey        []byte
}

func NewMiddleWare() *Middleware {
	return &Middleware{
		AccessTokenDuration: 30 * time.Minute,
		JWTSecretKey:        jwtKey,
	}
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func (m Middleware) GenerateToken(user model.User) (*model.UserAuthResponse, error) {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(m.AccessTokenDuration)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(m.JWTSecretKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return nil, err
	}

	return &model.UserAuthResponse{
		User:        user,
		AccessToken: tokenString,
		AccessExp:   expirationTime,
	}, nil
}

func (m *Middleware) Verify(accessToken string) (*uint, error) {

	claims := new(Claims)
	tkn, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		//if err == jwt.ErrSignatureInvalid {
		return nil, err
		//}
		//w.WriteHeader(http.StatusBadRequest)
		//return
	}
	if !tkn.Valid {
		return nil, errors.New("unauthorized access")
	}

	return &claims.UserID, nil
}
