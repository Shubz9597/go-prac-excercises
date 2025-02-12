package handlers

import (
	"encoding/json"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type pathConfig struct {
	Path string `json: "path" yaml: "path"`
	Url  string `json: "url" yaml: "url"`
}

type yamlConfig struct {
	Paths []pathConfig `json: "paths" yaml: "paths"`
}

var Config yamlConfig

func ParseFileAndReturn(fileName string, isYaml bool) {
	var file string
	if isYaml {
		file = fileName + ".yaml"
	} else {
		file = fileName + ".json"
	}

	reader, err := os.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	if isYaml {
		err = yaml.Unmarshal(reader, &Config)
	} else {
		err = json.Unmarshal(reader, &Config)
	}

	if err != nil {
		log.Fatal(err)
	}
}
