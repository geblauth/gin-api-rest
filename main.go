package main

import (
	"github.com/geblauth/gin-api-rest/database"
	"github.com/geblauth/gin-api-rest/routes"
)

func main() {
	database.ConectaBandoDeDados()

	routes.HandleRequest()
}
