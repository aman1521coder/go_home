package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	UploadDir   = "uploads/images"
	MaxFileSize = 5 * 1024 * 1024 // 5MB
	MaxImages   = 10               // Maximum images per item
)

// Allowed image MIME types
var allowedMimeTypes = map[string]bool{
	"image/jpeg": true,
	"image/jpg":  true,
	"image/png":  true,
	"image/gif":  true,
	"image/webp": true,
}

// ValidateImageFile validates an image file
func ValidateImageFile(fileHeader *multipart.FileHeader) error {
	// Validate file size
	if fileHeader.Size > MaxFileSize {
		return fmt.Errorf("file size exceeds maximum allowed size of 5MB")
	}

	if fileHeader.Size == 0 {
		return fmt.Errorf("file is empty")
	}

	// Open and validate MIME type
	file, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return fmt.Errorf("failed to read file: %w", err)
	}

	if n < 512 {
		buffer = buffer[:n]
	}

	mimeType := http.DetectContentType(buffer)
	if !allowedMimeTypes[mimeType] {
		return fmt.Errorf("invalid file type: %s. Allowed types: jpeg, jpg, png, gif, webp", mimeType)
	}

	return nil
}

// SaveUploadedImage saves an uploaded image file and returns the file path
func SaveUploadedImage(fileHeader *multipart.FileHeader, userID string) (string, error) {
	// Validate first
	if err := ValidateImageFile(fileHeader); err != nil {
		return "", err
	}

	// Open the uploaded file
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create upload directory if it doesn't exist
	if err := os.MkdirAll(UploadDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Generate unique filename
	ext := filepath.Ext(fileHeader.Filename)
	if ext == "" {
		// Default to .jpg if no extension
		ext = ".jpg"
	}
	filename := fmt.Sprintf("%s_%d%s", userID, time.Now().UnixNano(), ext)
	filePath := filepath.Join(UploadDir, filename)

	// Create the file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	// Copy file content
	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	// Return relative path for database storage
	return filePath, nil
}

// SaveMultipleImages saves multiple uploaded images
func SaveMultipleImages(fileHeaders []*multipart.FileHeader, userID string) ([]string, error) {
	if len(fileHeaders) > MaxImages {
		return nil, fmt.Errorf("maximum %d images allowed per item", MaxImages)
	}

	var imagePaths []string
	var savedPaths []string

	for _, fileHeader := range fileHeaders {
		imagePath, err := SaveUploadedImage(fileHeader, userID)
		if err != nil {
			// If one fails, delete already saved images
			for _, path := range savedPaths {
				DeleteImage(path)
			}
			return nil, err
		}
		imagePaths = append(imagePaths, imagePath)
		savedPaths = append(savedPaths, imagePath)
	}

	return imagePaths, nil
}

// DeleteImage deletes an image file
func DeleteImage(imagePath string) error {
	if imagePath == "" {
		return nil
	}
	if err := os.Remove(imagePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete image: %w", err)
	}
	return nil
}

// DeleteMultipleImages deletes multiple image files
func DeleteMultipleImages(imagePaths []string) error {
	for _, path := range imagePaths {
		if err := DeleteImage(path); err != nil {
			return err
		}
	}
	return nil
}

