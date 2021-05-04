package main

import (
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/Kyps/container-api/utils"
)

// Controller for GET request
// Returns JSON object that contains an array of all container names in database
func (c *Config) GetAllFiles() ([]byte, error) {
	files, err := utils.GetFileList(c.DatabaseDir)
	if err != nil {
		return nil, err
	}
	var filesArray []string
	for _, f := range files {
		filename := filepath.Base(f)
		name := strings.TrimSuffix(filename, filepath.Ext(filename))
		filesArray = append(filesArray, name)
	}
	dataToJson := map[string]interface{}{"results": filesArray}

	res, _ := json.Marshal(dataToJson)
	return res, nil
}
