package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jadidbourbaki/gofarm/geography"
)

// nearestNInternal returns the records as its first return value and true as its second return value if
// we were able to get the nearestN records. Otherwise, it returns a false as its second return value.
func (service *Service) nearestNInternal(w http.ResponseWriter, r *http.Request) (*NearestNTemplateData, bool) {
	query := r.URL.Query()

	nString := query.Get("n")
	n, err := strconv.Atoi(nString)
	if err != nil {
		service.sugaredLogger.Errorf("incorrect value for n: %w", err)
		w.WriteHeader(http.StatusBadRequest)
		return nil, false

	}

	latitudeString := query.Get("latitude")
	latitude, err := strconv.ParseFloat(latitudeString, 64)

	if err != nil {
		service.sugaredLogger.Errorf("incorrect value for latitude: %w", err)
		w.WriteHeader(http.StatusBadRequest)
		return nil, false
	}

	longitudeString := query.Get("longitude")
	longitude, err := strconv.ParseFloat(longitudeString, 64)

	if err != nil {
		service.sugaredLogger.Errorf("incorrect value for longitude: %w", err)
		w.WriteHeader(http.StatusBadRequest)
		return nil, false
	}

	datasourceString := query.Get("datasource")
	api, ok := service.ApiForDataSource(datasourceString)

	if !ok {
		service.sugaredLogger.Errorf("could not find api for datasource: %s", datasourceString)
		w.WriteHeader(http.StatusBadRequest)
		return nil, false
	}

	point := geography.HaversinePoint{Latitude: latitude, Longitude: longitude}
	records, err := api.NearestN(n, point)

	if err != nil {
		service.sugaredLogger.Errorf("nearestN: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil, false
	}

	data, err := NewNearestNTemplateData(n, records)
	if err != nil {
		service.sugaredLogger.Errorf("loading nearestN template data: %w", err)
	}

	return &data, true
}

func (service *Service) nearestNJsonHandler(w http.ResponseWriter, r *http.Request) {
	data, ok := service.nearestNInternal(w, r)

	if !ok {
		return
	}

	recordsJson, err := json.Marshal(data.Records)
	if err != nil {
		service.sugaredLogger.Errorf("marshalling json: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(recordsJson)
}

func (service *Service) nearestNHTMLHandler(w http.ResponseWriter, r *http.Request) {
	data, ok := service.nearestNInternal(w, r)

	if !ok {
		return
	}

	err := defaultView.nearestN.Execute(w, data)
	if err != nil {
		service.sugaredLogger.Errorf("executing template: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
