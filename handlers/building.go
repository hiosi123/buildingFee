package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hiosi123/buildingFee/storage"
)

type BuildingHandler struct {
	Storage *storage.BuildingStorage
}

func NewBuildingHandler(storage *storage.BuildingStorage) *BuildingHandler {
	return &BuildingHandler{Storage: storage}
}

func (b *BuildingHandler) CreateCharge(c *fiber.Ctx) error {
	var charge storage.Charge

	err := c.BodyParser(&charge)
	if err != nil {
		return err
	}

	fmt.Println(charge)

	id, err := b.Storage.CreateNewCharge(charge)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"success": true, "id": id})
}
