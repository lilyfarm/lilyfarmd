package geography

// MeterstoMiles converts length in meters to length in miles
func MetersToMiles(meters float64) float64 {
	// Conversion factor: 1 meter is approximately 0.000621371 miles
	miles := meters * 0.000621371
	return miles
}

// MetersToKilometersInteger converts meters to kilometers and then casts
// it to an int64.
func MetersToKilometersInteger(meters float64) int64 {
	return int64(meters / 1000.0)
}
