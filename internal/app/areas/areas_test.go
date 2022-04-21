package areas

import (
	"context"
	"testing"

	"github.com/kmulvey/trashmap/internal/app/config"
	"github.com/kmulvey/trashmap/internal/app/users"
	"github.com/kmulvey/trashmap/internal/pkg/gps"
	"github.com/stretchr/testify/assert"
)

func TestAreaFlow(t *testing.T) {
	t.Parallel()

	var schema = "testareaflow"

	var config, err = config.NewTestConfig(schema)
	assert.NoError(t, err)

	addUserID, err := users.Add(config, "testareaflow@site.com", "password", true)
	assert.NoError(t, err)
	assert.True(t, addUserID > 0)

	boulder, err := gps.NewAreaFromPostGISString("40.259822802779816 -105.65290936674407,40.26201885227386 -105.05519389236237,39.95833557541779 -105.05494458234654,39.93788639054093 -105.68899269947714,40.259822802779816 -105.65290936674407")
	assert.NoError(t, err)
	areaID, err := SaveArea(config, addUserID, boulder)
	assert.NoError(t, err)
	assert.True(t, areaID > 0)

	colorado, err := gps.NewAreaFromPostGISString("40.99837454922104 -109.05421673800787, 41.0004139327614 -102.05758033526878, 37.015379641011805 -109.05244682441632, 37.004114233165524 -102.04615148062369, 40.99837454922104 -109.05421673800787")
	assert.NoError(t, err)
	coloradoMap, err := GetPickupAreasWithinArea(config, colorado)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(coloradoMap))

	err = RemoveArea(config, addUserID, boulder)
	assert.NoError(t, err)

	err = users.Remove(config, "testareaflow@site.com")
	assert.NoError(t, err)

	_, err = config.DBConn.Exec(context.Background(), "drop schema "+schema+" cascade")
	assert.NoError(t, err)
}
