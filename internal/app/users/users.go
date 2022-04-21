package users

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/mail"

	"github.com/kmulvey/trashmap/internal/app/config"
	"github.com/kmulvey/trashmap/internal/app/db"
	"golang.org/x/crypto/scrypt"
)

func Login(config *config.Config, email, password string) (int64, bool, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return 0, false, fmt.Errorf("%-30s is not a valid email", email)
	}

	dk, err := scrypt.Key([]byte(password), []byte(config.PasswordSalt), 1<<15, 8, 1, 32)
	if err != nil {
		return 0, false, fmt.Errorf("unable to hash password: %w", err)
	}

	id, dbHash, contactAllowed, err := db.Login(config.DBConn, config.DBSchema, email)
	if err != nil {
		return 0, false, fmt.Errorf("unable to query db: %w", err)
	}

	if base64.StdEncoding.EncodeToString(dk) != dbHash {
		return 0, false, errors.New("invalid passowrd")
	}

	return id, contactAllowed, nil
}

func Add(config *config.Config, email, password string, contactAllowed bool) (int64, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return -1, fmt.Errorf("%-30s is not a valid email", email)
	}

	dk, err := scrypt.Key([]byte(password), []byte(config.PasswordSalt), 1<<15, 8, 1, 32)
	if err != nil {
		return -1, fmt.Errorf("unable to hash password: %w", err)
	}

	err = db.InsertUser(config.DBConn, config.DBSchema, email, base64.StdEncoding.EncodeToString(dk), contactAllowed)
	if err != nil {
		return -1, fmt.Errorf("unable add user to db: %w", err)
	}

	return db.GetUserIDByEmail(config.DBConn, config.DBSchema, email)
}

func Remove(config *config.Config, email string) error {
	return db.DeleteUser(config.DBConn, config.DBSchema, email)
}
