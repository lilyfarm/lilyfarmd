package api

import (
	"github.com/jadidbourbaki/gofarm/geography"
)

// FarmersMarketApi is the interface for all the public functions our API
// will be providing
type FarmersMarketApi interface {
	// Refresh the underlying data being used by the API.
	Refresh() error
	// NearestN returns the nearest N farmer's markets to a given location.
	// If n is set to -1 it returns all the farmer's markets in ascending order of distance.
	NearestN(n int, location geography.HaversinePoint) ([]FarmersMarketRecord, error)

	// NearestNByZipCode returns the nearest N farmers' markets to a given zipcode
	// If n is set to -1 it returns all the farmer's markets in ascending order of distance.
	NearestNByZipCode(n int, zipCode string) ([]FarmersMarketRecord, error)
}
