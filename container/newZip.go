package container

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// NewContainer creates a new empty zip file to given folder with
// given uuid name and loops over given filename list calling
// addFilesToContainer function
func NewZip(uuid, dirName string, fileList []string) error {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	container, err := os.OpenFile(dirName+".zip", flags, 0644)
	if err != nil {
		return err
	}
	defer container.Close()
	zipwriter := zip.NewWriter(container)
	defer zipwriter.Close()

	for _, filename := range fileList {
		if err := addFilesToZip(filename, zipwriter); err != nil {
			return err
		}
	}

	return nil
}

// addFilesToContainer gets given filename and zipWriter to
// create files in the .zip container
// function is not exported
func addFilesToZip(filename string, zipwriter *zip.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	splitFilename := strings.Split(filename, string(os.PathSeparator))
	splitFilename = splitFilename[2:]
	joinFilename := filepath.Join(splitFilename...)
	wr, err := zipwriter.Create(joinFilename)
	if err != nil {
		return err
	}
	if _, err := io.Copy(wr, file); err != nil {
		return err
	}
	file.Close()
	return nil
}
