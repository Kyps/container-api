package main

import (
	"encoding/json"
	"testing"
)

type Result struct {
	Results []string `json:"results"`
}

func TestGetAllFiles(t *testing.T) {
	conf := GetConfig("conf.json")
	results, err := conf.GetAllFiles()
	if err != nil {
		t.Errorf("GetAllFiles failed, expected error nil, got %v", err)
	}
	if len(results) == 0 {
		t.Errorf("GetAllFiles failed, expected length > 0, got %v", len(results))
	}
	var result Result
	err = json.Unmarshal(results, &result)
	if err != nil {
		t.Errorf("GetAllFiles failed to unmarhal []byte, expected error nil, got %v", err)
	}
}
