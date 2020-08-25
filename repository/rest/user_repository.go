package rest

import (
	"encoding/json"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/trungkien71297/go_api_management/domain/users"
	"github.com/trungkien71297/go_oauth/utils/errors"
	"time"
)

var (
	restClient = rest.RequestBuilder{
		BaseURL: "",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestError)
}
type userRepository struct{}

func NewUserRepository() RestUserRepository {
	return &userRepository{}
}
func (u *userRepository) LoginUser(username string, passowrd string) (*users.User, *errors.RestError) {
	rest := users.UserLogin{
		Username: username,
		Password: passowrd,
	}

	response := restClient.Post("/users/login", rest)
	if response == nil || response.Response == nil {
		return nil, &errors.RestError{
			Error:   "",
			Message: "",
			Code:    131,
		}
	}

	if response.StatusCode > 299 {
		return nil, &errors.RestError{
			Code:    299,
			Message: "Chay tung tang",
			Error:   "Nguoi ra trung ra bac vo nam",
		}
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, &errors.RestError{
			Code:    299,
			Message: "Chay tung tang",
			Error:   "Nguoi ra trung ra bac vo nam",
		}
	}
	return &user, nil
}
