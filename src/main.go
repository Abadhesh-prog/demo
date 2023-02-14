package main

import (
	"fmt"
	"io/ioutil"
    "github.com/Abadhesh-prog/demo"
	"gopkg.in/yaml.v3"
)

// type Config struct {
// 	HostName string `yaml:"hostname"`
// 	Port  int    `yaml:"port"`
// }
type APIComp struct {
	Prerequisites []struct {
		Name         string `yaml:"name"`
		Version     string `yaml:"version"`
		Downloadurl string    `yaml:"downloadurl"`
	} `yaml:"prerequisites"`
}

func main() {
	// Read the file
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a struct to hold the YAML data
	var config APIComp

	// Unmarshal the YAML data into the struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the data
	fmt.Println(config)
}