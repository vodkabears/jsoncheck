package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mgutz/ansi"
)

// Version from ldflags
var version string

func readResources(config string) (*[]Resource, error) {
	file, err := ioutil.ReadFile(config)
	if err != nil {
		return nil, err
	}

	var resources []Resource
	err = json.Unmarshal(file, &resources)
	if err != nil {
		return nil, err
	}

	return &resources, nil
}

func printValidResult(result *ResourceResult) {
	color := "green"
	fmt.Println(ansi.Color("Status: valid", color+"+i"))
	fmt.Println(ansi.Color("Document: "+result.URL, color))
	fmt.Println(ansi.Color("Schema: "+result.Schema, color))
}

func printInvalidResult(result *ResourceResult) {
	color := "red"
	fmt.Println(ansi.Color("Status: invalid", color+"+i"))
	fmt.Println(ansi.Color("Document: "+result.URL, color))
	fmt.Println(ansi.Color("Schema: "+result.Schema, color))
	fmt.Println(ansi.Color("Errors:", color))
	for _, desc := range result.Errors() {
		fmt.Println(ansi.Color(" * "+desc.String(), color))
	}
}

func main() {
	showVersion := flag.Bool("version", false, "prints current version")
	flag.Usage = func() {
		fmt.Println("Usage: resource-validator [config]")
	}

	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	files := flag.Args()
	if len(files) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	resources, err := readResources(files[0])
	if err != nil {
		log.Fatal(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	chans := make([]chan *ResourceResult, len(*resources))
	for i, resource := range *resources {
		chans[i] = make(chan *ResourceResult)
		go resource.Check(wd, chans[i])
	}

	var hasErrors bool
	for _, ch := range chans {
		for result := range ch {
			if result.Valid() {
				printValidResult(result)
			} else {
				hasErrors = true
				printInvalidResult(result)
			}
			fmt.Println("--------------------------------")
		}
	}

	if hasErrors {
		os.Exit(1)
	}
}
