package controllers

import (
	"net/http"

	"github.com/LeticiaBHB/Go-gin-API/database"
	"github.com/LeticiaBHB/Go-gin-API/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", beleza ?",
	})
}

func CriaNovoAluno(c*gin.Context){
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
			return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}