package routes

import (
	"github.com/LouisMatos/web_ponto_go/models"
	"github.com/LouisMatos/web_ponto_go/service"
	"github.com/gofiber/fiber/v2"
)

func CarregaRotas(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	app.Get("/index-old", func(c *fiber.Ctx) error {
		todosOsProdutos := service.BuscaTodosOsProdutos()
		return c.Render("index-old", todosOsProdutos)
	})

	app.Get("/layout", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "login/login")
	})

	app.Get("/delete", func(c *fiber.Ctx) error {
		service.DeletaProduto(c.Query("id"))
		return c.Redirect("/")
	})

	app.Get("/new", func(c *fiber.Ctx) error {
		return c.Render("new", nil)
	})

	app.Post("/insert", func(c *fiber.Ctx) error {

		p := new(models.Produto)
		if err := c.BodyParser(p); err != nil {
			return err
		}

		service.CriaNovoProduto(p.Nome, p.Descricao, p.Preco, p.Quantidade)
		return c.Redirect("/")
	})

	app.Get("/edit", func(c *fiber.Ctx) error {
		editProduto := service.EditaProduto((c.Query("id")))

		return c.Render("edit", editProduto, "layouts/main")
	})

	app.Post("/update", func(c *fiber.Ctx) error {

		p := new(models.Produto)
		if err := c.BodyParser(p); err != nil {
			return err
		}

		service.AtualizaProduto(p.Id, p.Nome, p.Descricao, p.Preco, p.Quantidade)
		return c.Redirect("/")
	})

	// a custom 404 handler at router tail
	//app.ErrorHandler(func(c *fiber.Ctx) error {
	//	c.Status(404).Redirect("")
	//	return nil
	//})
}
