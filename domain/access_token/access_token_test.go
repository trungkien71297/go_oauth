package access_token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpiredTime(t *testing.T) {
	if expireTime != 24 {
		t.Error("expireTime must be 24h")
	}
	assert.EqualValues(t, 24, expireTime, "expireTime must be 24h")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpred(), "New Access token cannnot be expired")
	assert.EqualValues(t, "", at.AccessToken, "Shold not have")
	assert.True(t, at.UserId == 0, "Shold not have")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpred(), "New Access token cannnot be expired")
}
