package app

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"

	notmuch "github.com/zenhack/go.notmuch"

	home "github.com/jamesdury/attachments/app/home"
	email "github.com/jamesdury/attachments/pkg/notmuch"
)

//go:embed static/template/*
var embedDirTemplate embed.FS

//go:embed static/dist/*
var embedDirStatic embed.FS

func Build(db *notmuch.DB) *fiber.App  {
	engine := html.NewFileSystem(http.FS(embedDirTemplate), ".html")
	engine.AddFunc("filetype", TemplateFunctionFileType)
	engine.AddFunc("truncate", TemplateFunctionTruncate)
	engine.AddFunc("contact", TemplateFunctionContact)
	engine.AddFunc("email", TemplateFunctionEmail)
	engine.AddFunc("gravatar", TemplateFunctionGravatar)
	engine.AddFunc("prettydate", TemplateFunctionPrettyDate)
	engine.AddFunc("bytesize", TemplateFunctionBytesize)
	engine.AddFunc("escape", TemplateFunctionEscape)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use("/static", filesystem.New(filesystem.Config{
		Browse:     true,
		Root:       http.FS(embedDirStatic),
		PathPrefix: "static/dist",
	}))

	emailService := email.NewService(email.NewRepo(db))

	home.Router(app.Group("/"), emailService)

	return app
}
