package models

import "gopkg.in/validator.v2"

type User struct {
	Id         uint    `gorm:"primary_key" json:"id"`
	Nome       string  `gorm:"not null" json:"nome"`
	Descricao  string  `gorm:"not null" json:"descricao"`
	Preco      float64 `gorm:"not null" json:"preco"`
	Quantidade int     `gorm:"not null" json:"quantidade"`
}

type UserDTO struct {
	Descricao string `json:"descricao" validate:"nonzero, nonnil"`
	Valor     string `json:"valor" validate:"nonzero, min=4, nonnil, regexp=^[\\d]+[.][\\d]{2}$"`
	Data      string `json:"data" validate:"nonzero, min=10, nonnil, regexp=^([0]?[1-9]|[1|2][0-9]|[3][0|1])[./-]([0]?[1-9]|[1][0-2])[./-]([0-9]{4}|[0-9]{2})$"`
	Categoria string `json:"categoria"`
}

func ValidateUserData(user *UserDTO) error {
	if err := validator.Validate(user); err != nil {
		return err
	}
	return nil
}
