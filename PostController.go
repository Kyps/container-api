package main

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Kyps/container-api/container"
	"github.com/Kyps/container-api/utils"
)

// PostController handles incoming POST request where user sends their files to be signed and containerised
// This function reads multipart form data and writes it to a temporary folder with an uuid where a manifest
// is generated and then also signed using KSI Blockchain. After that received files are copied to a .zip
// container and finally temporary folder is removed.
// Manual file closes are used because temporary files cleanup is needed at the end on the function.
func (c *Config) PostController(w http.ResponseWriter, req *http.Request) {
	uuid, err := utils.NewUuid()
	// create temporary folder tree
	os.MkdirAll(c.TempDir+uuid+"/META-INF", 0755)
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	// Create manifest to META-INF directory
	manifest, err := os.Create(c.TempDir + uuid + "/META-INF/manifest1.tlv")
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	// Set how big chunks are read to memory
	err = req.ParseMultipartForm(10000000)
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	m := req.MultipartForm
	files := m.File["data"]
	// Read files from multipart form
	for i := range files {
		file, err := files[i].Open()
		if err != nil {
			ErrorInternal(w, req, err)
			return
		}
		defer file.Close()
		dst, err := os.Create(c.TempDir + uuid + "/" + files[i].Filename)
		if err != nil {
			ErrorInternal(w, req, err)
			return
		}
		if _, err := io.Copy(dst, file); err != nil {
			ErrorInternal(w, req, err)
			return
		}
		// Generating file hash for manifest and write a file entry
		fileHash, err := generateHash(dst)
		if err != nil {
			ErrorInternal(w, req, err)
			return
		}
		splitHash := strings.Split(fileHash, ":")
		if _, err := io.WriteString(manifest, "datafile\n\n\turi="+files[i].Filename+"\n\talgorithm="+splitHash[0]+"\n\thash:"+splitHash[1]+"\n"); err != nil {
			ErrorInternal(w, req, err)
			return
		}
		dst.Close()
	}
	if _, err := io.WriteString(manifest, "\nsignature-uri=/META-INF/signature1.ksi"); err != nil {
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
	signatureFile, err := os.Create(c.TempDir + uuid + "/META-INF/signature1.ksi")
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	if _, err := io.WriteString(signatureFile, signature.String()); err != nil {
		ErrorInternal(w, req, err)
		return
	}
	signatureFile.Close()
	// Generate file list for container function
	fileList, err := utils.GetFileList(c.TempDir + uuid)
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	// Create container
	container.NewZip(uuid, c.DatabaseDir+uuid, fileList)
	// Remove temporary directory
	err = os.RemoveAll(c.TempDir + uuid)
	if err != nil {
		ErrorInternal(w, req, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":200, "message":"Document(s) successfully signed"}`))
}
