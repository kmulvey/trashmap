package auth

import (
	"github.com/kmulvey/trashmap/internal/app/config"
	"github.com/kmulvey/trashmap/internal/app/db"
)

func ExpireOldSessions(config *config.Config) error {
	return db.ExpireSessions(config.DBConn)
}
