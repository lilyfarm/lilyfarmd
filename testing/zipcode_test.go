package testing

import (
	"testing"

	"github.com/jadidbourbaki/gofarm/api"
	"github.com/stretchr/testify/require"
)

// Test ZipCode to HaversinePoint Lookup
func TestZipCodeToHaversinePoint(t *testing.T) {
	err := api.DefaultCredentials.LoadGeoNamesCredentials()
	require.Nil(t, err)

	randomZipCode := "77001"
	point, err := api.ZipCodeToHaversinePoint(randomZipCode)

	require.Nil(t, err)

	expectedLatitudeInt := 29
	expectedLongitudeInt := -95

	require.Equal(t, expectedLatitudeInt, int(point.Latitude))
	require.Equal(t, expectedLongitudeInt, int(point.Longitude))
}
