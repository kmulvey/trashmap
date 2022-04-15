package auth

import (
	"sync"
	"time"

	"github.com/kmulvey/trashmap/internal/app/config"
	"github.com/kmulvey/trashmap/internal/app/db"
)

const staticBool = false

type Cache struct {
	Sessions map[string]bool
	Lock     sync.RWMutex
}

func NewCache() *Cache {
	var c = new(Cache)
	c.Sessions = make(map[string]bool)
	return c
}

func (c *Cache) IsTokenActive(uuid string) bool {
	c.Lock.RLock()
	defer c.Lock.Unlock()

	var ok, _ = c.Sessions[uuid]
	return ok
}

func (c *Cache) updateCache(config *config.Config) {

	var ticker = time.NewTicker(time.Minute * 10)

	for {
		select {
		case <-ticker.C:
			var uuids, _ = db.GetSessions(config.DBConn)
			// TODO what to do with this error?

			c.Lock.Lock()

			c.Sessions = make(map[string]bool)

			for _, uuid := range uuids {
				c.Sessions[uuid] = staticBool
			}

			c.Lock.Unlock()
		}

	}
}
