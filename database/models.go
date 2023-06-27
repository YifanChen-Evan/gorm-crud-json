package database

// define data model based on JSON file
type Employee struct {
	// fields with an initial capital letter can be accessed in other packages
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
	Email       string `json:"email"`
	Department  string `json:"department"`
	Reason      string `json:"reason"`
	StartDate   string `json:"startDate"`
	DaysOff     int    `json:"daysOff"`
}
