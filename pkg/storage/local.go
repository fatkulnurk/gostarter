package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/fatkulnurk/gostarter/pkg/utils"
)

type LocalStorage struct {
	BasePath string
	BaseURL  string
}

// NewLocalStorage creates a new local storage instance
func NewLocalStorage(basePath, baseURL string) (*LocalStorage, error) {
	// Ensure the base path exists
	if err := os.MkdirAll(basePath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create base directory: %w", err)
	}

	// Ensure baseURL ends with a slash
	if baseURL != "" && !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}

	return &LocalStorage{
		BasePath: basePath,
		BaseURL:  baseURL,
	}, nil
}

// Upload saves a file to the local storage
func (s *LocalStorage) Upload(ctx context.Context, input UploadInput) (*UploadOutput, error) {
	// Create full path
	filePath := filepath.Join(s.BasePath, input.FileName)

	// Ensure directory exists
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	// Convert content to []byte
	var content []byte
	var size int64

	switch v := input.Content.(type) {
	case []byte:
		content = v
		size = int64(len(v))
	case string:
		content = []byte(v)
		size = int64(len(content))
	case io.Reader:
		var err error
		content, err = io.ReadAll(v)
		if err != nil {
			return nil, fmt.Errorf("failed to read content: %w", err)
		}
		size = int64(len(content))
	default:
		return nil, errors.New("unsupported content type")
	}

	// Write file
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return nil, fmt.Errorf("failed to write file: %w", err)
	}

	return &UploadOutput{
		Name:      filepath.Base(input.FileName),
		Path:      input.FileName,
		Size:      size,
		SizeHuman: utils.FormatSize(size),
	}, nil
}

// Delete removes a file from the local storage
func (s *LocalStorage) Delete(ctx context.Context, path string) error {
	filePath := filepath.Join(s.BasePath, path)

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil // File doesn't exist, consider it already deleted
	}

	// Delete the file
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

// Copy copies a file within the local storage
func (s *LocalStorage) Copy(ctx context.Context, sourcePath, destinationPath string) error {
	srcPath := filepath.Join(s.BasePath, sourcePath)
	dstPath := filepath.Join(s.BasePath, destinationPath)

	// Check if source file exists
	if _, err := os.Stat(srcPath); os.IsNotExist(err) {
		return fmt.Errorf("source file does not exist: %w", err)
	}

	// Ensure destination directory exists
	dstDir := filepath.Dir(dstPath)
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// Read source file
	content, err := os.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("failed to read source file: %w", err)
	}

	// Write to destination file
	if err := os.WriteFile(dstPath, content, 0644); err != nil {
		return fmt.Errorf("failed to write destination file: %w", err)
	}

	return nil
}

// Move moves a file within the local storage
func (s *LocalStorage) Move(ctx context.Context, sourcePath, destinationPath string) error {
	srcPath := filepath.Join(s.BasePath, sourcePath)
	dstPath := filepath.Join(s.BasePath, destinationPath)

	// Check if source file exists
	if _, err := os.Stat(srcPath); os.IsNotExist(err) {
		return fmt.Errorf("source file does not exist: %w", err)
	}

	// Ensure destination directory exists
	dstDir := filepath.Dir(dstPath)
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// Move the file
	if err := os.Rename(srcPath, dstPath); err != nil {
		return fmt.Errorf("failed to move file: %w", err)
	}

	return nil
}

// Get retrieves a file's content from the local storage
func (s *LocalStorage) Get(ctx context.Context, path string) ([]byte, error) {
	filePath := filepath.Join(s.BasePath, path)

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file does not exist: %w", err)
	}

	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return content, nil
}

// File gets information about a single file
func (s *LocalStorage) File(ctx context.Context, path string, expiryTempUrl *time.Duration) (*FileStorage, error) {
	filePath := filepath.Join(s.BasePath, path)

	// Get file info
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file does not exist: %w", err)
		}
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	// Determine mime type
	mimeType := "application/octet-stream"
	ext := filepath.Ext(path)
	if ext != "" {
		if mType := mime.TypeByExtension(ext); mType != "" {
			mimeType = mType
		}
	}

	// Create URL
	url := ""
	tempUrl := ""
	if s.BaseURL != "" {
		url = s.BaseURL + path
		tempUrl = url // For local storage, temp URL is the same as the regular URL
	}

	return &FileStorage{
		Name:         filepath.Base(path),
		Path:         path,
		Size:         fileInfo.Size(),
		SizeHuman:    utils.FormatSize(fileInfo.Size()),
		MimeType:     mimeType,
		LastModified: fileInfo.ModTime(),
		Visibility:   VisibilityPublic, // Local files are always considered public
		Url:          url,
		TempUrl:      tempUrl,
	}, nil
}

// Files lists files in a directory
func (s *LocalStorage) Files(ctx context.Context, dir string, expiryTempUrl *time.Duration) ([]FileStorage, error) {
	dirPath := filepath.Join(s.BasePath, dir)

	// Check if directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return []FileStorage{}, nil // Return empty list if directory doesn't exist
	}

	// Read directory
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var files []FileStorage
	for _, entry := range entries {
		if entry.IsDir() {
			continue // Skip directories
		}

		relPath := filepath.Join(dir, entry.Name())
		file, err := s.File(ctx, relPath, expiryTempUrl)
		if err != nil {
			return nil, err
		}

		files = append(files, *file)
	}

	// Sort files by name
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name < files[j].Name
	})

	return files, nil
}

// Directories lists subdirectories in a directory
func (s *LocalStorage) Directories(ctx context.Context, dir string) ([]string, error) {
	dirPath := filepath.Join(s.BasePath, dir)

	// Check if directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return []string{}, nil // Return empty list if directory doesn't exist
	}

	// Read directory
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var dirs []string
	for _, entry := range entries {
		if !entry.IsDir() {
			continue // Skip files
		}

		relPath := filepath.Join(dir, entry.Name())
		dirs = append(dirs, relPath)
	}

	// Sort directories by name
	sort.Strings(dirs)

	return dirs, nil
}

// Exists checks if a file exists
func (s *LocalStorage) Exists(ctx context.Context, path string) (bool, error) {
	filePath := filepath.Join(s.BasePath, path)

	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, fmt.Errorf("failed to check if file exists: %w", err)
}
