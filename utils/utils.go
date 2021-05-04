package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Generate UUID
// returns result string and error
func NewUuid() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random)
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

// Checks if given path is file or a directory
// returns true of false
func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}


// Function uses path parameter to walk the directory and return a file 
// list that contains aslo subdirectories
func GetFileList(path string) ([]string, error) {
	var filenames []string
	err := filepath.Walk(path,
		func(fpath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			finfo, err := os.Stat(fpath)
			if err != nil {
				return err
			}
			fMode := finfo.Mode()
			if !fMode.IsDir() {
				// fmt.Println(fpath, info.Size())
				filenames = append(filenames, fpath)
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	return filenames, nil
}


// Function uses path parameter to walk the directory and matches found filenames
// to the given pattern parameter. Finally returns a slice of strings containing 
// matched filepaths.
func FileListMatch(path, pattern string) ([]string, error) {
    var matches []string
    err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            return nil
        }
        if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
            return err
        } else if matched {
            matches = append(matches, path)
        }
        return nil
    })
    if err != nil {
        return nil, err
    }
    return matches, nil
}

