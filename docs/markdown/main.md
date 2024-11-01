The Lily Farm API currently has two API calls. No API Key or credentials are required. However, please keep request traffic manageable, our infrastructure is very modest currently.

GNU GPL Licensed source code is available [here](https://github.com/jadidbourbaki/lilyfarm-public).

#### Data Sources

The Lily Farm API supports multiple data sources. There are two data sources currently supported:

- `usda` (data fetched from the United States Department of Agriculture)
- `newyork` (data fetched from the New York State government)

#### API Reference

##### HTTP GET nearestNJson

Get the nearest `N` Farmers' Markets from a given `latitude` and `longitude`. Only accepts 
`GET` requests. 

###### Example:

The following returns the nearest 10 Farmers' Markets from Location `(LAT, LON)`.

```
lilyfarm.org/nearestNJson?n=10&latitude=LAT&longitude=LON&datasource=usda
```


##### HTTP GET nearestNJsonByZipCode

Get the nearest `N` Farmers' Markets from a given `zipCode`. Only accepts `GET` requests.

###### Example:

The following returns the nearest 10 Farmers' Markets from Zip code `ZIP`

```
lilyfarm.org/nearestNJsonByZipCode?n=10&zipCode=ZIP&datasource=usda
```

##### Serialization

For both the above API calls, Lily Farm returns an array of JSON records of type FarmersMarketRecord.

```go
type FarmersMarketRecord struct {
	Name                          string                   `json:"name"`
	Description                   string                   `json:"description,omitempty"`
	Address                       FarmersMarketAddress     `json:"address"`
	Distance                      float64                  `json:"distance"`
	Website                       string                   `json:"website,omitempty"`
	Location                      geography.HaversinePoint `json:"location"`
	OperationHours                string                   `json:"operation_hours,omitempty"`       // Day and Hours the market is open
	OperationSeason               string                   `json:"operation_season,omitempty"`      // Month and day when the market opens and closes for the year
	OperationMonthsCode           string                   `json:"operation_months_code,omitempty"` 
	FarmersMarketNutritionProgram bool                     `json:"fmnp,omitempty"`                  // true indicates that this market is part of the Farmers Market Nutrition Program.
	SnapStatus                    bool                     `json:"snap_status,omitempty"`           // true indicates the market accepts SNAP; www.snaptomarket.com
}
```

More features to come in the future! 🙂

Please see our [about](/about) page for information on contacting us for developer API related questions, bug reports, and feature requests.