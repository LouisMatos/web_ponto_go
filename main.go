package main

import (
	"log"
	"net/http"

	"github.com/LouisMatos/web_ponto_go/models"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/template/html"
)

func main() {

	engine := html.NewFileSystem(http.Dir("./views"), ".html")

	engine.Reload(true)

	engine.Debug(false)

	engine.Layout("embed")

	engine.Delims("{{", "}}")

	engine.AddFunc("greet", func(name string) string {
		return "Hello, " + name + "!"
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./views/resources/")

	app.Get("/", func(c *fiber.Ctx) error {
		todosOsProdutos := models.BuscaTodosOsProdutos()
		return c.Render("index", todosOsProdutos)
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
		models.DeletaProduto(c.Query("id"))
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

		models.CriaNovoProduto(p.Nome, p.Descricao, p.Preco, p.Quantidade)
		return c.Redirect("/")
	})

	app.Get("/edit", func(c *fiber.Ctx) error {
		editProduto := models.EditaProduto((c.Query("id")))

		return c.Render("edit", editProduto)
	})

	app.Post("/update", func(c *fiber.Ctx) error {

		p := new(models.Produto)
		if err := c.BodyParser(p); err != nil {
			return err
		}

		models.AtualizaProduto(p.Id, p.Nome, p.Descricao, p.Preco, p.Quantidade)
		return c.Redirect("/")
	})

	log.Fatal(app.Listen(":3000"))
}
