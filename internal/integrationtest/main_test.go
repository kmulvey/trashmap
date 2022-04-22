package integrationtest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/kmulvey/trashmap/internal/app/config"
	"github.com/kmulvey/trashmap/internal/app/webserver"
	"github.com/kmulvey/trashmap/internal/pkg/testhelpers"
	"github.com/stretchr/testify/assert"
)

// {"contact_allowed":true,"id":2}
type CreateUserResp struct {
	ContactAllowed bool `json:"contact_allowed"`
	ID             int  `json:"id"`
}

func TestHelloWorld(t *testing.T) {
	t.Parallel()

	// start webserver
	var config, err = config.NewTestConfig("schema")
	assert.NoError(t, err)

	go func() {
		err = webserver.StartWebServer(config)
		assert.NoError(t, err)
	}()
	time.Sleep(time.Second) // wait for webser to come up

	jar, err := cookiejar.New(nil)
	assert.NoError(t, err)

	var client = &http.Client{
		Jar: jar,
	}

	////////////////// CREATE AREA LOGGED OUT (should fail)
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8000/area", strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", "0")

	res, err := client.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusForbidden, res.StatusCode)

	////////////////// CREATE USER
	var loginCreds = url.Values{
		"email":           {testhelpers.RandomString(5) + "@email.com"},
		"password":        {"password"},
		"contact_allowed": {"true"},
	}

	var createURL = "http://localhost:8000/user"

	req, err = http.NewRequest(http.MethodPut, "http://localhost:8000/user", strings.NewReader(loginCreds.Encode())) // URL-encoded payload
	assert.NoError(t, err)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginCreds.Encode())))

	res, err = client.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)

	var createUserResp CreateUserResp
	err = json.Unmarshal(body, &createUserResp)
	assert.NoError(t, err)
	assert.True(t, createUserResp.ContactAllowed)
	assert.True(t, createUserResp.ID > 0)

	assert.NoError(t, res.Body.Close())

	u, err := url.Parse(createURL)
	assert.NoError(t, err)

	var cookies = jar.Cookies(u)
	assert.Equal(t, 1, len(cookies))
	assert.Equal(t, "web-session", cookies[0].Name)

	////////////////// ADD AREA
	areaJSON, err := testhelpers.GetRandomArea().ToJSON()
	assert.NoError(t, err)
	var areaData = url.Values{
		"user_id": {strconv.Itoa(createUserResp.ID)},
		"area":    {areaJSON},
	}
	req, err = http.NewRequest(http.MethodPut, "http://localhost:8000/area", strings.NewReader(areaData.Encode()))
	assert.NoError(t, err)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", "0")

	res, err = client.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, err = ioutil.ReadAll(res.Body)
	assert.NoError(t, err)

	createUserResp = CreateUserResp{} // this is lazy but the responses are similar enough
	err = json.Unmarshal(body, &createUserResp)
	assert.NoError(t, err)
	assert.True(t, createUserResp.ID > 0)

	assert.NoError(t, res.Body.Close())
}
