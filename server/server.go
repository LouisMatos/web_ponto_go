package server

import (
	"log"
	"net/http"
	"time"

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
		Prefork:       true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Web Ponto Eletronico App - v1.0.0",
		Views:         engine,
	})

	//app := fiber.New(fiber.Config{
	//	Views: engine,
	//})

	config := fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Index:         "index.html",
		CacheDuration: 60 * time.Second,
		MaxAge:        3600,
		Next:          nil,
	}
	app.Static("/static", "./views/resources/", config)

	return app
}

func RunNewServer(app *fiber.App, port string) {
	log.Fatal(app.Listen(":3000"))
}
