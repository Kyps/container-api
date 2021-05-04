package main

import "testing"

func TestConfig(t *testing.T) {
	configPath := "conf.json"

	result := GetConfig(configPath)
	if result.DatabaseDir == "" {
		t.Errorf("GetConfig failed, expected not empty DatabaseDir, got %v", result.DatabaseDir)
	}
	if result.TempDir == "" {
		t.Errorf("GetConfig failed, expected not empty TempDir, got %v", result.TempDir)
	}
	if result.Port == "" {
		t.Errorf("GetConfig failed, expected not empty Port, got %v", result.Port)
	}
	if result.GTURL == "" {
		t.Errorf("GetConfig failed, expected not empty GTURL, got %v", result.GTURL)
	}
	if result.GTUser == "" {
		t.Errorf("GetConfig failed, expected not empty GTUser, got %v", result.GTUser)
	}
	if result.GTPass == "" {
		t.Errorf("GetConfig failed, expected not empty GTPass, got %v", result.GTPass)
	}
}
