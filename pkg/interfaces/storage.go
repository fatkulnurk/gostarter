package interfaces

import "context"

type UploadInput struct {
	FileName string
	Content  []byte
	MimeType string
}

type IStorage interface {
	Upload(ctx context.Context, input UploadInput) (string, error) // returns URL or path
	Delete(ctx context.Context, path string) error
}
