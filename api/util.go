package api

import (
	"github.com/jadidbourbaki/gofarm/geography"
)

// HaversinePoint takes the latitude, longitude of a given record and converts it into a geography.HaversinePoint
func (record FarmersMarketRecord) HaversinePoint() geography.HaversinePoint {
	return geography.HaversinePoint{Latitude: record.Location.Latitude, Longitude: record.Location.Longitude}
}
