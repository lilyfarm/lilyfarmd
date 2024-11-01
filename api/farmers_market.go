package api

import (
	"github.com/jadidbourbaki/gofarm/geography"
)

// FarmersMarketAddress is the address of the Farmers' Market
type FarmersMarketAddress struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

// FarmersMarketRecord is the default structure for storing information about Farmers'
// Market entries from various APIs. this should contain all the relevant information
// that is common amongst most Farmers' Market datasets. Aside from the Name,
// the Address, the Distance, and the Location, everything else is optional.
type FarmersMarketRecord struct {
	Name                          string                   `json:"name"`
	Description                   string                   `json:"description,omitempty"`
	Address                       FarmersMarketAddress     `json:"address"`
	Distance                      float64                  `json:"distance"`
	Website                       string                   `json:"website,omitempty"`
	Location                      geography.HaversinePoint `json:"location"`
	OperationHours                string                   `json:"operation_hours,omitempty"`       // Day and Hours the market is open
	OperationSeason               string                   `json:"operation_season,omitempty"`      // Month and day when the market opens and closes for the year
	OperationMonthsCode           string                   `json:"operation_months_code,omitempty"` // See "Note on Operation Months Code" in NewYorkFarmersMarketRecord
	FarmersMarketNutritionProgram bool                     `json:"fmnp,omitempty"`                  // true indicates that this market is part of the Farmers Market Nutrition Program.
	SnapStatus                    bool                     `json:"snap_status,omitempty"`           // true indicates the market accepts SNAP; www.snaptomarket.com
}
