package main

import (
	"strings"

	"github.com/google/uuid"
)

// GetUUID returns a new uuid without hyphens
func GetUUID() string {
	var uuidWithHyphen = uuid.New()
	return strings.Replace(uuidWithHyphen.String(), "-", "", -1)
}
