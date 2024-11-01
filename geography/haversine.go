package geography

import (
	"fmt"
	"math"
)

// HaversinePoint is a point in the HaversineMetricSpace
type HaversinePoint struct {
	Latitude  float64
	Longitude float64
}

// Dimensions implements the Dimensions function from the Point interface
func (h HaversinePoint) Dimensions() int {
	return 2
}

// Value implements the Value function from the Point interface
func (h HaversinePoint) Value(dimension int) (float64, error) {
	if dimension == 0 {
		return h.Latitude, nil
	}

	if dimension == 1 {
		return h.Longitude, nil
	}

	return 0, fmt.Errorf("invalid dimension: %v", dimension)
}

// HaversineMetricSpace uses Haversine distance
type HaversineMetricSpace struct {
}

var DefaultHaversineMetricSpace HaversineMetricSpace

// Haversin(θ) function. Taken from [1]
//
// [1]: https://gist.github.com/cdipaolo/d3f8db3848278b49db68
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// Taken from [1]:
// Distance function returns the distance (in meters) between two points of
//
//	a given longitude and latitude relatively accurately (using a spherical
//	approximation of the Earth) through the Haversin Distance Formula for
//	great arc distance on a sphere with accuracy for small distances
//
// point coordinates are supplied in degrees and converted into rad. in the func
//
// distance returned is METERS!!!!!!
// http://en.wikipedia.org/wiki/Haversine_formula
//
// [1]: https://gist.github.com/cdipaolo/d3f8db3848278b49db68
func haversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

func (hms HaversineMetricSpace) Distance(a Point, b Point) (float64, error) {
	if a.Dimensions() != 2 {
		return 0, fmt.Errorf("point a has incorrect dimensions")
	}

	if b.Dimensions() != 2 {
		return 0, fmt.Errorf("point b has incorrect dimensions")
	}

	la1, err := a.Value(0)
	if err != nil {
		return 0, err
	}

	lo1, err := a.Value(1)
	if err != nil {
		return 0, err
	}

	la2, err := b.Value(0)
	if err != nil {
		return 0, err
	}

	lo2, err := b.Value(1)
	if err != nil {
		return 0, err
	}

	return haversineDistance(la1, lo1, la2, lo2), nil
}

// Interface guard
var _ MetricSpace = (*HaversineMetricSpace)(nil)
