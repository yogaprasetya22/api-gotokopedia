package main

import (
	"log"
	"os"

	"github.com/yogaprasetya22/api-gotokopedia/internal/db"
	"github.com/yogaprasetya22/api-gotokopedia/internal/env"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgresql://jagres:Jagres112.@localhost:5432/gotokopedia?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}
	// Custom EncoderConfig
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:      "time",
		LevelKey:     "level",
		MessageKey:   "msg",
		CallerKey:    "caller",
		EncodeTime:   zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"), // Format waktu yang lebih readable
		EncodeLevel:  zapcore.CapitalColorLevelEncoder,                   // Level log dengan warna
		EncodeCaller: zapcore.ShortCallerEncoder,                         // File dan line number
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // Output ke terminal
		zapcore.AddSync(os.Stdout),               // Kirim ke stdout
		zapcore.DebugLevel,                       // Level log minimal
	)

	logger := zap.New(core)
	defer logger.Sync()

	// Main Database
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		logger.Fatal(err.Error())
	}

	defer db.Close()
	logger.Info("database connection pool established")

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	if err := app.run(mux); err != nil {
		log.Fatal(err)
	}
}
