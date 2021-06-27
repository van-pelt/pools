package model

import (
	"github.com/go-pg/pg/v10"
	"strconv"
)

const tbl = "tbl_users"

type User struct {
	Id         int64
	FullName   string
	Email      string
	Password   string
	DateCreate string
	Hash       string
}

func (U *User) String() string {
	return "ID:" + strconv.Itoa(int(U.Id)) + " FullName:" + U.FullName + " Email:" + U.Email + " Password:" + U.Password + " DateCreate:" + U.DateCreate + " Hash:" + U.Hash
}

func GetUser(db *pg.DB, id int64) (*User, error) {
	var user User
	_, err := db.QueryOne(&user, `SELECT * FROM `+tbl+` WHERE id = ?`, id)
	return &user, err
}

func GetUsers(db *pg.DB) (*[]User, error) {
	var users []User
	_, err := db.Query(&users, `SELECT * FROM `+tbl)
	return &users, err
}

func GetUserByEmail(db *pg.DB, email string) (*User, error) {
	var user User
	_, err := db.QueryOne(&user, `SELECT * FROM `+tbl+` WHERE email = ?`, email)
	return &user, err
}

func GetUserByHash(db *pg.DB, hash string) (*User, error) {
	var user User
	_, err := db.QueryOne(&user, `SELECT * FROM `+tbl+` WHERE hash = ?`, hash)
	return &user, err
}

func DeleteUserByID(db *pg.DB, id int64) error {
	//var users []User
	_, err := db.Query(nil, `DELETE FROM `+tbl+` WHERE id=?`, id)
	return err
}
