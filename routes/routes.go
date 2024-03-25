package routes

import (
	"github.com/geblauth/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/alunos", controllers.RetornaAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.RetornaAlunoId)
	r.DELETE("/alunos/:id", controllers.DeletaAlunos)
	r.PATCH("/alunos/:id", controllers.AtualizaAlunos)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoCPF)
	r.GET("/index", controllers.RenderizaPaginaIndex)
	r.NoRoute(controllers.RotaNaoEncontrada)
	r.Run()
}
