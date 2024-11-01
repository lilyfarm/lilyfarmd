package service

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// nearestNByZipCodeInternal returns the records as its first return value and true as its second return value if
// we were able to get the nearestN records. Otherwise, it returns a false as its second return value.
func (service *Service) nearestNByZipCodeInternal(w http.ResponseWriter, r *http.Request) (*NearestNTemplateData, bool) {
	query := r.URL.Query()

	nString := query.Get("n")
	n, err := strconv.Atoi(nString)
	if err != nil {
		service.sugaredLogger.Errorf("incorrect value for n: %w", err)
		w.WriteHeader(http.StatusBadRequest)
		return nil, false

	}

	zipCodeString := query.Get("zipCode")

	// This is for validation, we will not be using the value
	// for this.
	_, err = strconv.Atoi(zipCodeString)

	if err != nil {
		service.sugaredLogger.Errorf("incorrect value for zipcode: %w", err)
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

	records, err := api.NearestNByZipCode(n, zipCodeString)
	if err != nil {
		service.sugaredLogger.Errorf("nearestNByZipCode: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil, false
	}

	data, err := NewNearestNTemplateData(n, records)
	if err != nil {
		service.sugaredLogger.Errorf("loading nearestN template data: %w", err)
	}

	return &data, true
}

// nearestNJsonByZipCodeHandler returns the nearest N Farmers' Markets to a particular zip code
// in JSON format
func (service *Service) nearestNJsonByZipCodeHandler(w http.ResponseWriter, r *http.Request) {
	data, ok := service.nearestNByZipCodeInternal(w, r)

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

// nearestNByZipCodeHTMLHandler returns the nearest N Farmers' Markets to a particular zip code
// in HTML format
func (service *Service) nearestNByZipCodeHTMLHandler(w http.ResponseWriter, r *http.Request) {
	data, ok := service.nearestNByZipCodeInternal(w, r)

	if !ok {
		return
	}

	err := defaultView.nearestN.Execute(w, data)
	if err != nil {
		service.sugaredLogger.Errorf("executing template: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
