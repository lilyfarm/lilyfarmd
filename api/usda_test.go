package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseUSDADatasetEmptyDataset(t *testing.T) {
	// An empty dataset should not be an error.
	emptyDatasetRaw := []byte("{\"data\": []}")
	emptyDataset, err := ParseUSDADataset(emptyDatasetRaw)

	require.Nil(t, err)
	require.Equal(t, len(emptyDataset), 0)
}

func TestParseUSDADatasetOneRecordDataset(t *testing.T) {
	// A one record dataset
	oneRecordDatasetRaw := []byte(`
	{
	"data": [
		{
			"brief_desc": "",
			"contact_email": null,
			"contact_name": null,
			"contact_phone": null,
			"directory_name": "farmers market",
			"directory_type": "farmersmarket",
			"distance": "0.2492313843284663",
			"listing_desc": null,
			"listing_id": "312052",
			"listing_image": "default-farmersmarket-4-3.jpg",
			"listing_name": "Potsdam Farmers Market",
			"location_address": "3 Riverside Dr, New York, New York 10023",
			"location_city": "New York",
			"location_state": "New York",
			"location_street": "3 Riverside Dr",
			"location_x": "-73.98557883420489",
			"location_y": "40.7805342225485",
			"location_zipcode": "10023",
			"media_blog": null,
			"media_facebook": null,
			"media_instagram": null,
			"media_pinterest": null,
			"media_twitter": null,
			"media_website": null,
			"media_youtube": null,
			"mydesc": "",
			"term": "",
			"updatetime": "Mar 27th, 2023"
		}
		]
	}`)

	expectedRecord := USDARecord{
		BriefDesc:       "",
		ContactEmail:    "",
		ContactName:     "",
		ContactPhone:    "",
		DirectoryName:   "farmers market",
		DirectoryType:   "farmersmarket",
		Distance:        "0.2492313843284663",
		ListingDesc:     "",
		ListingID:       "312052",
		ListingImage:    "default-farmersmarket-4-3.jpg",
		ListingName:     "Potsdam Farmers Market",
		LocationAddress: "3 Riverside Dr, New York, New York 10023",
		LocationCity:    "New York",
		LocationState:   "New York",
		LocationStreet:  "3 Riverside Dr",
		LocationX:       "-73.98557883420489",
		LocationY:       "40.7805342225485",
		LocationZipcode: "10023",
		MediaBlog:       "",
		MediaFacebook:   "",
		MediaInstagram:  "",
		MediaPinterest:  "",
		MediaTwitter:    "",
		MediaWebsite:    "",
		MediaYoutube:    "",
		MyDesc:          "",
		Term:            "",
		UpdateTime:      "Mar 27th, 2023",
	}

	oneRecordDataset, err := ParseUSDADataset(oneRecordDatasetRaw)
	require.Nil(t, err)
	require.Equal(t, len(oneRecordDataset), 1)
	require.Equal(t, oneRecordDataset[0], expectedRecord)
}

func TestParseUSDADatasetMultiRecordDataset(t *testing.T) {
	// A multi record dataset
	multiRecordDatasetRaw := []byte(`
	{
	"data": [
		{
			"brief_desc": "",
			"contact_email": null,
			"contact_name": null,
			"contact_phone": null,
			"directory_name": "farmers market",
			"directory_type": "farmersmarket",
			"distance": "0.2492313843284663",
			"listing_desc": null,
			"listing_id": "312052",
			"listing_image": "default-farmersmarket-4-3.jpg",
			"listing_name": "Potsdam Farmers Market",
			"location_address": "3 Riverside Dr, New York, New York 10023",
			"location_city": "New York",
			"location_state": "New York",
			"location_street": "3 Riverside Dr",
			"location_x": "-73.98557883420489",
			"location_y": "40.7805342225485",
			"location_zipcode": "10023",
			"media_blog": null,
			"media_facebook": null,
			"media_instagram": null,
			"media_pinterest": null,
			"media_twitter": null,
			"media_website": null,
			"media_youtube": null,
			"mydesc": "",
			"term": "",
			"updatetime": "Mar 27th, 2023"
		},
		{
			"brief_desc": "",
			"contact_email": null,
			"contact_name": null,
			"contact_phone": null,
			"directory_name": "farmers market",
			"directory_type": "farmersmarket",
			"distance": "0.3345007893966277",
			"listing_desc": null,
			"listing_id": "311799",
			"listing_image": "default-farmersmarket-4-3.jpg",
			"listing_name": "79th Street Greenmarket",
			"location_address": "366 Columbus Ave, New York, New York 10024",
			"location_city": "New York",
			"location_state": "New York",
			"location_street": "366 Columbus Ave",
			"location_x": "-73.97616213937675",
			"location_y": "40.78102834727456",
			"location_zipcode": "10024",
			"media_blog": null,
			"media_facebook": null,
			"media_instagram": null,
			"media_pinterest": null,
			"media_twitter": null,
			"media_website": null,
			"media_youtube": null,
			"mydesc": "",
			"term": "",
			"updatetime": "Mar 27th, 2023"
		}
		]
	}`)

	multiRecordDataset, err := ParseUSDADataset(multiRecordDatasetRaw)
	require.Nil(t, err)
	require.Equal(t, len(multiRecordDataset), 2)
}
