// Code generated by protoc-gen-gin. DO NOT EDIT.

package test

import (
	gin "github.com/gin-gonic/gin"
	http "net/http"
)

type HelloHttpServer struct {
	server HelloServer
	router gin.IRouter
}

func RegisterHelloHttpServer(srv HelloServer, r gin.IRouter) {
	s := HelloHttpServer{
		server: srv,
		router: r,
	}

	s.RegisterHttpService()
}

func (s *HelloHttpServer) SayHello(c *gin.Context) {
	var in SayHelloRequest

	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	out, err := s.server.SayHello(c, &in)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, out)
}

func (s *HelloHttpServer) SayGoodbye(c *gin.Context) {
	var in SayGoodbyeRequest

	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	out, err := s.server.SayGoodbye(c, &in)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, out)
}

func (s *HelloHttpServer) RegisterHttpService() {
	s.router.Handle("POST", "/v1/SayHello", s.SayHello)
	s.router.Handle("POST", "/v1/SayGoodbye", s.SayGoodbye)
}
