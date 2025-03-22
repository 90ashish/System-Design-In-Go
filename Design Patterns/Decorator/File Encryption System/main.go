package main

import (
	"fmt"
)

// FileHandler Interface
type FileHandler interface {
	Save(data string)
}

// BasicFileSaver - Core functionality for saving files
type BasicFileSaver struct{}

func (b *BasicFileSaver) Save(data string) {
	fmt.Printf("[BasicFileSaver] Data saved: %s\n", data)
}

// EncryptionDecorator - Adds Encryption Layer
type EncryptionDecorator struct {
	wrapped FileHandler
}

func (e *EncryptionDecorator) Save(data string) {
	encryptedData := fmt.Sprintf("🔒 [Encrypted] %s", data)
	fmt.Println("[Encryption Layer] Encrypting data...")
	e.wrapped.Save(encryptedData)
}

// CompressionDecorator - Adds Compression Layer
type CompressionDecorator struct {
	wrapped FileHandler
}

func (c *CompressionDecorator) Save(data string) {
	compressedData := fmt.Sprintf("📦 [Compressed] %s", data)
	fmt.Println("[Compression Layer] Compressing data...")
	c.wrapped.Save(compressedData)
}

// ChecksumDecorator - Adds Checksum Layer
type ChecksumDecorator struct {
	wrapped FileHandler
}

func (cs *ChecksumDecorator) Save(data string) {
	checksumData := fmt.Sprintf("✅ [Checksum] %s", data)
	fmt.Println("[Checksum Layer] Adding checksum to data...")
	cs.wrapped.Save(checksumData)
}

func main() {
	// Core file saver
	basicSaver := &BasicFileSaver{}

	// Example 1: File saved with Encryption only
	fmt.Println("\n🔹 Saving file with Encryption:")
	encryptedSaver := &EncryptionDecorator{wrapped: basicSaver}
	encryptedSaver.Save("User_Profile_Data")

	// Example 2: File saved with Compression and Encryption
	fmt.Println("\n🔹 Saving file with Compression + Encryption:")
	compressedEncryptedSaver := &CompressionDecorator{
		wrapped: &EncryptionDecorator{
			wrapped: basicSaver,
		},
	}
	compressedEncryptedSaver.Save("Website_Analytics_Report")

	// Example 3: File saved with Checksum, Compression, and Encryption
	fmt.Println("\n🔹 Saving file with Checksum + Compression + Encryption:")
	fullProtectionSaver := &ChecksumDecorator{
		wrapped: &CompressionDecorator{
			wrapped: &EncryptionDecorator{
				wrapped: basicSaver,
			},
		},
	}
	fullProtectionSaver.Save("Server_Logs_2023")
}
