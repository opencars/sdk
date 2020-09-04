package toolkit

import (
	"time"
)

const (
	// APIKeyHeader ...
	APIKeyHeader = "X-Api-Key"
)

// Operation represents public registrations of transport.
type Operation struct {
	Person      string   `json:"person"`
	RegAddress  *string  `json:"reg_addr_koatuu"`
	RegCode     int16    `json:"registration_code"`
	Reg         string   `json:"registration"`
	Date        string   `json:"date"`
	DepCode     int32    `json:"dep_code"`
	Dep         string   `json:"dep"`
	Brand       string   `json:"brand"`
	Model       string   `json:"model"`
	Year        int16    `json:"year"`
	Color       string   `json:"color"`
	Kind        string   `json:"kind"`
	Body        string   `json:"body"`
	Purpose     string   `json:"purpose"`
	Fuel        *string  `json:"fuel"`
	Capacity    *int     `json:"capacity"`
	OwnWeight   *float64 `json:"own_weight"`
	TotalWeight *float64 `json:"total_weight"`
	Number      string   `json:"number"`
}

// Registration represents information from vehicle registration document.
type Registration struct {
	Brand        *string  `json:"brand"`
	Capacity     *int     `json:"capacity"`
	Color        string   `json:"color"`
	FirstRegDate *string  `json:"first_reg_date"`
	Date         *string  `json:"date"`
	Fuel         *string  `json:"fuel"`
	Kind         *string  `json:"kind"`
	Year         int      `json:"year"`
	Model        *string  `json:"model"`
	Code         string   `json:"code"`
	Number       string   `json:"number"`
	NumSeating   *int     `json:"num_seating"`
	NumStanding  *int     `json:"num_standing"`
	OwnWeight    *float64 `json:"own_weight"`
	RankCategory *string  `json:"rank_category"`
	TotalWeight  *float64 `json:"total_weight"`
	VIN          *string  `json:"vin"`
}

// WantedVehicle represents information about wanted vehicle.
type WantedVehicle struct {
	ID            string    `json:"id"`
	Brand         string    `json:"brand"`
	Maker         *string   `json:"maker,omitempty"`
	Model         *string   `json:"model,omitempty"`
	Color         *string   `json:"color,omitempty"`
	Number        *string   `json:"number,omitempty"`
	BodyNumber    *string   `json:"body_number,omitempty"`
	ChassisNumber *string   `json:"chassis_number,omitempty"`
	EngineNumber  *string   `json:"engine_number,omitempty"`
	OVD           string    `json:"ovd"`
	Kind          string    `json:"kind"`
	Status        string    `json:"status"`
	RevisionID    string    `json:"revision_id"`
	TheftDate     string    `json:"theft_date"`
	InsertDate    time.Time `json:"insert_date"`
}

// CoordinateALPR represents
type CoordinateALPR struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// ResultALPR represents result of recognition.
type ResultALPR struct {
	Coordinates []CoordinateALPR `json:"coordinates"`
	Plate       string           `json:"plate"`
}
