package db

import (
	"github.com/gocql/gocql"
	"github.com/trungkien71297/go_oauth/clients/cassandra"
	"github.com/trungkien71297/go_oauth/domain/access_token"
	"github.com/trungkien71297/go_oauth/utils/errors"
)

const (
	getAccessToken    = "Select access_token, user_id, client_id, expires from oauth.access_tokens where access_token=?;"
	createAccessToken = "insert into access_tokens(access_token, user_id, client_id, expires) values(?,?,?,?);"
	updateAccessToken = "update access_tokens set expires=? where access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken) *errors.RestError
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}
type dbRepository struct{}

func New() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	session := cassandra.GetSession()
	var result access_token.AccessToken
	if err := session.Query(getAccessToken, id).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, &errors.RestError{
				Code:    123,
				Message: "Not found Nigma",
				Error:   "OG",
			}
		}
		return nil, &errors.RestError{
			Code:    111,
			Message: "Db error",
			Error:   "Not valid",
		}
	}
	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestError {

	session := cassandra.GetSession()
	if err := session.Query(createAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec(); err != nil {
		return &errors.RestError{
			Code:    117,
			Message: "Insert error",
			Error:   "Not valid",
		}
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestError {
	session := cassandra.GetSession()
	if err := session.Query(updateAccessToken, at.Expires, at.AccessToken).Exec(); err != nil {
		return &errors.RestError{
			Code:    117,
			Message: "Insert error",
			Error:   "Not valid",
		}
	}

	return nil

}
