package container

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// extractContainer unzips the container from given path with the pathname and
// extracts it to the given target directory. Because it the folder hierarchy is
// known, the META-INF folder is created at the beginning with target direcotry.
func ExtractZip(archive, target string) error {
	container, err := zip.OpenReader(archive)
	if err != nil {
		// ErrorInternal(w, req, err)
		return err
	}

	if err := os.MkdirAll(target+"/META-INF", 0755); err != nil {
		return err
	}
	for _, file := range container.File {
		path := filepath.Join(target, file.Name)
		// fmt.Printf("%s\n", file.Name)
		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}
	container.Close()
	return nil
}
