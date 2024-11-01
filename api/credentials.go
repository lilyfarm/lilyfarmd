package api

import (
	"fmt"
	"os"
)

// usdaCredentialsEnvironmentVariable is the variable used to get the API Key
// for the USDA Farmers' Market API. This should be set by the environment.
const usdaCredentialsEnvironmentVariable = "LILYFARM_USDA_CREDENTIALS"

// geonamesCredentialsEnvironmentVariable is the variable used to get the
// username for the geonames.org API. This should be set by the environment.
const geonamesCredentialsEnvironmentVariable = "LILYFARM_GEONAMES_CREDENTIALS"

// Credentials stores the credentials for all APIs used by this package
type Credentials struct {
	// usda stores credentials for the USDA Farmers' Market API
	usda string

	// geonames stores credentials for the geonames.org API we use
	// for zipcode to longitude, latitude conversion
	geonames string
}

// loadUSDACredentials populates the usda field in the credentials struct,
// or returns an error if usdaCredentialsEnvironmentVariable was not set.
func (c *Credentials) LoadUSDACredentials() error {
	c.usda = os.Getenv(usdaCredentialsEnvironmentVariable)
	if c.usda == "" {
		return fmt.Errorf("usda credentials not found in environment variable: %s", usdaCredentialsEnvironmentVariable)
	}

	return nil
}

// loadGeoNamesCredentials populates the geonames field in the credentials struct,
// or returns an error if the geoNamesCredentialsEnvironmentVariable was not set.
func (c *Credentials) LoadGeoNamesCredentials() error {
	c.geonames = os.Getenv(geonamesCredentialsEnvironmentVariable)
	if c.geonames == "" {
		return fmt.Errorf("geonames credentials not found in environment variable: %s", geonamesCredentialsEnvironmentVariable)
	}

	return nil
}

var DefaultCredentials Credentials
