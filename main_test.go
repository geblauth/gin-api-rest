package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/geblauth/gin-api-rest/controllers"
	"github.com/geblauth/gin-api-rest/database"
	"github.com/geblauth/gin-api-rest/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func setupRotasTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	rotas := gin.Default()
	return rotas

}

func TestVerificaStatusCodeSaudacao(t *testing.T) {
	r := setupRotasTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/Germano", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais!")
	mockDaResposta := `{"API diz:":"Eai Germano, Tudo Beleza?"}`
	respostBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostBody))
}

func TestListandoAlunos(t *testing.T) {
	database.ConectaBandoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := setupRotasTeste()
	r.GET("/alunos", controllers.RetornaAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)

}
func TestBuscaCPF(t *testing.T) {
	database.ConectaBandoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := setupRotasTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/67867867867", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)

}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Aluno Teste", CPF: "67867867867", RG: "222121212"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}
func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)

}

func TestBuscaAlunoId(t *testing.T) {
	database.ConectaBandoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := setupRotasTeste()
	r.GET("/alunos/:id", controllers.RetornaAlunoId)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Aluno Teste", alunoMock.Nome)
	assert.Equal(t, "67867867867", alunoMock.CPF)
	assert.Equal(t, "222121212", alunoMock.RG)
}

func TestAlunoDelete(t *testing.T) {
	database.ConectaBandoDeDados()
	CriaAlunoMock()
	r := setupRotasTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAlunos)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)

}

func TestAtualiza(t *testing.T) {
	database.ConectaBandoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := setupRotasTeste()
	r.PATCH("/alunos/:id", controllers.AtualizaAlunos)
	aluno := models.Aluno{Nome: "Aluno Teste", CPF: "12345678912", RG: "123456000"}
	valorJson, _ := json.Marshal(aluno)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, "12345678912", alunoMockAtualizado.CPF)
	assert.Equal(t, "123456000", alunoMockAtualizado.RG)
	assert.Equal(t, "Aluno Teste", alunoMockAtualizado.Nome)

}
