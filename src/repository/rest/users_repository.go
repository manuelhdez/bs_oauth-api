package rest

import (
	"encoding/json"
	"time"

	"github.com/manuelhdez/bs_oauth-api/src/domains/users"
	rest_errors "github.com/manuelhdez/bs_oauth-api/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *rest_errors.RestErr)
}

type usersRepository struct{}

func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *rest_errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response := usersRestClient.Post("/auth/login", request)

	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError("invalid restclient response when trying to login user")
	}

	if response.StatusCode > 299 {
		apiErr := rest_errors.NewRestErrorFromBytes(response.Bytes())
		if apiErr != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, apiErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal users login response")
	}
	return &user, nil
}
