package db

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/jackc/pgx/v5/pgxpool"
    "your_project/internal/config"
)

type PostgresDB struct {
    Pool *pgxpool.Pool
}

func NewPostgresDB(cfg *config.Config) *PostgresDB {
    dsn := fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?sslmode=%s",
        cfg.DBUser,
        cfg.DBPassword,
        cfg.DBHost,
        cfg.DBPort,
        cfg.DBName,
        cfg.DBSSLMode,
    )

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    pool, err := pgxpool.New(ctx, dsn)
    if err != nil {
        log.Fatalf("Unable to create DB pool: %v", err)
    }

    // Ping DB
    if err := pool.Ping(ctx); err != nil {
        log.Fatalf("Unable to connect to DB: %v", err)
    }

    log.Println("✅ Connected to PostgreSQL")

    return &PostgresDB{Pool: pool}
}