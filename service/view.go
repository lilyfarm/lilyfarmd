package service

import (
	"embed"
	"text/template"
)

//go:embed templates/*.html
var templateFS embed.FS

type View struct {
	// loaded is true only if the templates have been successfuly loaded
	// from embed FS
	loaded bool

	head           string
	headingAndMenu string

	nearestN           *template.Template
	getLocation        *template.Template
	supportUs          *template.Template
	developerResources *template.Template
}

func loadDefaultViewAndTemplates() error {
	if err := defaultView.load(); err != nil {
		return err
	}

	if err := defaultGenericPageTemplateData.load(); err != nil {
		return err
	}

	return nil
}

// load reads and parses all the templates and other files for the view
func (view *View) load() error {
	headBytes, err := templateFS.ReadFile("templates/head.html")

	if err != nil {
		return err
	}

	view.head = string(headBytes)

	headingAndMenuBytes, err := templateFS.ReadFile("templates/headingAndMenu.html")

	if err != nil {
		return err
	}

	view.headingAndMenu = string(headingAndMenuBytes)

	nearestNBytes, err := templateFS.ReadFile("templates/nearestN.html")

	if err != nil {
		return err
	}

	nearestNTemplate, err := template.New("nearestN").Parse(string(nearestNBytes))

	if err != nil {
		return err
	}

	view.nearestN = nearestNTemplate

	getLocationBytes, err := templateFS.ReadFile("templates/getLocation.html")

	if err != nil {
		return err
	}

	getLocationTemplate, err := template.New("getLocation").Parse(string(getLocationBytes))

	if err != nil {
		return err
	}

	view.getLocation = getLocationTemplate

	supportUsBytes, err := templateFS.ReadFile("templates/supportUs.html")

	if err != nil {
		return err
	}

	supportUsTemplate, err := template.New("supportUs").Parse(string(supportUsBytes))

	if err != nil {
		return err
	}

	view.supportUs = supportUsTemplate

	developerResourcesBytes, err := templateFS.ReadFile("templates/developerResources.html")

	if err != nil {
		return err
	}

	developerResourcesTemplate, err := template.New("developerResources").Parse(string(developerResourcesBytes))

	if err != nil {
		return err
	}

	view.developerResources = developerResourcesTemplate

	view.loaded = true

	return nil
}

var defaultView View
