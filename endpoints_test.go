package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestPostEndpoint(t *testing.T) {
	file, err := os.Open("conf.json")
	if err != nil {
		t.Errorf("Post endpoint test failed, expected error not nil, got %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("data", filepath.Base("./conf.json"))
	if err != nil {
		t.Errorf("Post endpoint test failed, expected error not nil, got %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Errorf("Post endpoint test failed, expected error not nil, got %v", err)
	}

	err = writer.Close()
	if err != nil {
		t.Errorf("Post endpoint test failed, expected error not nil, got %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:4000/", body)
	if err != nil {
		t.Errorf("Post endpoint test failed, expected error not nil, got %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(handleIndex)

	handler.ServeHTTP(res, req)

	if http.StatusOK != res.Code {
		t.Errorf("GetAllFiles endpoint test failed, expected %v, got %v", http.StatusOK, res.Code)
	}
}

func TestPutEndpoint(t *testing.T) {
	sourceFile, err := os.Open("./container/testcontainer/aa80c3d0-6721-4167-8ae3-d8ad238e8ad7.zip")
	if err != nil {
		t.Errorf("Put endpoint test failed, expected error not nil, got %v", err)
	}
	defer sourceFile.Close()

	// Create new file
	newFile, err := os.Create("./database/aa80c3d0-6721-4167-8ae3-d8ad238e8ad7.zip")
	if err != nil {
		t.Errorf("Put endpoint test failed, expected error not nil, got %v", err)
	}
	_, err = io.Copy(newFile, sourceFile)
	if err != nil {
		t.Errorf("Put endpoint test failed, expected error not nil, got %v", err)
	}
	newFile.Close()
	req, err := http.NewRequest(http.MethodPut, "http://localhost:4000/sign/?key=aa80c3d0-6721-4167-8ae3-d8ad238e8ad7", nil)
	if err != nil {
		t.Errorf("Put endpoint test failed, expected error not nil, got %v", err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(handleSignatures)

	handler.ServeHTTP(res, req)

	if http.StatusOK != res.Code {
		t.Errorf("Put endpoint test failed, expected %v, got %v", http.StatusOK, res.Code)
	}
	e := os.Remove("./database/aa80c3d0-6721-4167-8ae3-d8ad238e8ad7.zip")
	if e != nil {
		t.Errorf("Put endpoint test failed, expected error not nil, got %v", err)
	}
}
func TestDeleteEndpoint(t *testing.T) {
	sourceFile, err := os.Open("./container/testcontainer/aa80c3d0-6721-4167-8ae3-d8ad238e8ad7.zip")
	if err != nil {
		t.Errorf("Delete endpoint test failed, expected error not nil, got %v", err)
	}
	defer sourceFile.Close()

	// Create new file
	newFile, err := os.Create("./database/aa80c3d0-6721-4167-8ae3-d8ad238e8ad7.zip")
	if err != nil {
		t.Errorf("Delete endpoint test failed, expected error not nil, got %v", err)
	}
	_, err = io.Copy(newFile, sourceFile)
	if err != nil {
		t.Errorf("Delete endpoint test failed, expected error not nil, got %v", err)
	}
	newFile.Close()
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:4000/sign/?key=aa80c3d0-6721-4167-8ae3-d8ad238e8ad7", nil)
	if err != nil {
		t.Errorf("Delete endpoint test failed, expected error not nil, got %v", err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(handleSignatures)

	handler.ServeHTTP(res, req)

	if http.StatusOK != res.Code {
		t.Errorf("Delete endpoint test failed, expected %v, got %v", http.StatusOK, res.Code)
	}
	e := os.Remove("./database/aa80c3d0-6721-4167-8ae3-d8ad238e8ad7.zip")
	if e != nil {
		t.Errorf("Delete endpoint test failed, expected error not nil, got %v", err)
	}
}
func TestGetAllFilesEndpoint(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:4000/getall", nil)
	if err != nil {
		t.Errorf("GetAllFiles endpoint test failed, expected error not nil, got %v", err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(handleGetAll)

	handler.ServeHTTP(res, req)

	if http.StatusOK != res.Code {
		t.Errorf("GetAllFiles endpoint test failed, expected %v, got %v", http.StatusOK, res.Code)
	}
}
