package test

import (
	"log"
	"testing"

	"github.com/yogaprasetya22/api-gotokopedia/internal/db"
	"github.com/yogaprasetya22/api-gotokopedia/internal/env"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"

	"context"
	"database/sql"
	"time"
)

type Config struct {
	Addr         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

func New(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Addr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)

	duration, err := time.ParseDuration(cfg.MaxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(duration)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func NewTestDB(t *testing.T) *sql.DB {
	addr := env.GetString("DB_ADDR", "postgresql://jagres:Jagres112.@localhost/gotokopedia?sslmode=disable")

	log.Println("Database address:", addr)

	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		t.Fatalf("Failed to connect to test DB: %v", err)
	}

	t.Cleanup(func() {
		conn.Close()
	})

	return conn
}

func NewTestStorage(t *testing.T) (*store.Storage, *sql.DB) {
	db := NewTestDB(t)
	storage := store.NewStorage(db)
	return &storage, db
}
func WithTransaction(t *testing.T, db *sql.DB, fn func(tx *sql.Tx) error) {
    tx, err := db.Begin()
    if err != nil {
        t.Fatalf("failed to begin transaction: %v", err)
    }

    var fnErr error

    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p)
        } else if t.Failed() || fnErr != nil {
            tx.Rollback()
        } else {
            if err := tx.Commit(); err != nil {
                t.Fatalf("failed to commit transaction: %v", err)
            }
        }
    }()

    fnErr = fn(tx)
}