package storage

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type BuildingStorage struct {
	Conn *sqlx.DB
}

func NewBuildingStorage(conn *sqlx.DB) *BuildingStorage {
	return &BuildingStorage{Conn: conn}
}

type Building struct {
	Id                   int64  `json:"id" db:"id"`
	Name                 string `json:"name" db:"name"`
	Owner                string `json:"owner" db:"owner"`
	Number_floors        int8   `json:"number_floors" db:"number_floors"`
	Number_floors_ground int8   `json:"number_floors_ground" db:"number_floors_ground"`
	Created_at           string `json:"created_at" db:"created_at"`
	Discarded            string `json:"discarded" db:"discarded" `
}

type Floor struct {
	Id          int64  `json:"id" db:"id"`
	Floor       string `json:"floor" db:"floor"`
	Name        string `json:"name" db:"name"`
	Area        string `json:"area" db:"area"`
	Building_id int64  `json:"building_id" db:"building_id"`
}

type Charge struct {
	Id               int64   `json:"id" db:"id"`
	Year             string  `json:"year" db:"year"`
	Date             string  `json:"date" db:"date"`
	Electric_measure float32 `json:"electric_measure" db:"electric_measure"`
	Water_measure    float32 `json:"water_measure" db:"water_measure"`
	Floor_id         int64   `json:"floor_id" db:"floor_id"`
}

func (b *BuildingStorage) CreateNewBuilding(data Building) (int, error) {
	query, values, err := InsertQueryStr(data)
	if err != nil {
		return 0, err
	}

	res, err := b.Conn.Exec(fmt.Sprintf(`INSERT INTO building %s`, query), values...)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (b *BuildingStorage) CreateNewFloor(data Floor) (int, error) {
	query, values, err := InsertQueryStr(data)
	if err != nil {
		return 0, err
	}

	res, err := b.Conn.Exec(fmt.Sprintf(`INSERT INTO floor %s`, query), values...)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (b *BuildingStorage) GetFloor(Id int64) (Floor, error) {
	var floor Floor

	fQuery, err := GetPureFieldStr(Floor{})
	if err != nil {
		return Floor{}, err
	}

	query := fmt.Sprintf("SELECT %s FROM floor WHERE id = ?", fQuery)
	err = b.Conn.Get(&floor, query, Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return Floor{}, nil
		}
		return Floor{}, err
	}

	return floor, nil
}
