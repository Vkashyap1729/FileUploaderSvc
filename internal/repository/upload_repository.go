package repository

import (
    "context"
    "file-uploader/internal/db"
	"time"
)

type UploadRepository struct {
    db *db.PostgresDB
}

func NewUploadRepository(db *db.PostgresDB) *UploadRepository {
    return &UploadRepository{db: db}
}

func (r *UploadRepository) CreateUpload(ctx context.Context, id string, fileName string, size int64, key string) error {
    query := `
        INSERT INTO uploads (id, file_name, file_size, status, storage_key)
        VALUES ($1, $2, $3, $4, $5)
    `
    _, err := r.db.Pool.Exec(ctx, query, id, fileName, size, "initiated", key)
    return err
}

func (r *UploadRepository) GetAllUploads(ctx context.Context) ([]map[string]interface{}, error) {
	query := `
		SELECT id, file_name, file_size, status, storage_key, created_at
		FROM uploads
		ORDER BY created_at DESC
	`

	rows, err := r.db.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}

	for rows.Next() {
		var (
			id         string
			fileName   string
			fileSize   int64
			status     string
			storageKey string
			createdAt  time.Time
		)

		err := rows.Scan(&id, &fileName, &fileSize, &status, &storageKey, &createdAt)
		if err != nil {
			return nil, err
		}

		row := map[string]interface{}{
			"id":          id,
			"file_name":   fileName,
			"file_size":   fileSize,
			"status":      status,
			"storage_key": storageKey,
			"created_at":  createdAt,
		}

		results = append(results, row)
	}

	return results, nil
}