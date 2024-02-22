package userservice

import (
	"errors"
	"strings"

	"github.com/hector-leite/meli-notification/domain/constants"
	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/domain/entity"
	"github.com/hector-leite/meli-notification/domain/service/serviceutil"
	"github.com/hector-leite/meli-notification/infra/context"
	"github.com/hector-leite/meli-notification/infra/utils"

	uuid "github.com/satori/go.uuid"
)

type UserService struct {
	authTokenService contract.AuthTokenService
	data             contract.DataManager
}

func NewUserService(
	authTokenService contract.AuthTokenService,
	data contract.DataManager,
) contract.UserService {
	return UserService{
		authTokenService: authTokenService,
		data:             data,
	}
}

func (s UserService) GetFromContext(context context.Context) (entity.User, error) {
	rawUser, ok := context.Get(constants.ContextKeyUser)
	if ok {
		return rawUser.(entity.User), nil
	}

	rawUUID, ok := context.Get(constants.ContextKeyUserUUID)
	if !ok {
		return entity.User{}, errors.New("context-error")
	}
	userUUID := rawUUID.(string)

	user, err := s.data.User().FindByUUID(userUUID)
	if err != nil {
		return user, err
	}

	context.Set(constants.ContextKeyUser, user)

	return user, nil
}

func (s UserService) SignUp(password string, user entity.User) (err error) {
	user.Name = strings.ToLower(user.Name)
	user.UUID = uuid.NewV4().String()
	user.HashPassword, err = utils.Hash(utils.SaltPassword(password, user.UUID))
	if err != nil {
		return err
	}

	_, err = s.data.User().SignUp(user)
	if err != nil {
		return err
	}

	return nil
}

func (s UserService) SignIn(password string, user entity.User) (entity.User, error) {
	user, err := s.data.User().FindByEmail(user.Email)
	if err != nil {
		return entity.User{}, err
	}

	password = utils.SaltPassword(password, user.UUID)
	valid, err := utils.ValidateHash(password, user.HashPassword)
	if err != nil {
		return entity.User{}, err
	}
	if !valid {
		return entity.User{}, errors.New("invalid-credentials")
	}

	user.Tokens, err = s.getTokens(user)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (s UserService) Find() ([]entity.User, error) {
	users, err := s.data.User().Find()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s UserService) RefreshToken(hashAuthToken string) (entity.User, error) {
	tokenBody, err := serviceutil.UnwrapTokenBody(hashAuthToken)
	if err != nil {
		return entity.User{}, errors.New("invalid-refresh-token")
	}

	authToken, err := s.authTokenService.FindUnexpiredByUUIDAndType(tokenBody.UUID, constants.TokenTypeRefresh)
	if err != nil {
		return entity.User{}, err
	}

	valid, err := utils.ValidateHash(tokenBody.Token, authToken.HashToken)
	if err != nil {
		return entity.User{}, err
	}
	if !valid {
		return entity.User{}, errors.New("invalid-refresh-token")
	}

	user, err := s.data.User().FindByUUID(tokenBody.UserUUID)
	if err != nil {
		return entity.User{}, err
	}

	err = s.authTokenService.DeleteByID(authToken.ID)
	if err != nil {
		return entity.User{}, err
	}

	user.Tokens, err = s.getTokens(user)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (s UserService) getTokens(user entity.User) ([]entity.TokenWrapper, error) {
	var tokens []entity.TokenWrapper
	for _, tokenType := range []string{constants.TokenTypeRefresh, constants.TokenTypeAccess} {
		tokenWrapper, err := s.authTokenService.NewTokenWrapper(user, tokenType)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, tokenWrapper)
	}

	return tokens, nil
}
