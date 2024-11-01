package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jadidbourbaki/gofarm/api"
	"go.uber.org/zap"
)

// Service contains all the state we need for the Farmers' Market service
type Service struct {
	apis          map[string]api.FarmersMarketApi // a map from the datasource to a specific api
	logger        zap.Logger
	sugaredLogger zap.SugaredLogger
}

// loadApis loads all the available APIs for Farmers' Market Data
func (service *Service) loadApis() {
	service.apis = make(map[string]api.FarmersMarketApi)

	// Initialize the USDA data source
	usdaApi, err := api.NewUSDAFarmersMarketApi()
	if err != nil {
		service.sugaredLogger.Fatalf("could not load api: %w", err)
	}

	service.apis["usda"] = usdaApi

	newyorkApi, err := api.NewNewYorkMarketApi()
	if err != nil {
		service.sugaredLogger.Fatalf("could not load api: %w", err)
	}

	service.apis["newyork"] = newyorkApi

}

// New creates a new service struct and initializes loggers and the API as
// well as other internal state
func New() *Service {
	service := &Service{}

	// Initialize the logger.
	// No need to handle errors here, if even the logger isn't running
	// then... well.. there's not much we can do.
	service.logger = *zap.Must(zap.NewProduction())
	service.sugaredLogger = *service.logger.Sugar()

	service.loadApis()

	return service
}

func (service *Service) ApiForDataSource(dataSource string) (api.FarmersMarketApi, bool) {
	api, ok := service.apis[dataSource]
	return api, ok
}

// Run runs the service
func (service *Service) Run(port int, enableTls bool) {
	if err := loadDefaultViewAndTemplates(); err != nil {
		service.sugaredLogger.Fatalf("could not load default view: %w", err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/nearestNJson", service.nearestNJsonHandler).Methods("GET")
	router.HandleFunc("/nearestNHtml", service.nearestNHTMLHandler).Methods("GET")
	router.HandleFunc("/nearestNJsonByZipCode", service.nearestNJsonByZipCodeHandler).Methods("GET")
	router.HandleFunc("/nearestNHtmlByZipCode", service.nearestNByZipCodeHTMLHandler).Methods("GET")
	router.HandleFunc("/", service.getLocationHTMLHandler).Methods("GET")

	router.HandleFunc("/about", service.supportUsHandler).Methods("GET")

	router.HandleFunc("/developerResources/{path}", service.developerResourcesHandler).Methods("GET")
	router.HandleFunc("/developerResources", service.developerResourcesHandler).Methods("GET")

	runOnPort := fmt.Sprintf(":%v", port)
	service.sugaredLogger.Infof("running on port %s", runOnPort)

	if !enableTls {
		err := http.ListenAndServe(runOnPort, router)
		if err != nil {
			service.sugaredLogger.Fatal(err)
		}
		return
	}

	if err := defaultTlsCredentials.load(); err != nil {
		service.sugaredLogger.Fatal(err)
	}

	err := http.ListenAndServeTLS(runOnPort, defaultTlsCredentials.certificate, defaultTlsCredentials.key, router)
	if err != nil {
		service.sugaredLogger.Fatal(err)
	}
}

// Shutdown should be called prior to closing the service
func (service *Service) Shutdown() {
	service.logger.Sync()
}
