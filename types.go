package toolkit

import (
	"time"
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
	DFirstReg    *string  `json:"d_first_reg"`
	DReg         *string  `json:"d_reg"`
	Fuel         *string  `json:"fuel"`
	Kind         *string  `json:"kind"`
	MakeYear     int      `json:"make_year"`
	Model        *string  `json:"model"`
	NDoc         string   `json:"n_doc"`
	SDoc         string   `json:"s_doc"`
	NRegNew      string   `json:"n_reg_new"`
	NSeating     *int     `json:"n_seating"`
	NStanding    *int     `json:"n_standing"`
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
