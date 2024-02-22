package data

import (
	"github.com/hector-leite/meli-notification/domain/entity"
	"gorm.io/gorm"
)

type authTokenRepo struct {
	db *gorm.DB
}

func newAuthTokenRepo(db *gorm.DB) authTokenRepo {
	return authTokenRepo{
		db: db,
	}
}

func (r authTokenRepo) Insert(authToken entity.AuthToken) (uint, error) {
	result := r.db.Create(&authToken)
	if result.Error != nil {
		return 0, result.Error
	}

	return authToken.ID, nil
}

func (r authTokenRepo) FindUnexpiredByIDAndType(authTokenID uint, authTokenType string) (entity.AuthToken, error) {
	var token entity.AuthToken
	result := r.db.Where("ID = ? AND type = ? and expiration_date >= NOW()", authTokenID, authTokenType).First(&token)
	if result.Error != nil {
		return entity.AuthToken{}, result.Error
	}

	return token, nil
}

func (r authTokenRepo) FindUnexpiredByUUIDAndType(authTokenUUID string, authTokenType string) (entity.AuthToken, error) {
	var token entity.AuthToken
	result := r.db.Where("UUID = ? AND type = ? and expiration_date >= NOW()", authTokenUUID, authTokenType).First(&token)
	if result.Error != nil {
		return entity.AuthToken{}, result.Error
	}

	return token, nil
}

func (r authTokenRepo) DeleteByID(authTokenID uint) error {
	result := r.db.Delete(&entity.AuthToken{}, authTokenID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
