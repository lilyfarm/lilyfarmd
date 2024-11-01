package testing

import (
	"encoding/json"
	"testing"

	"github.com/jadidbourbaki/gofarm/api"
	"github.com/jadidbourbaki/gofarm/geography"
	"github.com/stretchr/testify/require"
)

var houstonLocation = geography.HaversinePoint{Latitude: 29.784203, Longitude: -95.553394}

// TestFetchingUSDAData checks if the data is still accessible with out API Key and
// whether the USDA data is still up.
func TestFetchingUSDAData(t *testing.T) {
	err := api.DefaultCredentials.LoadUSDACredentials()
	require.Nil(t, err)

	// Just a random location to test
	_, err = api.FetchUSDADataByLocation(houstonLocation)
	require.Nil(t, err)
}

// TestMarshallignUSDAData checks if the data is being delivered in *a* correct JSON format
func TestMarshallingUSDAData(t *testing.T) {
	err := api.DefaultCredentials.LoadUSDACredentials()
	require.Nil(t, err)

	data, err := api.FetchUSDADataByLocation(houstonLocation)
	require.Nil(t, err)

	var records map[string]interface{}
	err = json.Unmarshal(data, &records)

	require.Nil(t, err)

	_, err = json.Marshal(records)
	require.Nil(t, err)
}
