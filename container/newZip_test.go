package container

import (
	"archive/zip"
	"os"
	"testing"

	"github.com/Kyps/container-api/utils"
)

func TestNewZip(t *testing.T) {
	uuid := "test"
	filelist, err := utils.GetFileList("./testfiles/")
	if err != nil {
		t.Errorf("CreateContainer failed, expected filelist error nil, got %v", err)
	}
	path := "./testcontainer/test"
	err = NewZip(uuid, path, filelist)
	if err != nil {
		t.Errorf("CreateContainer failed, expected error nil, got %v", err)
	}
}
func TestAddFilesToZip(t *testing.T) {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	container, err := os.OpenFile("./testcontainer/addToZip.zip", flags, 0644)
	if err != nil {
		t.Errorf("addFilesToZip failed, expected error nil, got %v", err)
	}
	defer container.Close()
	zipwriter := zip.NewWriter(container)
	defer zipwriter.Close()
	filename := "testfiles" + string(os.PathSeparator) + "META-INF" + string(os.PathSeparator) + "manifest1.tlv"

	if err := addFilesToZip(filename, zipwriter); err != nil {
		t.Errorf("addFilesToZip failed, expected error nil, got %v", err)
	}
}
