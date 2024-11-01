package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseNewYorkFarmersMarketDatasetEmptyDataset(t *testing.T) {
	// An empty dataset should not be an error.
	emptyDatasetRaw := []byte("[]")
	emptyDataset, err := ParseNewYorkFarmersMarketDataset(emptyDatasetRaw)

	require.Nil(t, err)
	require.Equal(t, len(emptyDataset), 0)
}

func TestParseNewYorkFarmersMarketDatasetOneRecordDataset(t *testing.T) {
	// A one record dataset
	oneRecordDatasetRaw := []byte(`[
		{
			"county": "Albany",
			"market_name": "Albany County Farmers' Market",
			"market_location": "51 S. Pearl St (in front of  MVP Arena)",
			"address_line_1": "51 S. Pearl St",
			"city": "Albany",
			"state": "NY",
			"zip": "12207",
			"contact": "Jevan Dollard",
			"phone": "5184652143",
			"market_link": {
			"url": "https://www.downtownalbany.org/albany-county-farmers-market"
			},
			"operation_hours": "Sun 10am-2pm",
			"operation_season": "July 14-September 29",
			"operation_months_code": "M",
			"fmnp": "N",
			"snap_status": "Y",
			"latitude": "42.64819",
			"longitude": "-73.75383"
		}]`)

	oneRecordDatasetExpectedRecord := NewYorkFarmersMarketRecord{
		County:                        "Albany",
		MarketName:                    "Albany County Farmers' Market",
		Location:                      "51 S. Pearl St (in front of  MVP Arena)",
		AddressLine1:                  "51 S. Pearl St",
		City:                          "Albany",
		StateCode:                     "NY",
		Zip:                           "12207",
		Contact:                       "Jevan Dollard",
		Telephone:                     "5184652143",
		MarketLink:                    NewYorkFarmersMarketLink{Url: "https://www.downtownalbany.org/albany-county-farmers-market"},
		OperationHours:                "Sun 10am-2pm",
		OperationSeason:               "July 14-September 29",
		OperationMonthsCode:           "M",
		FarmersMarketNutritionProgram: "N",
		SnapStatus:                    "Y",
		Latitude:                      "42.64819",
		Longitude:                     "-73.75383",
		latitudeFloat:                 42.64819,
		longitudeFloat:                -73.75383,
	}

	oneRecordDataset, err := ParseNewYorkFarmersMarketDataset(oneRecordDatasetRaw)
	require.Nil(t, err)
	require.Equal(t, len(oneRecordDataset), 1)
	require.Equal(t, oneRecordDataset[0], oneRecordDatasetExpectedRecord)
}

func TestParseNewYorkFarmersMarketDatasetMultiRecordDataset(t *testing.T) {
	// A multi record dataset
	multiRecordDatasetRaw := []byte(`[
		{
			"county": "Albany",
			"market_name": "Albany County Farmers' Market",
			"market_location": "51 S. Pearl St (in front of  MVP Arena)",
			"address_line_1": "51 S. Pearl St",
			"city": "Albany",
			"state": "NY",
			"zip": "12207",
			"contact": "Jevan Dollard",
			"phone": "5184652143",
			"market_link": {
			"url": "https://www.downtownalbany.org/albany-county-farmers-market"
			},
			"operation_hours": "Sun 10am-2pm",
			"operation_season": "July 14-September 29",
			"operation_months_code": "M",
			"fmnp": "N",
			"snap_status": "Y",
			"latitude": "42.64819",
			"longitude": "-73.75383"
		},
		 {
		"county": "New York",
		"market_name": "82nd Street Greenmarket",
		"market_location": "82nd St btwn 1st & York Aves",
		"address_line_1": "408 East 82nd Street",
		"city": "New York",
		"state": "NY",
		"zip": "10128",
		"contact": "Tutu Badaru",
		"phone": "2127887900",
		"market_link": {
		"url": "http://www.grownyc.org/greenmarket/manhattan/82nd-street"
		},
		"operation_hours": "Sat 9am-2:30pm",
		"operation_season": "Year-round",
		"operation_months_code": "YR",
		"fmnp": "Y",
		"snap_status": "Y",
		"latitude": "40.77394",
		"longitude": "-73.9506"
		}]`)

	multiRecordDataset, err := ParseNewYorkFarmersMarketDataset(multiRecordDatasetRaw)
	require.Nil(t, err)
	require.Equal(t, len(multiRecordDataset), 2)
}
