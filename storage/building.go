package storage

import (
	"github.com/jmoiron/sqlx"
)

type BuildingStorage struct {
	Conn *sqlx.DB
}

func NewBuildingStorage(conn *sqlx.DB) *BuildingStorage {
	return &BuildingStorage{Conn: conn}
}

type NewBuilding struct {
	Id                   int64  `json:"id" db:"id"`
	Name                 string `json:"name" db:"name"`
	Owner                string `json:"owner" db:"owner"`
	Number_floors        int8   `json:"number_floors" db:"number_floors"`
	Number_floors_ground int8   `json:"number_floors_ground" db:"number_floors_ground"`
	Created_at           string `json:"created_at" db:"created_at"`
	Discarded            string `json:"discarded" db:"discarded" `
}

type NewFloor struct {
	Id               int64   `json:"id" db:"id"`
	Floor            string  `json:"floor" db:"floor"`
	Name             string  `json:"name" db:"name"`
	Area             string  `json:"area" db:"area"`
	Electric_measure float32 `json:"electric_measure" db:"electric_measure"`
	Water_measure    float32 `json:"water_measure" db:"water_measure"`
	Building_id      int64   `json:"building_id" db:"building_id"`
}

func (b *BuildingStorage) CreateNewBuilding(data NewBuilding) (int, error) {

	res, err := b.Conn.Exec(`
	INSERT INTO building (name, owner, number_floors, number_floors_ground, created_at) 
	values (?, ?, ?, ?, ?)`, data.Name, data.Owner, data.Number_floors, data.Number_floors_ground, data.Created_at)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
