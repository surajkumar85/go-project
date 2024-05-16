package main

import (
	"fmt"
	"log"

	"github.com/surajkumar85/go-project/server"
)

func main() {
	fmt.Println("This is the entry file of this project")

	server := server.NewAPIServer(":8000")

	if err :=server.Run(); err != nil {
		log.Fatal("Somthing went wrong with the server")
	}

}