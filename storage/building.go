package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type BuildingStorage struct {
	Conn *sqlx.DB
}

func NewBuildingStorage(conn *sqlx.DB) *BuildingStorage {
	return &BuildingStorage{Conn: conn}
}

type Building struct {
	Id                   *int64  `json:"id" db:"id"`
	Name                 *string `json:"name" db:"name"`
	Owner                *string `json:"owner" db:"owner"`
	Number_floors        *int8   `json:"number_floors" db:"number_floors"`
	Number_floors_ground *int8   `json:"number_floors_ground" db:"number_floors_ground"`
	Built_at             *string `json:"built_at" db:"built_at"`
	Created_at           *string `json:"created_at" db:"created_at"`
	Discarded            *bool   `json:"discarded" db:"discarded" `
}

type Floor struct {
	Id          *int64  `json:"id" db:"id"`
	Floor       *int8   `json:"floor" db:"floor"`
	Name        *string `json:"name" db:"name"`
	Area        *string `json:"area" db:"area"`
	Created_at  *string `json:"created_at" db:"created_at"`
	Building_id *int64  `json:"building_id" db:"building_id"`
}

type Charge struct {
	Id               *int64   `json:"id" db:"id"`
	Year             *string  `json:"year" db:"year"`
	Month            *string  `json:"month" db:"month"`
	Date             *string  `json:"date" db:"date"`
	Electric_measure *float32 `json:"electric_measure" db:"electric_measure"`
	Water_measure    *float32 `json:"water_measure" db:"water_measure"`
	Created_at       *string  `json:"created_at" db:"created_at"`
	Floor_id         *int64   `json:"floor_id" db:"floor_id"`
}

func (b *BuildingStorage) GetBuilding(Id int64) (Building, error) {
	var building Building

	bQuery, err := GetPureFieldStr(Building{})
	if err != nil {
		return Building{}, err
	}

	query := fmt.Sprintf("SELECT %s FROM building WHERE id = ?", bQuery)
	err = b.Conn.Get(&building, query, Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return Building{}, nil
		}
		return Building{}, err
	}

	return building, nil
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

func (b *BuildingStorage) GetCharge(Id int64) (Charge, error) {
	var charge Charge

	cQuery, err := GetPureFieldStr(Charge{})
	if err != nil {
		return Charge{}, err
	}

	query := fmt.Sprintf("SELECT %s FROM floor WHERE id = ?", cQuery)
	err = b.Conn.Get(&charge, query, Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return Charge{}, nil
		}
		return Charge{}, err
	}

	return charge, nil
}

func (b *BuildingStorage) CreateNewBuilding(data Building) (int, error) {
	query, values, err := InsertQueryStr(data)
	if err != nil {
		return 0, err
	}
	fmt.Println(query, values)

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

func (b *BuildingStorage) CreateNewCharge(data Charge) (int, error) {
	query, values, err := InsertQueryStr(data)
	if err != nil {
		return 0, err
	}

	res, err := b.Conn.Exec(fmt.Sprintf(`INSERT INTO charge %s`, query), values...)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (b *BuildingStorage) UpdateBuilding(data Building) error {
	query, values, err := UpdateQueryStr("building", data, "id")
	if err != nil {
		return err
	}

	res, err := b.Conn.Exec(query, values...)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println(count)

	return nil
}

func (b *BuildingStorage) UpdateFloor(data Floor) error {
	query, values, err := UpdateQueryStr("floor", data, "id")
	if err != nil {
		return err
	}

	res, err := b.Conn.Exec(query, values...)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println(count)

	return nil
}

func (b *BuildingStorage) UpdateCharge(data Charge) error {
	query, values, err := UpdateQueryStr("charge", data, "id")
	if err != nil {
		return err
	}

	res, err := b.Conn.Exec(query, values...)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Println(count)

	return nil
}
