package chatroutes

import (
	"database/sql"
	"net/http"

	chatHandler "github.com/Atharv7901/E2EEChatApp/chat-service/handler"
	"github.com/Atharv7901/E2EEChatApp/chat-service/middleware"
	"github.com/gin-gonic/gin"
)

type ChatServer struct {
	addr string
	db   *sql.DB
}

func NewChatServer(addr string, db *sql.DB) *ChatServer {
	return &ChatServer{
		addr: addr,
		db:   db,
	}
}

func (c *ChatServer) Run(chatHandler chatHandler.ChatHandlerInterface) error {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/users", middleware.AuthMiddleware(), chatHandler.GetAllUsers)
		api.GET("/users/:id", middleware.AuthMiddleware(), chatHandler.GetUserByID)
		api.POST("/chat/send", middleware.AuthMiddleware(), chatHandler.SendMessage)
		api.GET("/messages", middleware.AuthMiddleware(), chatHandler.FetchMessagesHandler)
	}

	return http.ListenAndServe(c.addr, r)
}
