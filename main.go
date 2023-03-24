package main

import (
	"github.com/LouisMatos/web_ponto_go/config"
	"github.com/LouisMatos/web_ponto_go/db"
	"github.com/LouisMatos/web_ponto_go/middlewares"
	"github.com/LouisMatos/web_ponto_go/routes"
	"github.com/LouisMatos/web_ponto_go/server"
)

func main() {

	config.LoadAppConfig()

	db.ConexaoComBancoDados(config.AppConfig.ConnectionString)

	db.Migrate()

	app := server.StartNewServer()

	middlewares.StartLogger(app)

	routes.CarregaRotas(app)

	server.RunNewServer(app, ":3000")

}
