package service

import (
	"fmt"

	"github.com/jadidbourbaki/gofarm/api"
)

type GenericPageTemplateData struct {
	HtmlHead       string
	HeadingAndMenu string
}

var defaultGenericPageTemplateData GenericPageTemplateData

func (data *GenericPageTemplateData) load() error {
	if !defaultView.loaded {
		return fmt.Errorf("default view not loaded")
	}

	data.HtmlHead = defaultView.head
	data.HeadingAndMenu = defaultView.headingAndMenu

	return nil
}

type NearestNTemplateData struct {
	GenericPageTemplateData // embedded struct
	Count                   int
	Records                 []api.FarmersMarketRecord
}

// NewNearestTemplateData returns a NearestNTemplateData struct
// pre-populated with the HtmlHead and the HeadingAndMenu. It
// returns an error if the templates have not been loaded yet.
func NewNearestNTemplateData(count int, records []api.FarmersMarketRecord) (NearestNTemplateData, error) {
	returnData := NearestNTemplateData{}

	if !defaultView.loaded {
		return returnData, fmt.Errorf("templates not loaded")
	}

	returnData.HtmlHead = defaultView.head
	returnData.HeadingAndMenu = defaultView.headingAndMenu
	returnData.Count = count
	returnData.Records = records

	return returnData, nil
}

type DeveloperResourcesTemplateData struct {
	GenericPageTemplateData // embedded struct
	MarkdownHTML            string
}

// NewDeveloperResourcesTemplateData returns a DeveloperResourceTemplateData
// struct pre-populated with the HtmlHead and the HeadingAndMenu. It
// takes the *already converted* (to HTML) markdown as its argument It
// returns an error if the templates have not been loaded yet.
func NewDeveloperResourcesTemplateData(markdownHtml string) (DeveloperResourcesTemplateData, error) {
	returnData := DeveloperResourcesTemplateData{}

	if !defaultView.loaded {
		return returnData, fmt.Errorf("templates not loaded")
	}

	returnData.HtmlHead = defaultView.head
	returnData.HeadingAndMenu = defaultView.headingAndMenu
	returnData.MarkdownHTML = markdownHtml

	return returnData, nil
}
