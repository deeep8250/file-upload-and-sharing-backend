package models

import "time"

type File struct {
	ID          uint      `gorm:"primaryKey; autoIncrement" json:"id"`
	OwnerID     uint      `json:"owner_id"`
	FileName    string    `json:"file_name"`
	BucketName  string    `json:"bucket_name"`
	ObjectKey   string    `json:"object_key"`
	SizeBytes   int64     `json:"size_bytes"`
	ContentType string    `json:"content_type"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UploadedAt  time.Time `json:"uploaded_at"`
}
