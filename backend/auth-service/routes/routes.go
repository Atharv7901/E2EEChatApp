package authroutes

import (
	"database/sql"
	"net/http"

	authHandler "github.com/Atharv7901/E2EEChatApp/handler"
	"github.com/gin-gonic/gin"
)

type AuthServer struct {
	addr string
	db   *sql.DB
}

func NewAuthServer(addr string, db *sql.DB) *AuthServer {
	return &AuthServer{
		addr: addr,
		db:   db,
	}
}

func (a *AuthServer) Run(authHandler authHandler.AuthHandlerInterface) error {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/signup", authHandler.SignUp)
	}

	return http.ListenAndServe(a.addr, r)
}
