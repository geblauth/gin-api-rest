package routes

import (
	"github.com/geblauth/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.RetornaAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.RetornaAlunoId)
	r.DELETE("/alunos/:id", controllers.DeletaAlunos)
	r.PATCH("/alunos/:id", controllers.AtualizaAlunos)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoCPF)
	r.Run()
}
