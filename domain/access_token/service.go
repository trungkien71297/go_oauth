package access_token

import (
	"strings"

	"github.com/trungkien71297/go_oauth/utils/errors"
)

type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpirationTime(AccessToken) *errors.RestError
}

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpirationTime(AccessToken) *errors.RestError
}
type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(access_token_id string) (*AccessToken, *errors.RestError) {
	accessToken, err := s.repository.GetById(access_token_id)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(at AccessToken) *errors.RestError {
	if len(strings.TrimSpace(at.AccessToken)) == 0 {
		return &errors.RestError{
			Code:    543,
			Message: "token not valid",
			Error:   "Cassiopia",
		}
	}
	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestError {
	return s.repository.UpdateExpirationTime(at)
}
