package app

import (
	"github.com/gin-gonic/gin"
	cassandra_db "github.com/manuelhdez/bs_oauth-api/src/clients/cassandra"
	"github.com/manuelhdez/bs_oauth-api/src/domains/access_token"
	"github.com/manuelhdez/bs_oauth-api/src/http"
	"github.com/manuelhdez/bs_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, err := cassandra_db.GetSession()
	if err != nil {
		panic(err.Error())
	}
	session.Close()

	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8080")
}
