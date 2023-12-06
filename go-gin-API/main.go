package main

import (
	"github.com/LeticiaBHB/Go-gin-API/models"
	"github.com/LeticiaBHB/Go-gin-API/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Joana", CPF:"0000000000", RG: "1234567890"},
		{Nome: "Ana", CPF:"11111111111", RG: "0987654321"},
	}
	routes.HandleRequests()
}

//linhas de código para subir um servidor