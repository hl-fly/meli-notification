package authtokenservice

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"

	"github.com/hector-leite/meli-notification/domain/constants"
	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/domain/entity"
	"github.com/hector-leite/meli-notification/domain/service/serviceutil"
	"github.com/hector-leite/meli-notification/infra/auth"
	uuid "github.com/satori/go.uuid"
)

type AuthTokenService struct {
	data          contract.DataManager
	jwtPrivateKey string
}

func NewAuthTokenService(
	data contract.DataManager,
	jwtPrivateKey string,
) contract.AuthTokenService {
	return AuthTokenService{
		data:          data,
		jwtPrivateKey: jwtPrivateKey,
	}
}

func (s AuthTokenService) NewTokenWrapper(user entity.User, tokenType string) (entity.TokenWrapper, error) {
	if tokenType == constants.TokenTypeAccess {
		tokenWrapper, err := s.newAccessTokenWrapper(user)
		if err != nil {
			return entity.TokenWrapper{}, err
		}

		return tokenWrapper, nil
	}

	token, hashToken, err := auth.NewToken()
	if err != nil {
		return entity.TokenWrapper{}, err
	}

	var authToken entity.AuthToken
	authToken.HashToken = hashToken
	authToken.Type = tokenType
	authToken.UserID = &user.ID
	authToken.UUID = uuid.NewV4().String()

	if tokenType != constants.TokenTypeRefresh {
		return entity.TokenWrapper{}, errors.New("invalid-token-type")
	}
	authToken.ExpirationDate = time.Now().AddDate(0, 0, 30)

	authToken.ID, err = s.data.AuthToken().Insert(authToken)
	if err != nil {
		return entity.TokenWrapper{}, err
	}

	authToken, err = s.data.AuthToken().FindUnexpiredByIDAndType(authToken.ID, authToken.Type)
	if err != nil {
		return entity.TokenWrapper{}, err
	}

	var body serviceutil.TokenBody
	body.Token = token
	body.UUID = authToken.UUID
	body.UserUUID = user.UUID

	dataBytes, err := json.Marshal(body)
	if err != nil {
		return entity.TokenWrapper{}, err
	}

	var tokenWrapper entity.TokenWrapper
	tokenWrapper.Type = tokenType
	tokenWrapper.HashAuthToken = base64.RawURLEncoding.EncodeToString(dataBytes)

	return tokenWrapper, nil
}

func (s AuthTokenService) FindUnexpiredByUUIDAndType(authTokenUUID string, authTokenType string) (entity.AuthToken, error) {
	authToken, err := s.data.AuthToken().FindUnexpiredByUUIDAndType(authTokenUUID, authTokenType)
	if err != nil {
		return entity.AuthToken{}, err
	}

	return authToken, nil
}

func (s AuthTokenService) DeleteByID(authTokenID uint) error {
	return s.data.AuthToken().DeleteByID(authTokenID)
}

func (s AuthTokenService) newAccessTokenWrapper(user entity.User) (entity.TokenWrapper, error) {
	claims := new(auth.UserClaims)
	claims.UUID = user.UUID
	claims.Name = user.Name
	claims.Role = user.Type

	accessToken, err := auth.NewAccessToken(claims.UUID, claims, s.jwtPrivateKey)
	if err != nil {
		return entity.TokenWrapper{}, err
	}

	var tokenWrapper entity.TokenWrapper
	tokenWrapper.Type = constants.TokenTypeAccess
	tokenWrapper.HashAuthToken = accessToken

	return tokenWrapper, nil
}
