package main

import (
	"io"
	"os"

	"github.com/guardtime/goksi/hash"
	"github.com/guardtime/goksi/service"
	"github.com/guardtime/goksi/signature"
)

// Generate a hash for manifest file of an uploaded file
// It receives a file pointer and uses Guardtime KSI SDK to generate the hash
func generateHash(docFile *os.File) (string, error) {
	// Create document hash.
	hsr, err := hash.Default.New()
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(hsr, docFile); err != nil {
		return "", err
	}
	docHash, err := hsr.Imprint()
	if err != nil {
		return "", err
	}
	// log.Println(docHash)
	return docHash.String(), nil
}

// Function to create a signature for signature file. It receives a manifest file
// pointer and uses Guardtime KSI signing service to generate a signature
func signFile(filename *os.File, url, user, pass string) (*signature.Signature, error) {
	// Create document hash.
	hsr, err := hash.Default.New()
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(hsr, filename); err != nil {
		return nil, err
	}
	docHash, err := hsr.Imprint()
	if err != nil {
		return nil, err
	}
	// log.Println(docHash)
	// return docHash.String(), nil
	// Initialize signing service and sign the document hash.
	signer, err := service.NewSigner(service.OptEndpoint(url, user, pass))
	if err != nil {
		return nil, err
	}
	ksiSignature, err := signer.Sign(docHash)
	if err != nil {
		return nil, err
	}
	return ksiSignature, nil

}
