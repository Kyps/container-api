package main

import (
	"os"
	"testing"
)

var tests = []struct {
	input string
}{
	{"conf.json"},
	{"ksi.go"},
	{"main.go"},
}

func TestGenerateHash(t *testing.T) {
	for _, i := range tests {
		openFile, err := os.Open(i.input)
		if err != nil {
			t.Errorf("Opening test file failed, expected error nil, got %v", err)
		}
		defer openFile.Close()
		fileHash, err := generateHash(openFile)
		if err != nil {
			t.Errorf("generateHash failed, expected error nil, got %v", err)
		}
		if fileHash == "" {
			t.Errorf("generateHash failed, expected result not empty string, got %v", err)
		}
	}

}

func TestSignFile(t *testing.T) {
	conf := GetConfig("conf.json")
	for _, i := range tests {
		openFile, err := os.Open(i.input)
		if err != nil {
			t.Errorf("Opening test file failed, expected error nil, got %v", err)
		}
		defer openFile.Close()
		fileHash, err := signFile(openFile, conf.GTURL, conf.GTUser, conf.GTPass)
		if err != nil {
			t.Errorf("generateHash failed, expected error nil, got %v", err)
		}
		if fileHash.String() == "" {
			t.Errorf("signFile failed, expected result not empty string, got %v", err)
		}
	}

}
