package server

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"

	helpers "github.com/jamesdury/attachments/internal/helpers"
)


func Setup(
	dirTemplate embed.FS,
	dirStatic embed.FS,
) *fiber.App  {
	engine := html.NewFileSystem(http.FS(dirTemplate), ".html")
	engine.AddFunc("filetype", helpers.FileType)
	engine.AddFunc("truncate", helpers.Truncate)
	engine.AddFunc("contact", helpers.Contact)
	engine.AddFunc("email", helpers.Email)
	engine.AddFunc("gravatar", helpers.Gravatar)
	engine.AddFunc("prettydate", helpers.PrettyDate)
	engine.AddFunc("bytesize", helpers.Bytesize)
	engine.AddFunc("escape", helpers.Escape)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use("/static", filesystem.New(filesystem.Config{
		Browse:     true,
		Root:       http.FS(dirStatic),
		PathPrefix: "static/dist",
	}))

	return app
}
