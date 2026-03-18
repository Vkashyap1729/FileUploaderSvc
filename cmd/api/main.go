package main

import (
    "file-uploader/internal/config"
    "file-uploader/internal/db"
)

func main() {
    cfg := config.LoadConfig()

    postgres := db.NewPostgresDB(cfg)

    _ = postgres // use later
}