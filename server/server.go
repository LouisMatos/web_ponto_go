package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/template/html"
)

func StartNewServer() *fiber.App {

	engine := html.NewFileSystem(http.Dir("./views"), ".html")

	engine.Reload(false)

	engine.Debug(false)

	engine.Layout("embed")

	engine.Delims("{{", "}}")

	app := fiber.New(fiber.Config{
		Prefork:           false,
		StrictRouting:     true,
		ServerHeader:      "Fiber",
		AppName:           "Web Ponto Eletronico App - v1.0.0",
		Views:             engine,
		ReduceMemoryUsage: false,
	})

	// Or extend your config for customization
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
		Realm: "Forbidden",
		Authorizer: func(user, pass string) bool {
			if user == "john" && pass == "doe" {
				return true
			}
			if user == "admin" && pass == "123456" {
				return true
			}
			return false
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Render("reload", nil, "layouts/main")
		},
		ContextUsername: "_user",
		ContextPassword: "_pass",
	}))

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
