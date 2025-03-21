package main

import (
	"fmt"
	"sync"
	"time"
)

// StorageService Interface
type StorageService interface {
	UploadFile(fileName string)
}

// GoogleDrive Struct
type GoogleDrive struct{}

func (g *GoogleDrive) UploadFile(fileName string) {
	time.Sleep(1 * time.Second) // Simulate upload delay
	fmt.Printf("[Google Drive] File '%s' uploaded successfully.\n", fileName)
}

// Dropbox Struct
type Dropbox struct{}

func (d *Dropbox) UploadFile(fileName string) {
	time.Sleep(1 * time.Second) // Simulate upload delay
	fmt.Printf("[Dropbox] File '%s' uploaded successfully.\n", fileName)
}

// AWS S3 Struct
type AWSS3 struct{}

func (a *AWSS3) UploadFile(fileName string) {
	time.Sleep(1 * time.Second) // Simulate upload delay
	fmt.Printf("[AWS S3] File '%s' uploaded successfully.\n", fileName)
}

// StorageFactory - Factory for creating storage services
type StorageFactory struct{}

// CreateStorageService - Factory method to generate the correct storage service
func (f *StorageFactory) CreateStorageService(service string) (StorageService, error) {
	switch service {
	case "googledrive":
		return &GoogleDrive{}, nil
	case "dropbox":
		return &Dropbox{}, nil
	case "awss3":
		return &AWSS3{}, nil
	default:
		return nil, fmt.Errorf("unsupported storage service: %s", service)
	}
}

// UploadFileConcurrently - Handles concurrent file uploads using Goroutines
func UploadFileConcurrently(wg *sync.WaitGroup, factory *StorageFactory, service string, fileName string) {
	defer wg.Done()

	storageService, err := factory.CreateStorageService(service)
	if err != nil {
		fmt.Printf("[ERROR] %v\n", err)
		return
	}
	storageService.UploadFile(fileName)
}

func main() {
	factory := &StorageFactory{}
	var wg sync.WaitGroup

	// Bulk File Uploads
	files := []struct {
		service  string
		fileName string
	}{
		{"googledrive", "project_report.pdf"},
		{"dropbox", "presentation.pptx"},
		{"awss3", "backup_data.zip"},
		{"googledrive", "design_sketch.jpg"},
		{"onedrive", "confidential_notes.docx"}, // Unsupported Service
	}

	// Concurrently upload files
	for _, file := range files {
		wg.Add(1)
		go UploadFileConcurrently(&wg, factory, file.service, file.fileName)
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	fmt.Println("✅ All file uploads have been processed successfully.")
}
