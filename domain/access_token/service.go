package access_token

import "github.com/trungkien71297/go_oauth/utils/errors"

type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
}

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestError)
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
