package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/jadidbourbaki/gofarm/geography"
)

// webBrowserUserAgent is used to specify the User-Agent for a well-known web browser to be used to get
// access to the Farmers' Market API. This is because it seems like the USDA Data Sharing API
// returns a 403 Access Forbidden when we do not specify the User-Agent for a popular web browser
// even if we have a valid API Key.
const webBrowserUserAgent = "Mozilla/5.0 (X11; Linux x86_64; rv:124.0) Gecko/20100101 Firefox/124.0"

// usdaDefaultMiles is the default range to return Farmers' Markets in for the USDA API
const usdaDefaultMiles = 100

// FetchUSDADataByLocation fetches the data from the USDA API endpoint
// location is the location to center the dataset on
// The radius it uses is api.usdaDefaultMiles
func FetchUSDADataByLocation(location geography.HaversinePoint) ([]byte, error) {
	return FetchUSDADataByLocationAndRadius(location, usdaDefaultMiles)
}

// FetchUSDADataByLocationAndRadius fetches the data from the USDA API endpoint
// location is the location to center the dataset on
// radius is the radius to fetch the data within, in miles
// From experimentation, it seems like the results return in ascending
// order of distance from the provided location.
func FetchUSDADataByLocationAndRadius(location geography.HaversinePoint, radius int) ([]byte, error) {
	const url = "https://www.usdalocalfoodportal.com/api/farmersmarket"

	logPrefix := "Fetching USDA Data"

	apiKey := DefaultCredentials.usda

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logPrefix, err)
	}

	query := req.URL.Query()

	query.Add("apikey", apiKey)
	query.Add("x", fmt.Sprint(location.Longitude))
	query.Add("y", fmt.Sprint(location.Latitude))
	query.Add("radius", fmt.Sprint(radius))

	req.URL.RawQuery = query.Encode()

	// The USDA Farmer's Market API seems to return an http.StatusForbidden unless
	// a User-Agent is specified.
	req.Header.Set("User-Agent", webBrowserUserAgent)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logPrefix, err)
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: incorrect status code %v", logPrefix, res.StatusCode)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", logPrefix, err)
	}

	return bodyBytes, nil
}

// USDARecord is a structure to unmarshal the json received from the USDA API
type USDARecord struct {
	BriefDesc       string `json:"brief_desc"`
	ContactEmail    string `json:"contact_email"`
	ContactName     string `json:"contact_name"`
	ContactPhone    string `json:"contact_phone"`
	DirectoryName   string `json:"directory_name"`
	DirectoryType   string `json:"directory_type"`
	Distance        string `json:"distance"`
	ListingDesc     string `json:"listing_desc"`
	ListingID       string `json:"listing_id"`
	ListingImage    string `json:"listing_image"`
	ListingName     string `json:"listing_name"`
	LocationAddress string `json:"location_address"`
	LocationCity    string `json:"location_city"`
	LocationState   string `json:"location_state"`
	LocationStreet  string `json:"location_street"`
	LocationX       string `json:"location_x"`
	LocationY       string `json:"location_y"`
	LocationZipcode string `json:"location_zipcode"`
	MediaBlog       string `json:"media_blog"`
	MediaFacebook   string `json:"media_facebook"`
	MediaInstagram  string `json:"media_instagram"`
	MediaPinterest  string `json:"media_pinterest"`
	MediaTwitter    string `json:"media_twitter"`
	MediaWebsite    string `json:"media_website"`
	MediaYoutube    string `json:"media_youtube"`
	MyDesc          string `json:"mydesc"`
	Term            string `json:"term"`
	UpdateTime      string `json:"updatetime"`
}

func (record USDARecord) FarmersMarketRecord() (FarmersMarketRecord, error) {
	emptyRecord := FarmersMarketRecord{}

	distance, err := strconv.ParseFloat(record.Distance, 64)
	if err != nil {
		return emptyRecord, fmt.Errorf("distance malformatted: %w", err)
	}

	latitude, err := strconv.ParseFloat(record.LocationY, 64)
	if err != nil {
		return emptyRecord, fmt.Errorf("latitude malformatted: %w", err)
	}

	longitude, err := strconv.ParseFloat(record.LocationX, 64)
	if err != nil {
		return emptyRecord, fmt.Errorf("longitude malformatted: %w", err)
	}

	return FarmersMarketRecord{
		Name:        record.ListingName,
		Description: record.BriefDesc,
		Address: FarmersMarketAddress{
			Street:  record.LocationStreet,
			City:    record.LocationCity,
			State:   record.LocationState,
			ZipCode: record.LocationZipcode,
		},
		Distance: distance,
		Location: geography.HaversinePoint{
			Latitude:  latitude,
			Longitude: longitude,
		},
		Website: record.MediaWebsite,
	}, nil
}

// USDADataset is a structure to unamrshal the json received from the USDA API
type USDADataset struct {
	Data []USDARecord `json:"data"`
}

// ParseUSDADataset parses data from the USDA API Endpoint into a set of
// Farmers' Market Records
func ParseUSDADataset(dataset []byte) ([]USDARecord, error) {
	logPrefix := "could not parse usda data"
	usdaDataset := USDADataset{}

	err := json.Unmarshal(dataset, &usdaDataset)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", logPrefix, err)
	}

	return usdaDataset.Data, nil
}

// USDAFarmersMarketApi is the implementation of the FarmersMarketApi that uses the USDA API endpoint
type USDAFarmersMarketApi struct {
}

func NewUSDAFarmersMarketApi() (*USDAFarmersMarketApi, error) {
	if err := DefaultCredentials.LoadGeoNamesCredentials(); err != nil {
		return nil, err
	}

	if err := DefaultCredentials.LoadUSDACredentials(); err != nil {
		return nil, err
	}

	return &USDAFarmersMarketApi{}, nil
}

// Refresh does not really do anything for the USDA API as
// all our data is collected live.
func (api *USDAFarmersMarketApi) Refresh() error {
	return nil
}

func (api *USDAFarmersMarketApi) NearestN(n int, location geography.HaversinePoint) ([]FarmersMarketRecord, error) {
	dataset, err := FetchUSDADataByLocation(location)
	if err != nil {
		return nil, fmt.Errorf("fetching usda data: %w", err)
	}

	parsedDataset, err := ParseUSDADataset(dataset)
	if err != nil {
		return nil, fmt.Errorf("parsing usda data: %w", err)
	}

	generalizedDataset := []FarmersMarketRecord{}

	for _, record := range parsedDataset {
		generalizedRecord, err := record.FarmersMarketRecord()

		if err != nil {
			return nil, fmt.Errorf("converting to farmers market record: %w", err)
		}

		generalizedDataset = append(generalizedDataset, generalizedRecord)
	}

	var records []FarmersMarketRecord

	if n == -1 {
		n = len(generalizedDataset)
	}

	// The only negative value we accept is -1, which is handled above
	if n < 0 {
		return records, fmt.Errorf("invalid value for n")
	}

	minLen := min(len(generalizedDataset), n)

	for i := 0; i < minLen; i++ {
		records = append(records, generalizedDataset[i])
	}

	if n > minLen {
		return records, fmt.Errorf("did not find all n records")
	}

	return records, nil
}

func (api *USDAFarmersMarketApi) NearestNByZipCode(n int, zipcode string) ([]FarmersMarketRecord, error) {
	location, err := ZipCodeToHaversinePoint(zipcode)
	if err != nil {
		return nil, fmt.Errorf("zip code lookup failed: %w", err)
	}

	return api.NearestN(n, location)
}

// Verify that USDAFarmersMarketApi implements FarmersMarketApi
var _ FarmersMarketApi = (*USDAFarmersMarketApi)(nil)
