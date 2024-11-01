# api

This is where the logic for the API exists. It needs the following environment variables:

1. `GOFARM_GEONAMES_CREDENTIALS` --- the username for geonames.org web services
2. `GOFARM_USDA_CREDENTIALS` --- the api key for the USDA Farmers' Market API

See `credentials.go` for more details on how these are loaded.

Only the first one is needed for the New York State specific API. (`new_york.go`)

Both are needed for the United States wide API. (`usda.go`)