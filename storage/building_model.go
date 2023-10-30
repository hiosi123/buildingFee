package storage

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
	MeasureType *int8   `json:"measureType" db:"measureType"` //1: 기본, 2: 25x
	Created_at  *string `json:"created_at" db:"created_at"`
	Building_id *int64  `json:"building_id" db:"building_id"`
}

type Charge struct {
	Id                  *int64   `json:"id" db:"id"`
	Year                *string  `json:"year" db:"year"`
	Month               *string  `json:"month" db:"month"`
	Date                *string  `json:"date" db:"date"`
	Measure_number      *int8    `json:"measure_number" db:"measure_number"`
	Electric_measure    *float32 `json:"electric_measure" db:"electric_measure"`
	Electric_difference *float32 `json:"electric_difference" db:"electric_difference"`
	Water_measure       *float32 `json:"water_measure" db:"water_measure"`
	Water_difference    *float32 `json:"water_difference" db:"water_difference"`
	Created_at          *string  `json:"created_at" db:"created_at"`
	Floor_id            *int64   `json:"floor_id" db:"floor_id"`
}
