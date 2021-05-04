package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Kyps/container-api/container"
	"github.com/Kyps/container-api/utils"
)

// AddSignatureController gets param from URL query for the container which to where a signature is going to be added
// The function extracts the container, finds the correct name for new manifest and signature files. New manifest is created
// by creating new hashes from the files in the container. The container is rebuilt and old container and temp files are cleanud up.
func (c *Config) AddSignatureController(w http.ResponseWriter, req *http.Request) {
	keys, ok := req.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		log.Println("No params in query")
		ErrorNotFound(w, req, nil)
		return
	}
	key := keys[0]
	// Check if given container exists
	if _, err := os.Stat(c.DatabaseDir + key + ".zip"); os.IsNotExist(err) {
		ErrorNotFound(w, req, nil)
		return
	}
	// Extract container to temp directory
	err := container.ExtractZip(c.DatabaseDir+key+".zip", c.TempDir+key)
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	existingManifests := 1
	for {
		if _, err := os.Stat(c.TempDir + key + "/META-INF/manifest" + strconv.Itoa(existingManifests) + ".tlv"); os.IsNotExist(err) {
			break
		}
		existingManifests += 1
	}
	// Create new manifest file
	manifest, err := os.Create(c.TempDir + key + "/META-INF/manifest" + strconv.Itoa(existingManifests) + ".tlv")
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	// Get file list of temp folder, in loop create new hashes and append to new manifest
	filenames, err := ioutil.ReadDir(c.TempDir + key)
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	for _, file := range filenames {
		filePath := c.TempDir + key + "/" + file.Name()
		isDir, err := utils.IsDirectory(filePath)
		if err != nil {
			ErrorInternal(w, req, err)
			return
		}
		if isDir {
			break
		}
		openF, err := os.Open(filePath)
		if err != nil {
			ErrorInternal(w, req, err)
			return
		}
		fileHash, err := generateHash(openF)
		if err != nil {
			ErrorInternal(w, req, err)
			return
		}
		splitHash := strings.Split(fileHash, ":")
		if _, err := io.WriteString(manifest, "datafile\n\n\turi="+file.Name()+"\n\talgorithm="+splitHash[0]+"\n\thash:"+splitHash[1]+"\n"); err != nil {
			ErrorInternal(w, req, err)
			return
		}
		openF.Close()

	}
	if _, err := io.WriteString(manifest, "\nsignature-uri=/META-INF/signature"+strconv.Itoa(existingManifests)+".ksi"); err != nil {
		ErrorInternal(w, req, err)
		return
	}
	// Get signature to manifest file
	signature, err := signFile(manifest, c.GTURL, c.GTUser, c.GTPass)
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	manifest.Close()
	signatureFile, err := os.Create(c.TempDir + key + "/META-INF/signature" + strconv.Itoa(existingManifests) + ".ksi")
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	if _, err := io.WriteString(signatureFile, signature.String()); err != nil {
		ErrorInternal(w, req, err)
		return
	}

	signatureFile.Close()
	fileList, err := utils.GetFileList(c.TempDir + key)
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	err = os.Remove(c.DatabaseDir + key + ".zip")
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	// Create container
	err = container.NewZip(key, c.DatabaseDir+key, fileList)
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	err = os.RemoveAll(c.TempDir + key)
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":200, "message":"Signature added to container"}`))
}

// DeleteSignatureController gets param from URL query for the container which last signature is going to be deleted
// The fuction extracts the container, finds the last singature and manifest, then deletes them. After that the old
// container is deleted and the extracted container rebuilt. At the end the functions returns status - 200 and success message
func (c *Config) DeleteSignatureController(w http.ResponseWriter, req *http.Request) {
	keys, ok := req.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
		log.Println("No params in query")
		ErrorNotFound(w, req, nil)
		return
	}
	key := keys[0]
	if _, err := os.Stat(c.DatabaseDir + key + ".zip"); os.IsNotExist(err) {
		ErrorNotFound(w, req, nil)
		return
	}
	err := container.ExtractZip(c.DatabaseDir+key+".zip", c.TempDir+key)
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	// Check how many signatue files are in the container
	signatureList, err := utils.FileListMatch(c.TempDir+key+"/META-INF", "signature*")
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}

	// Find most recent manifest and signature file. This relies on numbers in filenames. In the future also time comparison must be added.
	signatureListLength := len(signatureList)
	if _, err := os.Stat(c.TempDir + key + "/META-INF/signature" + strconv.Itoa(signatureListLength) + ".ksi"); os.IsNotExist(err) {
		ErrorInternal(w, req, err)
		return
	}
	if _, err := os.Stat(c.TempDir + key + "/META-INF/manifest" + strconv.Itoa(signatureListLength) + ".tlv"); os.IsNotExist(err) {
		ErrorInternal(w, req, err)
		return
	}
	err = os.Remove(c.TempDir + key + "/META-INF/signature" + strconv.Itoa(signatureListLength) + ".ksi")
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	err = os.Remove(c.TempDir + key + "/META-INF/manifest" + strconv.Itoa(signatureListLength) + ".tlv")
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}

	fileList, err := utils.GetFileList(c.TempDir + key)
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	// Remove old container
	err = os.Remove(c.DatabaseDir + key + ".zip")
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	// Rebuild container
	err = container.NewZip(key, c.DatabaseDir+key, fileList)
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	// Clean up temp files
	err = os.RemoveAll(c.TempDir + key)
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":200, "message":"Last signature removed from container"}`))
}
