package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jadidbourbaki/gofarm/docs"
)

func (service *Service) getLocationHTMLHandler(w http.ResponseWriter, _ *http.Request) {
	err := defaultView.getLocation.Execute(w, defaultGenericPageTemplateData)
	if err != nil {
		service.sugaredLogger.Errorf("executing template: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (service *Service) supportUsHandler(w http.ResponseWriter, _ *http.Request) {
	err := defaultView.supportUs.Execute(w, defaultGenericPageTemplateData)
	if err != nil {
		service.sugaredLogger.Errorf("executing template: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (service *Service) developerResourcesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path, ok := vars["path"]

	if !ok {
		// If no path specified, just load the main.md markdown file
		path = "main.md"
	}

	markdownHtml, err := docs.RenderHTML(path)
	if err != nil {
		service.sugaredLogger.Errorf("rendering html from markdown: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	developerResourcesTemplate, err := NewDeveloperResourcesTemplateData(markdownHtml)
	if err != nil {
		service.sugaredLogger.Errorf("creating template: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = defaultView.developerResources.Execute(w, developerResourcesTemplate)
	if err != nil {
		service.sugaredLogger.Errorf("executing template: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
