package main

import (
	"log"

	"github.com/xeipuuv/gojsonschema"
)

// Resource contains test data and schema
type Resource struct {
	Schema string   `json:"schema"`
	Data   []string `json:"data"`
}

// ResourceResult contains validation results
type ResourceResult struct {
	*gojsonschema.Result
	URL    string
	Schema string
}

// Check validates resource by its schema
func (r Resource) Check(wd string, ch chan *ResourceResult) {
	schema := r.Schema
	schemaLoader := gojsonschema.NewReferenceLoader("file://" + wd + "/" + schema)
	for _, url := range r.Data {
		result, err := gojsonschema.Validate(schemaLoader, gojsonschema.NewReferenceLoader(url))
		if err != nil {
			log.Printf("Error (%s; %s): %s\n", schema, url, err)
			continue
		}

		ch <- &ResourceResult{
			result,
			url,
			schema,
		}
	}

	close(ch)
}
