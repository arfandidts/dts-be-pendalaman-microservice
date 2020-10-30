package database

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Auth struct {
	ID       string `json:"-" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func (auth *Auth) SignUp(db *gorm.DB) error {
	// select * from users where username="user"
	if err := db.Where(&Auth{Username: auth.Username}).First(auth).Error; err != nil {
		if err == gorm.ErrRecordNotFound { // jika tidak ditemukan
			if err := db.Create(auth).Error; err != nil { // membuat user baru
				return err
			}
		}
		return err
	}
	return nil
}

func (auth *Auth) Login(db *gorm.DB) (*Auth, error) {
	if err := db.Where(&Auth{Username: auth.Username, Password: auth.Password}).First(auth).Error; err != nil {
		if err == gorm.ErrRecordNotFound { // jika tidak ditemukan
			return nil, errors.Errorf("Incorect email or password")
		}
	}
	return auth, nil
}
