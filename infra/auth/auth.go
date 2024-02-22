package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/hector-leite/meli-notification/infra/utils"
)

var (
	TokenSigningMethod = jwt.SigningMethodHS256
)

const (
	HTTPHeader            = "Authorization"
	AccessTokenExpiration = time.Hour
)

type UserClaims struct {
	jwt.StandardClaims
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func (c *UserClaims) SetStandardClaims(standardClaims jwt.StandardClaims) {
	c.StandardClaims = standardClaims
}

type Claims interface {
	jwt.Claims
	SetStandardClaims(standardClaims jwt.StandardClaims)
}

func NewAccessToken(uuid string, claims Claims, privateKey string) (string, error) {
	now := time.Now()
	claims.SetStandardClaims(jwt.StandardClaims{
		Issuer:    "Plimbou",
		Subject:   uuid,
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(AccessTokenExpiration).Unix(),
	})

	token, err := jwt.NewWithClaims(TokenSigningMethod, claims).SignedString([]byte(privateKey))
	if err != nil {
		return token, err
	}

	return token, nil
}

func NewToken() (string, []byte, error) {
	token, err := utils.RandomString(64)
	if err != nil {
		return "", nil, err
	}

	hashToken, err := utils.Hash(token)
	if err != nil {
		return "", nil, err
	}

	return token, hashToken, nil
}
