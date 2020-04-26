package config

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
)

func getConfig() Root {
	root := Root{}
	var configPath = flag.String("config", "", "The path to the config file")
	flag.Parse()
	fileBytes, err := ioutil.ReadFile(*configPath)
	if err != nil {
		fmt.Println(err)
		return root
	}
	err = xml.Unmarshal(fileBytes, &root)
	if err != nil {
		fmt.Println(err)
	}
	return root
}

// AppConfig represents the root element of the config file specified with --config
var AppConfig = getConfig()

// Root is the root of our config xml
type Root struct {
	XMLName xml.Name `xml:"Root"`
	Input   Input    `xml:"Input"`
	Output  Output   `xml:"Output"`
}

// Input is how we group all inputs
type Input struct {
	XMLName  xml.Name  `xml:"Input"`
	Patterns []Pattern `xml:"Pattern"`
}

// Pattern represents the source pattern we'll be collapsing
type Pattern struct {
	XMLName xml.Name `xml:"Pattern"`
	Path    string   `xml:"path,attr"`
	Size    int      `xml:"size,attr"`
}

// Output is how we group all outputs
type Output struct {
	XMLName xml.Name `xml:"Output"`
}
