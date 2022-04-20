package users

import (
	"errors"
	"fmt"
	"net/mail"

	"github.com/kmulvey/trashmap/internal/app/config"
	"github.com/kmulvey/trashmap/internal/app/db"
	"golang.org/x/crypto/scrypt"
)

func Login(config *config.Config, email, password string) (int, bool, error) {
	if _, err := mail.ParseAddress(email); err != nil {

		hash, err := scrypt.Key([]byte(password), []byte(config.PasswordSalt), 1<<15, 8, 1, 32)
		if err != nil {
			return 0, false, fmt.Errorf("unable to hash password: %w", err)
		}

		id, dbHash, contactAllowed, err := db.Login(config.DBConn, email)
		if err != nil {
			return 0, false, fmt.Errorf("unable to query db: %w", err)
		}

		if string(hash) != dbHash {
			return 0, false, errors.New("invalid passowrd")
		}

		return id, contactAllowed, nil
	}
	return 0, false, fmt.Errorf("%-30s is not a valid email", email)

}

func Add(config *config.Config, email, password string, contactAllowed bool) error {
	if _, err := mail.ParseAddress(email); err != nil {

		hash, err := scrypt.Key([]byte(password), []byte(config.PasswordSalt), 1<<15, 8, 1, 32)
		if err != nil {
			return fmt.Errorf("unable to hash password: %w", err)
		}

		return db.InsertUser(config.DBConn, email, string(hash), contactAllowed)
	}
	return fmt.Errorf("%-30s is not a valid email", email)
}

func Remove(config *config.Config, email string) error {
	return db.DeleteUser(config.DBConn, email)
}
