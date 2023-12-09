package main

import (
	"github.com/LeticiaBHB/Go-gin-API/database"
	"github.com/LeticiaBHB/Go-gin-API/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}

//linhas de código para subir um servidor