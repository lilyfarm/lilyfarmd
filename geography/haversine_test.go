package geography

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// These tests were adapted from [1]
// [1]: https://github.com/jftuga/geodist/

var newYork = HaversinePoint{40.7128, 74.0060}
var sanDiego = HaversinePoint{32.7157, 117.1611}
var elPaso = HaversinePoint{31.7619, 106.4850}

func TestHaversineDistanceNewYorkSanDiego(t *testing.T) {
	expectedDistance := int64(3911)

	distance, err := DefaultHaversineMetricSpace.Distance(newYork, sanDiego)
	assert.Nil(t, err)
	distanceInKm := MetersToKilometersInteger(distance)

	assert.Equal(t, distanceInKm, expectedDistance)

	// Should probably check for commutativity in at least one test
	reverseDistance, err := DefaultHaversineMetricSpace.Distance(sanDiego, newYork)
	assert.Nil(t, err)
	reverseDistanceInKm := MetersToKilometersInteger(reverseDistance)
	assert.Equal(t, reverseDistanceInKm, expectedDistance)
}

func TestHaversineDistanceSanDiegoElPaso(t *testing.T) {
	expectedDistance := int64(1010)

	distance, err := DefaultHaversineMetricSpace.Distance(sanDiego, elPaso)
	assert.Nil(t, err)
	distanceInKm := MetersToKilometersInteger(distance)

	assert.Equal(t, distanceInKm, expectedDistance)
}
