package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/rishitashaw/bank-api/database/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}



func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)

	server.router=router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorRequest(err error) gin.H{
	return gin.H{"error": err.Error()}
}