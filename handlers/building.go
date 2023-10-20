package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hiosi123/buildingFee/storage"
)

type BuildingHandler struct {
	Storage *storage.BuildingStorage
}

func NewBuildingHandler(storage *storage.BuildingStorage) *BuildingHandler {
	return &BuildingHandler{Storage: storage}
}

func (b *BuildingHandler) GetBuilding(c *fiber.Ctx) error {
	building_id := c.Params("building_id")

	id, err := strconv.ParseInt(building_id, 10, 64)
	if err != nil {
		return err
	}

	building, err := b.Storage.GetBuilding(id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"success": true, "building": building})
}

func (b *BuildingHandler) GetFloor(c *fiber.Ctx) error {
	floor_id := c.Params("floor_id")

	id, err := strconv.ParseInt(floor_id, 10, 64)
	if err != nil {
		return err
	}

	floor, err := b.Storage.GetFloor(id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"success": true, "building": floor})
}

func (b *BuildingHandler) GetCharge(c *fiber.Ctx) error {
	charge_id := c.Params("charge_id")

	id, err := strconv.ParseInt(charge_id, 10, 64)
	if err != nil {
		return err
	}

	charge, err := b.Storage.GetCharge(id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"success": true, "building": charge})
}

func (b *BuildingHandler) CreateBuilding(c *fiber.Ctx) error {
	var building storage.Building

	err := c.BodyParser(&building)
	if err != nil {
		return err
	}

	time := GetCurrentTime()
	building.Created_at = &time

	id, err := b.Storage.CreateNewBuilding(building)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"success": true, "id": id})
}

func (b *BuildingHandler) CreateFloor(c *fiber.Ctx) error {
	var floor storage.Floor

	err := c.BodyParser(&floor)
	if err != nil {
		return err
	}

	time := GetCurrentTime()
	floor.Created_at = &time

	id, err := b.Storage.CreateNewFloor(floor)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"success": true, "id": id})
}

func (b *BuildingHandler) CreateCharge(c *fiber.Ctx) error {
	var charge storage.Charge

	err := c.BodyParser(&charge)
	if err != nil {
		return err
	}

	time := GetCurrentTime()
	charge.Created_at = &time

	id, err := b.Storage.CreateNewCharge(charge)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"success": true, "id": id})
}

func (b *BuildingHandler) UpdateBuilding(c *fiber.Ctx) error {
	var building storage.Building

	err := c.BodyParser(&building)
	if err != nil {
		return err
	}

	err = b.Storage.UpdateBuilding(building)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"success": true})
}

func (b *BuildingHandler) UpdateFloor(c *fiber.Ctx) error {
	var floor storage.Floor

	err := c.BodyParser(&floor)
	if err != nil {
		return err
	}

	err = b.Storage.UpdateFloor(floor)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"success": true})
}

func (b *BuildingHandler) UpdateCharge(c *fiber.Ctx) error {
	var charge storage.Charge

	err := c.BodyParser(&charge)
	if err != nil {
		return err
	}

	err = b.Storage.UpdateCharge(charge)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"success": true})
}
