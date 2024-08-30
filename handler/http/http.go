package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type httpHandler struct {
	app *fiber.App
}

func New() *httpHandler {
	app := fiber.New()

	return &httpHandler{
		app: app,
	}
}

func (h *httpHandler) Start(port string) {
	log.Printf("Starting at port %v ...\n", port)
	log.Fatal(h.app.Listen(":" + port))
}
