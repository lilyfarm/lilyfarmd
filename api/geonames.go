package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jadidbourbaki/gofarm/geography"
)

// GeoNamesResponse is the response received from geonames.org
type GeoNamesResponse struct {
	PostalCodes []GeoNamesRecord `json:"postalcodes"`
}

// GeoNamesRecord are individual records in the response received from geonames.org
type GeoNamesRecord struct {
	AdminCode2  string  `json:"adminCode2"`
	AdminCode1  string  `json:"adminCode1"`
	AdminName2  string  `json:"adminName2"`
	Longitude   float64 `json:"lng"`
	CountryCode string  `json:"countryCode"`
	PostalCode  string  `json:"postalcode"`
	AdminName1  string  `json:"adminName1"`
	PlaceName   string  `json:"placeName"`
	Latitude    float64 `json:"lat"`
}

// Uses the geonames.org API to convert a US Zip code into a Haversine Point
// Returns a default HaversinePoint as well as an error if it fails
func ZipCodeToHaversinePoint(zipcode string) (geography.HaversinePoint, error) {
	returnPoint := geography.HaversinePoint{}

	req, err := http.NewRequest(http.MethodGet, "http://api.geonames.org/postalCodeLookupJSON", nil)
	if err != nil {
		return returnPoint, fmt.Errorf("failed to create geonames request: %w", err)
	}

	query := req.URL.Query()

	query.Add("username", DefaultCredentials.geonames)
	query.Add("country", "US")
	query.Add("postalcode", zipcode)

	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return returnPoint, fmt.Errorf("request failed: %w", err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return returnPoint, fmt.Errorf("incorrect status code: %v", res.StatusCode)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return returnPoint, fmt.Errorf("failed to read body: %w", err)
	}

	var data GeoNamesResponse

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return returnPoint, fmt.Errorf("unmarshalling json: %w", err)
	}

	// Potentially give out an error for when there is
	// more than 1 record returned?
	if len(data.PostalCodes) == 0 {
		return returnPoint, fmt.Errorf("failed to find zipcode: %s", zipcode)
	}

	firstRecord := data.PostalCodes[0]

	returnPoint.Latitude = firstRecord.Latitude
	returnPoint.Longitude = firstRecord.Longitude

	return returnPoint, nil
}
