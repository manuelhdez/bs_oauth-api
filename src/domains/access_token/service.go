package access_token

import (
	"strings"

	"github.com/manuelhdez/bs_oauth-api/src/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(id string) (*AccessToken, *errors.RestErr) {
	id = strings.TrimSpace(id)
	if len(id) == 0 {
		return nil, errors.NewBadRequestError("invalid access token")
	}

	accessToken, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
