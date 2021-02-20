package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Def *gin.Engine
}

func NewServer() Server {
	s := Server{Def: gin.Default()}
	return s
}
