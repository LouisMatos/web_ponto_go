package server

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func StartNewServer() *fiber.App {

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

	return app
}

func RunNewServer(app *fiber.App, port string) {
	log.Fatal(app.Listen(":3000"))
}
