package container

import (
	"os"
	"testing"
)

func TestExtractContainer(t *testing.T) {
	filePath := "./testcontainer/43c1bc6d-2d6e-4d10-8d6b-3bac98b550c1.zip"
	destinationPath := "./testfiles"
	err := ExtractZip(filePath, destinationPath)
	if err != nil {
		t.Errorf("CreateContainer failed, expected error nil, got %v", err)
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("CreateContainer failed, expected container exists, got %v", err)
	}
}
