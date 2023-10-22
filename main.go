package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hiosi123/buildingFee/config"
	"github.com/hiosi123/buildingFee/db"
	"github.com/hiosi123/buildingFee/handlers"
	"github.com/hiosi123/buildingFee/storage"
	"go.uber.org/fx"

	_ "github.com/go-sql-driver/mysql"
)

func newFiberServer(lc fx.Lifecycle, buildingHandlers *handlers.BuildingHandler) *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	buildingGroup := app.Group("/buildings")
	buildingGroup.Get("/building/:building_id", buildingHandlers.GetBuilding)
	buildingGroup.Get("/floor/:floor_id", buildingHandlers.GetFloor)
	buildingGroup.Get("/charge/:charge_id", buildingHandlers.GetCharge)
	buildingGroup.Get("/chargeList", buildingHandlers.GetChargeListByDate)
	buildingGroup.Post("/building", buildingHandlers.CreateBuilding)
	buildingGroup.Post("/floor", buildingHandlers.CreateFloor)
	buildingGroup.Post("/charge", buildingHandlers.CreateCharge)
	buildingGroup.Patch("/building", buildingHandlers.UpdateBuilding)
	buildingGroup.Patch("/floor", buildingHandlers.UpdateFloor)
	buildingGroup.Patch("/charge", buildingHandlers.UpdateCharge)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting fiber server on port 8080")
			go app.Listen(":8080")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}

func main() {
	fx.New(
		fx.Provide(
			// create: configEnvVars
			config.LoadEnv,
			// create: *sqlx.DB
			db.CreateMySqlConnection,
			// create: storage
			storage.NewBuildingStorage,
			// create: handler
			handlers.NewBuildingHandler,
		),
		fx.Invoke(newFiberServer),
	).Run()
}
