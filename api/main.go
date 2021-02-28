package main

import (
	"log"

	"github.com/shien/weightrec-backend/pkg/server"
	"github.com/shien/weightrec-backend/pkg/user"
)

func main() {

	ur := user.Init()
	defer ur.Close()
	if err := ur.AddUser("testuser@gmail.com"); err != nil {
		log.Println(err)
	}

	result, err := ur.IsUserExists("testuser@gmail.com")
	log.Println(result, err)

	server.Init()
}
