package db

import (
	cassandra_db "github.com/manuelhdez/bs_oauth-api/src/clients/cassandra"
	"github.com/manuelhdez/bs_oauth-api/src/domains/access_token"
	"github.com/manuelhdez/bs_oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra_db.GetSession()
	if err != nil {
		panic(err.Error())
	}
	defer session.Close()

	return nil, errors.NewInternalServerError("database connection not implemented")
}
