package users

import (
	"context"
	"testing"

	"github.com/kmulvey/trashmap/internal/app/config"
	"github.com/stretchr/testify/assert"
)

func TestUserFlow(t *testing.T) {
	t.Parallel()

	var schema = "testuserflow"

	var config, err = config.NewTestConfig(schema)
	assert.NoError(t, err)

	addUserID, err := Add(config, "testuserflow@site.com", "password", true)
	assert.NoError(t, err)
	assert.True(t, addUserID > 0)

	loginUserID, contact, err := Login(config, "testuserflow@site.com", "password")
	assert.NoError(t, err)
	assert.Equal(t, addUserID, loginUserID)
	assert.True(t, contact)

	err = Remove(config, "testuserflow@site.com")
	assert.NoError(t, err)

	_, err = config.DBConn.Exec(context.Background(), "drop schema "+schema+" cascade")
	assert.NoError(t, err)
}
