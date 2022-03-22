package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// MyClaims custom declare structure and embed jwt.StandardClaims
type MyClaims struct {
	MyID int `json:"id"`
	jwt.StandardClaims
}

// RefreshClaims
type RefreshClaims struct {
	Account        string `json:"account"`
	ClientId       string `json:"client_id"`
	Certificate    string `json:"certificate"`
	CertificateKey string `json:"certificateKey"`
	jwt.StandardClaims
}

var (
	// _tokenExpireDuration token expire time
	_tokenExpireDuration = time.Hour * 2 // default 2 hour
	// secret
	_secret = []byte(os.Getenv("TOKEN_SECURITY_KEY"))
)

// GenToken generate JWT
func GenToken(id int) (string, error) {
	// create
	c := MyClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(_tokenExpireDuration).Unix(),
			Issuer:    "tre-china",
		},
	}
	// sign
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// encode
	return token.SignedString(_secret)
}

// ParseToken parse JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// parse token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(*jwt.Token) (i interface{}, err error) {
		return _secret, nil
	})
	if err != nil {
		return nil, err
	}
	// valid token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
