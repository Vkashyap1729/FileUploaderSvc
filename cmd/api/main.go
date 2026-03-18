package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"file-uploader/internal/config"
    "file-uploader/internal/db"
    "file-uploader/internal/handler"
    "file-uploader/internal/repository"
    "file-uploader/internal/service"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Init DB
	postgres := db.NewPostgresDB(cfg)

	// Init layers
	repo := repository.NewUploadRepository(postgres)
	service := service.NewUploadService(repo)
	handler := handler.NewUploadHandler(service)

	// Setup router
	router := gin.Default()

	router.GET("/uploads", handler.GetUploads)

	log.Println("🚀 Server running on :8080")
	router.Run(":8080")
}