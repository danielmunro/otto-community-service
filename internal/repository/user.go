package repository

import (
	"errors"
	"github.com/danielmunro/otto-community-service/internal/constants"
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

func CreateUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{conn}
}

func (u *UserRepository) FindOne(id uint) (*entity.User, error) {
	user := &entity.User{}
	u.conn.Where("id = ?", id).Find(user)
	if user.ID == 0 {
		return nil, errors.New(constants.ErrorMessageUserNotFound)
	}
	return user, nil
}

func (u *UserRepository) FindOneByUuid(uuid string) (*entity.User, error) {
	user := &entity.User{}
	u.conn.Where("uuid = ?", uuid).Find(user)
	if user.ID == 0 {
		return nil, errors.New(constants.ErrorMessageUserNotFound)
	}
	return user, nil
}

func (u *UserRepository) FindUsersNotFollowing(userUuid uuid.UUID) []*entity.User {
	var users []*entity.User
	u.conn.Raw("SELECT * " +
	"FROM users " +
	"WHERE id NOT IN (SELECT following_id FROM follows f JOIN users u ON f.user_id = u.id WHERE u.uuid = ?)", userUuid.String()).
	Scan(&users)
	return users
}

func (u *UserRepository) Create(user *entity.User) {
	u.conn.Create(user)
}

func (u *UserRepository) Update(user *entity.User) {
	u.conn.Model(&user).Updates(user)
}

func (u *UserRepository) Delete(user *entity.User) {
	u.conn.Delete(user)
}
