package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type AdapterInterface interface {
	UploadToPinata(filePath string) (string, error)
}

type Adapter struct {
	cfg Config
}

func NewAdapter(cfg Config) AdapterInterface {
	return &Adapter{
		cfg: cfg,
	}
}

type UploadResponse struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	CID           string `json:"cid"`
	Size          string `json:"size"`
	NumberOfFiles string `json:"number_of_files"`
	MimeType      string `json:"mime_type"`
	UserId        string `json:"user_id"`
}

func (a *Adapter) UploadToPinata(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Create a buffer to store the request body
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Create the form file field
	formFile, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return "", fmt.Errorf("error creating form file: %w", err)
	}

	// Copy the file content into the form file field
	_, err = io.Copy(formFile, file)
	if err != nil {
		return "", fmt.Errorf("error copying file content: %w", err)
	}

	// Close the writer to finalize the form data
	writer.Close()

	// Create a new HTTP request
	url := "https://uploads.pinata.cloud/v3/files"
	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	// Set the headers
	req.Header.Set("Authorization", "Bearer "+a.cfg.PinataJWT)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a new HTTP client and perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error performing request: %w", err)
	}
	defer resp.Body.Close()

	// Decode the response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding response: %w", err)
	}

	// Check if the response status is OK
	if resp.StatusCode == http.StatusOK {
		log.Println(result)

		hash1, ok := result["data"].(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("error parsing IpfsHash from response")
		}
		cid, ok := hash1["cid"].(string)
		if !ok {
			return "", fmt.Errorf("error parsing IpfsHash from response")
		}
		return cid, nil
	}

	return "", fmt.Errorf("error uploading file: %s", result)
}
