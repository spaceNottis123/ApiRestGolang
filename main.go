package main

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Print("Starting server")
	routes.HandleRequest()
}
