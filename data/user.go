package data

import (
	"github.com/hector-leite/meli-notification/domain/contract"
	"github.com/hector-leite/meli-notification/domain/entity"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func newUserRepo(db *gorm.DB) contract.UserRepo {
	return userRepo{
		db: db,
	}
}

func (r userRepo) SignUp(user entity.User) (uint, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func (r userRepo) Find() ([]entity.User, error) {
	var users []entity.User

	result := r.db.Find(&users)
	if result.Error != nil {
		return []entity.User{}, result.Error
	}

	return users, nil
}

func (r userRepo) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	result := r.db.Where("email = ?", email).Find(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

func (r userRepo) FindByUUID(UUID string) (entity.User, error) {
	var user entity.User

	result := r.db.Where("uuid = ?", UUID).Find(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}
