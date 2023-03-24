package service

import (
	"github.com/LouisMatos/web_ponto_go/db"
	"github.com/LouisMatos/web_ponto_go/models"
)

func BuscaTodosOsProdutos() []models.Produto {

	produtos := []models.Produto{}

	db.DB.Find(&produtos)

	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {

	produto := models.Produto{
		Nome:       nome,
		Descricao:  descricao,
		Preco:      preco,
		Quantidade: quantidade,
	}

	db.DB.Create(&produto)

}

func DeletaProduto(id string) {

	var produto models.Produto

	db.DB.Delete(&produto, id)

}

func EditaProduto(id string) models.Produto {

	var produto models.Produto

	db.DB.First(&produto, id)

	return produto
}

func AtualizaProduto(id uint, nome, descricao string, preco float64, quantidade int) {

	produto := models.Produto{
		Id:         id,
		Nome:       nome,
		Descricao:  descricao,
		Preco:      preco,
		Quantidade: quantidade,
	}

	db.DB.Save(&produto)
}
