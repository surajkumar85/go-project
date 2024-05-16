package server

import (
	"fmt"
	"net/http"
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

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/" ,http.StripPrefix("/api/v1",router))

	server := http.Server{
		Addr: s.addr,
		Handler: v1,
	}

	fmt.Println("Our server is running on port: ",s.addr)

	return server.ListenAndServe() 
}