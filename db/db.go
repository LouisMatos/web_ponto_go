package db

import (
	"log"
	"time"

	"github.com/LouisMatos/web_ponto_go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConexaoComBancoDados(connectionString string) {

	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados!")
	} else {
		dbConfig, _ := DB.DB()
		dbConfig.SetMaxOpenConns(25)
		dbConfig.SetMaxIdleConns(25)
		dbConfig.SetConnMaxLifetime(5 * time.Minute)
	}
	log.Println("Conexão ao banco de dados realizado com sucesso!")

}

func Migrate() {
	DB.AutoMigrate(&models.Produto{})
	log.Println("Migração do banco de dados concluída...")
}
