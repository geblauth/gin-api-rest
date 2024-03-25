package database

import (
	"log"

	"github.com/geblauth/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBandoDeDados() {

	stringConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConexao))
	if err != nil {
		log.Panic("Erro para conectar no Banco")
	}
	DB.AutoMigrate(&models.Aluno{})
}
