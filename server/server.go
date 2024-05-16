package server

import (
	"fmt"
	"net/http"

	"github.com/surajkumar85/go-project/internals/user"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	userService := user.NewService()
	userHandler:=user.NewUserHandler(userService)
	userHandler.RegisteredRoutes(router)
	v1 := http.NewServeMux()
	v1.Handle("/api/v1/" ,http.StripPrefix("/api/v1",router))


	server := http.Server{
		Addr: s.addr,
		Handler: v1,
	}

	fmt.Println("Our server is running on port: ",s.addr)

	return server.ListenAndServe() 
}