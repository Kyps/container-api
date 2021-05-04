package main

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration file template
type Config struct {
	Port        string
	TempDir     string
	DatabaseDir string
	GTUser      string
	GTPass      string
	GTURL       string
}

// Read json configuration file and return decoded json struct
func GetConfig(path string) Config {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Println(err)
	}
	return configuration
}
