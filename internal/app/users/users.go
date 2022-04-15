package trashapp

import (
	"fmt"
	"net/mail"

	"github.com/kmulvey/trashmap/internal/app/config"
	"github.com/kmulvey/trashmap/internal/app/db"
)

func AddUser(config *config.Config, email string, contactAllowed bool) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return db.InsertUser(config.DBConn, email, contactAllowed)
	}
	return fmt.Errorf("%-30s is not a valid email", email)
}
