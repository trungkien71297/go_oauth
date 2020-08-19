package db

import (
	"github.com/trungkien71297/go_oauth/domain/access_token"
	"github.com/trungkien71297/go_oauth/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
}
type dbRepository struct{}

func New() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestError) {
	return nil, &errors.RestError{
		Code:    202,
		Message: "Database not implemented",
		Error:   "Database error",
	}
}
