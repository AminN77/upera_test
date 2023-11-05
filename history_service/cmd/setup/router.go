package setup

import (
	"github.com/AminN77/upera_test/history_service/api/controller"
	fiberPkg "github.com/AminN77/upera_test/history_service/pkg/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetRouter(c *controller.Controller) *fiber.App {
	app := fiberPkg.NewFiberRouter()

	// middlewares
	app.Use(recover.New())

	// routes
	v1 := app.Group("/api/v1/history")
	{
		v1.Get("/:productId/revision", c.FetchRevisionsOfOneProduct)
		v1.Get("/:revisionNumber", c.FetchRevision)
	}

	return app
}
