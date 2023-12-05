package main
import(
	"github.com/gin-gonic/gin"
)
func ExibeTodosAlunos(c *gin.Context){
	c.JSON(200,gin.H{
		"id":"1",
		"nome":"Maria",
	})
}

func main(){
	r := gin.Default()
	r.GET("/alunos", ExibeTodosAlunos)
	r.Run(":5000") //posso especificar aqui a porta que vai abrir no localhost
}
//linhas de código para subir um servidor