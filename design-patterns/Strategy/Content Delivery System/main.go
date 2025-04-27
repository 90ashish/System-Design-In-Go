package main

import (
	"fmt"
	"sync"
	"time"
)

// CompressionStrategy Interface
type CompressionStrategy interface {
	Compress(fileName string)
}

// ZIPCompression Strategy
type ZIPCompression struct{}

func (z *ZIPCompression) Compress(fileName string) {
	time.Sleep(2 * time.Second) // Simulate compression delay
	fmt.Printf("[ZIP] File '%s' compressed successfully using ZIP.\n", fileName)
}

// GZIPCompression Strategy
type GZIPCompression struct{}

func (g *GZIPCompression) Compress(fileName string) {
	time.Sleep(1 * time.Second) // Simulate compression delay
	fmt.Printf("[GZIP] File '%s' compressed successfully using GZIP.\n", fileName)
}

// LZCompression Strategy
type LZCompression struct{}

func (l *LZCompression) Compress(fileName string) {
	time.Sleep(3 * time.Second) // Simulate compression delay
	fmt.Printf("[LZ] File '%s' compressed successfully using LZ.\n", fileName)
}

// Context - Compression Manager
type CompressionManager struct {
	strategy CompressionStrategy
}

// SetStrategy - Assigns the chosen strategy
func (c *CompressionManager) SetStrategy(strategy CompressionStrategy) {
	c.strategy = strategy
}

// CompressFile - Executes the chosen compression strategy
func (c *CompressionManager) CompressFile(fileName string, wg *sync.WaitGroup) {
	defer wg.Done()

	if c.strategy == nil {
		fmt.Printf("Error: No compression strategy selected for file: %s\n", fileName)
		return
	}

	c.strategy.Compress(fileName)
}

func main() {
	var wg sync.WaitGroup

	// Create compression manager instances
	manager1 := &CompressionManager{}
	manager2 := &CompressionManager{}
	manager3 := &CompressionManager{}

	// Assign Strategies
	manager1.SetStrategy(&ZIPCompression{})
	manager2.SetStrategy(&GZIPCompression{})
	manager3.SetStrategy(&LZCompression{})

	// List of files to compress
	files := []struct {
		manager  *CompressionManager
		fileName string
	}{
		{manager1, "document.txt"},
		{manager2, "website_assets.js"},
		{manager3, "large_video.mp4"},
		{manager1, "report.pdf"},
		{manager2, "main.css"},
	}

	// Concurrent Compression using Goroutines
	for _, file := range files {
		wg.Add(1)
		go file.manager.CompressFile(file.fileName, &wg)
	}

	// Wait for all Goroutines to complete
	wg.Wait()

	fmt.Println("✅ All compression tasks have been completed successfully.")
}
