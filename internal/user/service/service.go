package service

import (
	"github.com/van-pelt/pools/internal/user/model"
	"github.com/van-pelt/pools/pkg/database"
)

type UserService struct {
	db *database.Storage
}

func New(db *database.Storage) *UserService {
	return &UserService{db: db}
}

func (U *UserService) CheckUserData(email string) (*model.User, error) {
	Usr, err := model.GetUserByEmail(U.db.DB, email)
	if err != nil {
		return nil, err
	}
	return Usr, nil
}

func (U *UserService) GetUsersData() (*[]model.User, error) {
	Usr, err := model.GetUsers(U.db.DB)
	if err != nil {
		return nil, err
	}
	return Usr, nil
}

func (U *UserService) DelUsersData(id int64) error {
	return model.DeleteUserByID(U.db.DB, id)
}

func (U *UserService) GetUserByID(id int64) (*model.User, error) {
	Usr, err := model.GetUser(U.db.DB, id)
	if err != nil {
		return nil, err
	}
	return Usr, nil
}
