package gps

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {

	var m = make([]*Area, 2)
	m[0], _ = NewAreaFromPostGISString("40.259822802779816 -105.65290936674407,40.26201885227386 -105.05519389236237")
	m[1], _ = NewAreaFromPostGISString("39.95833557541779 -105.05494458234654,39.93788639054093 -105.68899269947714")

	var js, _ = json.Marshal(m)
	fmt.Println(string(js))
}
