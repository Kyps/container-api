package utils

import "testing"

func TestNewUuid(t *testing.T) {
	uuid, err := NewUuid()
	if err != nil {
		t.Errorf("NewUuid failed, expected error nil, got %v", err)
	}
	if uuid == "" {
		t.Errorf("NewUuid failed, expected uuid, got empty string")
	}
	if len(uuid) != 36 {
		t.Errorf("NewUuid failed, expected length 36 , got length %v", len(uuid))
	}
}

func TestIsDirectory(t *testing.T) {
	filePath := "../conf.json"
	dirPath := "../utils"

	fileTest, err := IsDirectory(filePath)
	if err != nil {
		t.Errorf("IsDirectory failed, expected error nil, got %v", err)
	}
	if fileTest {
		t.Errorf("IsDirectory failed, expected false, got %v", fileTest)
	}
	dirTest, err := IsDirectory(dirPath)
	if err != nil {
		t.Errorf("IsDirectory failed, expected error nil, got %v", err)
	}
	if !dirTest {
		t.Errorf("IsDirectory failed, expected true, got %v", dirTest)
	}
}

func TestGetFileList(t *testing.T) {
	dirPath := "../utils"

	dirTest, err := GetFileList(dirPath)
	if err != nil {
		t.Errorf("GetFileList failed, expected error nil, got %v", err)
	}
	if len(dirTest) != 2 {
		t.Errorf("GetFileList failed, expected length 2, got %v", len(dirPath))
	}
}

func TestFileListMatch(t *testing.T) {
	dirPath := "../utils"
	passPattern := "util*"
	failPattern := "match*"

	dirTestPass, err := FileListMatch(dirPath, passPattern)
	if err != nil {
		t.Errorf("FileListMatch failed, expected error nil, got %v", err)
	}
	if len(dirTestPass) != 2 {
		t.Errorf("GetFileList failed, expected length 2, got %v", len(dirPath))
	}
	dirTestFail, err := FileListMatch(dirPath, failPattern)
	if err != nil {
		t.Errorf("FileListMatch failed, expected error nil, got %v", err)
	}
	if len(dirTestFail) > 0 {
		t.Errorf("GetFileList failed, expected length 2, got %v", len(dirPath))
	}
}
