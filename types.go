package sdk

// Operation represents public registrations of transport.
type Operation struct {
	Person      string `json:"person"`
	Address     string `json:"address"`
	Code        int    `json:"operation"`
	Description string `json:"description"`
	Date        string `json:"date"`
	OfficeID    int    `json:"office_id"`
	OfficeName  string `json:"office_name"`
	Brand       string `json:"brand"`
	Model       string `json:"model"`
	Year        int    `json:"year"`
	Color       string `json:"color"`
	Kind        string `json:"kind"`
	Body        string `json:"body"`
	Purpose     string `json:"purpose"`
	Fuel        string `json:"fuel"`
	Capacity    int    `json:"capacity"`
	Weight      int    `json:"weight"`
	Number      string `json:"number"`
}

// Registration represents information from vehicle registration document.
type Registration struct {
	Brand       string `json:"brand"`
	Capacity    int    `json:"capacity"`
	Color       string `json:"color"`
	FirstReg    string `json:"first_reg"`
	Date        string `json:"date"`
	Fuel        string `json:"fuel"`
	Kind        string `json:"kind"`
	Body        string `json:"body"`
	Year        int    `json:"year"`
	Model       string `json:"model"`
	Code        string `json:"code"`
	Number      string `json:"number"`
	TotalWeight int    `json:"total_weight`
	OwnWeight   int    `json:"own_weight"`
	Category    string `json:"category"`
	VIN         string `json:"vin"`
}
