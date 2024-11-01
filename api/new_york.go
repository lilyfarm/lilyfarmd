package api

import (
	"cmp"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"

	"github.com/jadidbourbaki/gofarm/geography"
)

// NewYorkFarmersMarketAPIEndpoint is the public endpoint for the New York State Farmer's Market API
const NewYorkFarmersMarketAPIEndpoint = "https://data.ny.gov/resource/farmersmarkets.json"

type NewYorkFarmersMarketLink struct {
	Url string `json:"url"`
}

// NewYorkFarmersMarketRecord is the individual record for each Farmer's Market in New York State.
// Take a look at NewYorkFarmersMarketDataset below to see how this struct was constructed.
// The things that were missing in the Data Layouts file by the New York Dept of Agriculture and
// Markets [1], we were able to find in documentation by Socrata [2]
// Note on Operation Months Code:
// We extracted the meanings of these values from [2]
// SPRING (P) = operating at some point in April or May;
// SUMMER (M) = operating at some point in June, July, August, September, October, or November;
// EXTENDED SEASON (X) = operating at some point in December;
// WINTER (W) = operating at some point in January, February, March;
// YEAR-ROUND (YR) = continually operating at the same location.
// Categories are NOT mutually exclusive.
//
// [1]: https://data.ny.gov/api/assets/68393954-B01F-422B-9600-71854BD5118B?download=true
// [2]: https://dev.socrata.com/foundry/data.ny.gov/xjya-f8ng
type NewYorkFarmersMarketRecord struct {
	County                        string                   `json:"county"`                // County the Farmers’ Market is located in
	MarketName                    string                   `json:"market_name"`           // Name of the Farmers’ Market
	Location                      string                   `json:"market_location"`       // General place where the market is set up
	AddressLine1                  string                   `json:"address_line_1"`        // 1st address line of the physical address for this location
	City                          string                   `json:"city"`                  // City name of the physical location
	StateCode                     string                   `json:"state"`                 // 2 character state code for the physical address for this location
	Zip                           string                   `json:"zip"`                   // Zip code for the physical address for this location
	Contact                       string                   `json:"contact"`               // Contact person for the Market
	Telephone                     string                   `json:"phone"`                 // Phone number of the market contact
	MarketLink                    NewYorkFarmersMarketLink `json:"market_link"`           // Website or other link to this market’s online presence
	OperationHours                string                   `json:"operation_hours"`       // Day and Hours the market is open
	OperationSeason               string                   `json:"operation_season"`      // Month and day when the market opens and closes for the year
	OperationMonthsCode           string                   `json:"operation_months_code"` // See "Note on Operation Months Code" above
	FarmersMarketNutritionProgram string                   `json:"fmnp"`                  // Y indicates that this market is part of the Farmers Market Nutrition Program.
	Latitude                      json.Number              `json:"latitude"`              // Latitude associated with this market location
	Longitude                     json.Number              `json:"longitude"`             // Longitude associated with this market location
	SnapStatus                    string                   `json:"snap_status"`           // Y indicates the market accepts SNAP; www.snaptomarket.com

	// We keep these versions for more convenient distance computation.
	// They are parsed from the json.Number values above.
	latitudeFloat  float64
	longitudeFloat float64
}

// FarmersMarketRecord converts a NewYorkFarmersMarketRecord to a FarmersMarketRecord
func (record NewYorkFarmersMarketRecord) FarmersMarketRecord() FarmersMarketRecord {
	return FarmersMarketRecord{
		Name:        record.MarketName,
		Description: "",
		Address: FarmersMarketAddress{
			Street:  record.AddressLine1,
			City:    record.City,
			State:   record.StateCode,
			ZipCode: record.Zip,
		},
		Location: geography.HaversinePoint{
			Latitude:  record.latitudeFloat,
			Longitude: record.longitudeFloat,
		},
		Website:                       record.MarketLink.Url,
		OperationHours:                record.OperationHours,
		OperationSeason:               record.OperationSeason,
		OperationMonthsCode:           record.OperationMonthsCode,
		FarmersMarketNutritionProgram: record.FarmersMarketNutritionProgram == "Y",
		SnapStatus:                    record.SnapStatus == "Y",

		// We do not have a distance value so we will use -1 as the default
		Distance: -1,
	}
}

// ParseNewYorkFarmersMarketDataset takes in raw bytes representing JSON content and
// returns a structured NewYorkFarmersMarketDataset object, or an error if unmarshalling
// failed.
func ParseNewYorkFarmersMarketDataset(dataset []byte) ([]NewYorkFarmersMarketRecord, error) {
	parsedDataset := []NewYorkFarmersMarketRecord{}

	err := json.Unmarshal(dataset, &parsedDataset)
	if err != nil {
		return parsedDataset, fmt.Errorf("could not parse new york farmer's market dataset: %w", err)
	}

	for idx := range parsedDataset {
		record := &parsedDataset[idx]

		latitudeFloat, err := record.Latitude.Float64()
		if err != nil {
			return parsedDataset, fmt.Errorf("could not parse latitude float: %w", err)
		}

		record.latitudeFloat = latitudeFloat

		longitudeFloat, err := record.Longitude.Float64()
		if err != nil {
			return parsedDataset, fmt.Errorf("could not parse longitude float: %w", err)
		}

		record.longitudeFloat = longitudeFloat
	}

	return parsedDataset, nil
}

// FetchNewYorkStateData returns all the Farmers' Market Data from the New York State Farmers'
// Market API endpoint.
func FetchNewYorkStateData() ([]byte, error) {
	failureErrorFormat := "could not fetch new york data: %w"

	req, err := http.NewRequest("GET", NewYorkFarmersMarketAPIEndpoint, nil)
	if err != nil {
		return nil, fmt.Errorf(failureErrorFormat, err)
	}

	// We expect a json file
	req.Header.Set("Accept", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf(failureErrorFormat, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf(failureErrorFormat, err)
	}

	return body, nil
}

// NewYorkFarmersMarketApi is the implementation of the FarmersMarketApi that uses the New York API endpoint
type NewYorkFarmersMarketApi struct {
	metricSpace geography.MetricSpace
	dataset     []FarmersMarketRecord
	// lazy load the API
	loadedApi bool
}

func (fmApi *NewYorkFarmersMarketApi) Refresh() error {
	body, err := FetchNewYorkStateData()
	if err != nil {
		return fmt.Errorf("could not fetch data: %w", err)
	}

	records, err := ParseNewYorkFarmersMarketDataset(body)
	if err != nil {
		return fmt.Errorf("could not parse data: %w", err)
	}

	// If the dataset is already populated, empty it
	fmApi.dataset = []FarmersMarketRecord{}

	for _, record := range records {
		fmApi.dataset = append(fmApi.dataset, record.FarmersMarketRecord())
	}

	return nil

}

// NewNewYorkMarketApi returns a pointer to a freshly constructed New York Farmers Market Api
func NewNewYorkMarketApi() (*NewYorkFarmersMarketApi, error) {
	fmApi := &NewYorkFarmersMarketApi{}

	fmApi.metricSpace = geography.DefaultHaversineMetricSpace

	if err := DefaultCredentials.LoadGeoNamesCredentials(); err != nil {
		return nil, err
	}

	return fmApi, nil
}

func (fmApi *NewYorkFarmersMarketApi) NearestN(n int, location geography.HaversinePoint) ([]FarmersMarketRecord, error) {
	if !fmApi.loadedApi {
		if err := fmApi.Refresh(); err != nil {
			return nil, err
		}

		fmApi.loadedApi = true
	}

	sortedDataset, err := fmApi.sortDatasetForLocation(location)
	if err != nil {
		return nil, fmt.Errorf("nearest n: %w", err)
	}

	var records []FarmersMarketRecord

	if n == -1 {
		n = len(sortedDataset)
	}

	// The only negative value we accept is -1, which is handled above
	if n < 0 {
		return records, fmt.Errorf("invalid value for n")
	}

	minLen := min(len(sortedDataset), n)

	for i := 0; i < minLen; i++ {
		records = append(records, sortedDataset[i])
	}

	if n > minLen {
		return records, fmt.Errorf("did not find all n records")
	}

	return records, nil
}

func (fmApi *NewYorkFarmersMarketApi) NearestNByZipCode(n int, zipcode string) ([]FarmersMarketRecord, error) {
	location, err := ZipCodeToHaversinePoint(zipcode)
	if err != nil {
		return nil, fmt.Errorf("zip code lookup failed: %w", err)
	}

	return fmApi.NearestN(n, location)
}

// sortDatasetForLocation sorts all farmer's markets in ascending order based on their Haversine distance from
// a given location
func (fmApi *NewYorkFarmersMarketApi) sortDatasetForLocation(location geography.Point) ([]FarmersMarketRecord, error) {
	dataset := fmApi.dataset

	// This is important, otherwise the deepcopy below will not work
	records := make([]FarmersMarketRecord, len(dataset))

	for idx, record := range dataset {
		distance, err := fmApi.metricSpace.Distance(record.HaversinePoint(), location)
		if err != nil {
			return nil, fmt.Errorf("could not compute distance: %w", err)
		}

		dataset[idx].Distance = distance
	}

	distanceCmp := func(a, b FarmersMarketRecord) int {
		return cmp.Compare(a.Distance, b.Distance)
	}

	// We need a deep copy here
	copy(records, dataset)

	slices.SortFunc(records, distanceCmp)

	return records, nil
}

// Verify that NewYorkFarmersMarketApi implements FarmersMarketApi
var _ FarmersMarketApi = (*NewYorkFarmersMarketApi)(nil)
