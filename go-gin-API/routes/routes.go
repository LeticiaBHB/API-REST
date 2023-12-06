package routes

import (
	"github.com/LeticiaBHB/Go-gin-API/controllers"
	"github.com/gin-gonic/gin"
)
func HandleRequests(){
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.Run(":5000") //posso especificar aqui a porta que vai abrir no localhost
}