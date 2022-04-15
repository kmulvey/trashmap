package auth

import (
	"github.com/kmulvey/trashmap/internal/app/config"
	"github.com/kmulvey/trashmap/internal/app/db"
)

func Login(config *config.Config, email string) (string, error) {
	var id, err = db.GetUserIDByEmail(config.DBConn, email)
	if err != nil {
		return "", err
	}

	var uuid = GetUUID()
	err = db.Login(config.DBConn, id, uuid)

	return uuid, err
}
