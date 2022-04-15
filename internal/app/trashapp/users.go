package trashapp

import (
	"fmt"
	"net/mail"
)

func AddUser(config *Config, email string, contactAllowed bool) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return InsertUser(config.DBConn, email, contactAllowed)
	}
	return fmt.Errorf("%-30s is not a valid email", email)
}
