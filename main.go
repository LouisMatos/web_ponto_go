package main

import (
	"fmt"
	"net/http"

	"github.com/LouisMatos/web_ponto_go/routes"
)

func main() {
	fmt.Println("Run!!!!")
	routes.CarregaRotas()
	fmt.Println("Run!!!!")
	http.ListenAndServe(":8000", nil)
}
