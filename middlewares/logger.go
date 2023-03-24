package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func StartLogger(app *fiber.App) {

	//"github.com/gofiber/fiber/v2/middleware/logger"
	//app.Use(logger.New())

	app.Use(logger.New(logger.Config{Next: nil,
		Done:         nil,
		Format:       "[${time}] | ${ip}:${port} | ${status} | ${latency} | ${method} | ${path} | ${body}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
	}))

}
