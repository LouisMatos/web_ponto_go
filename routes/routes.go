package routes

import (
	"time"

	"github.com/LouisMatos/web_ponto_go/models"
	"github.com/LouisMatos/web_ponto_go/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func CarregaRotas(app *fiber.App) {

	app.Get("/metrics", monitor.New(monitor.Config{
		Title:      "Fiber Monitor",
		Refresh:    3 * time.Second,
		APIOnly:    false,
		Next:       nil,
		CustomHead: "",
		FontURL:    "https://fonts.googleapis.com/css2?family=Roboto:wght@400;900&display=swap",
		ChartJsURL: "https://cdn.jsdelivr.net/npm/chart.js@2.9/dist/Chart.bundle.min.js",
	}))

	//web := app.Group("/web", func(c *fiber.Ctx) error { // middleware for /api/v1
	//	c.Set("Version", "v1")
	//	return c.Next()
	//})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil, "layouts/main")
	})

	app.Get("/index-old", func(c *fiber.Ctx) error {
		todosOsProdutos := service.BuscaTodosOsProdutos()
		return c.Render("index-old", todosOsProdutos)
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("index", nil, "login/login")
	})

	app.Get("/new", func(c *fiber.Ctx) error {
		return c.Render("new", nil, "layouts/main")
	})

	app.Get("/edit", func(c *fiber.Ctx) error {
		editProduto := service.EditaProduto((c.Query("id")))

		return c.Render("edit", editProduto, "layouts/main")
	})

	app.Post("/insert", func(c *fiber.Ctx) error {

		p := new(models.Produto)
		if err := c.BodyParser(p); err != nil {
			return err
		}

		service.CriaNovoProduto(p.Nome, p.Descricao, p.Preco, p.Quantidade)
		return c.Redirect("/")
	})

	app.Post("/update", func(c *fiber.Ctx) error {

		p := new(models.Produto)
		if err := c.BodyParser(p); err != nil {
			return err
		}

		service.AtualizaProduto(p.Id, p.Nome, p.Descricao, p.Preco, p.Quantidade)
		return c.Redirect("/")
	})

	app.Get("/delete", func(c *fiber.Ctx) error {
		service.DeletaProduto(c.Query("id"))
		return c.Redirect("/")
	})

}
