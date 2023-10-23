package handlers

import (
	"errors"
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

func (b *BuildingHandler) GetChargeListByDate(c *fiber.Ctx) error {
	year := c.Query("year")
	month := c.Query("month")

	if len(year) > 4 || len(month) > 2 {
		return errors.New("year should be 4 and month should be 2")
	}

	chargeList, err := b.Storage.GetChargeListByDate(year, month)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"success": true, "chargeList": chargeList})
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

	//중복 x
	currentCharge, err := b.Storage.GetChargeByInfo(*charge.Year, *charge.Month, *charge.Floor_id, *charge.Measure_number)
	if err != nil {
		return err
	}
	if currentCharge.Id != nil && *currentCharge.Id != 0 {
		return errors.New("current charge already exist")
	}

	year := *charge.Year
	currentMonth, err := strconv.Atoi(*charge.Month)
	if err != nil {
		return err
	}

	lm := currentMonth - 1
	lastMonth := strconv.Itoa(lm)
	if len(lastMonth) == 1 {
		lastMonth = "0" + lastMonth
	} else if lastMonth == "00" {
		lastMonth = "12"
		yearInt, err := strconv.Atoi(year)
		if err != nil {
			return err
		}
		year = strconv.Itoa(yearInt - 1)
	}

	lastCharge, err := b.Storage.GetChargeByInfo(year, lastMonth, *charge.Floor_id, *charge.Measure_number)
	if err != nil {
		return err
	}

	time := GetCurrentTime()
	charge.Created_at = &time

	if lastCharge.Id != nil {
		ed := *charge.Electric_measure - *lastCharge.Electric_measure
		wd := *charge.Water_measure - *lastCharge.Water_measure
		charge.Electric_difference = &ed
		charge.Water_difference = &wd
	}

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
