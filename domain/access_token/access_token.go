package access_token

import "time"

const (
	expireTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(expireTime * time.Hour).Unix(),
	}
}

func (a *AccessToken) IsExpred() bool {

	return time.Unix(a.Expires, 0).Before(time.Now().UTC())
}
