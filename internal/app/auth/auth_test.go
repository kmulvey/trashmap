package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var key = "thisis32bitlongpassphraseimusing"

func TestEncodeDecode(t *testing.T) {
	t.Parallel()

	var authenticator = NewAuthenticator(key)

	var token = NewToken("me@email.com")

	var ciphertext, err = authenticator.Encode(token)
	assert.NoError(t, err)

	decodedToken, err := authenticator.Decode(ciphertext)
	assert.NoError(t, err)

	assert.Equal(t, token.Email, decodedToken.Email)
	assert.True(t, time.Since(decodedToken.Issued) < time.Millisecond)
	var shouldExpire = time.Now().Add(time.Hour * 12)
	assert.True(t, shouldExpire.Sub(decodedToken.Expires) < time.Millisecond)
}
