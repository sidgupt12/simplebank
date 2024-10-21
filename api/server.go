package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/sidgupt12/simplebank/db/sqlc"
)

// this server serves all HTTP requests of our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// Create a new server instance and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}
