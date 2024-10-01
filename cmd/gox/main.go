package main

import (
	"fmt"
	"gox/api"
	"gox/config"
	"net/http"
)

func main() {

	addr := config.LoadConfig().ServerAddress

	http.HandleFunc("/", api.HomeHandler)
	http.HandleFunc("/users", api.GetAllUsersHandler)
	http.HandleFunc("/user", api.GetUserByIdHandler)

	// SERVER START
	fmt.Println("server is running on port", addr)

	err := http.ListenAndServe(addr, nil)

	if err != nil {
		fmt.Println("failed to start server", err)
	}

}
