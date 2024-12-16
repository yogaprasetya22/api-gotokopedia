package main

import (
	"expvar"
	"os"
	"runtime"

	"github.com/yogaprasetya22/api-gotokopedia/internal/db"
	"github.com/yogaprasetya22/api-gotokopedia/internal/env"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const version = "0.1"

//	@title			Tokopedia API
//	@description	API for Tokopedia, a social network for gohpers
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath					/api/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {
	cfg := config{
		addr:   env.GetString("ADDR", ":8080"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgresql://jagres:Jagres112.@localhost:5432/socialjagres?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	// Custom EncoderConfig
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:          "time",
		LevelKey:         "level",
		MessageKey:       "msg",
		CallerKey:        "caller",
		EncodeTime:       zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"), // Format waktu yang lebih readable
		EncodeLevel:      zapcore.CapitalColorLevelEncoder,                   // Level log dengan warna
		EncodeCaller:     zapcore.ShortCallerEncoder,                         // File dan line number
		ConsoleSeparator: "\t",                                               // Separator antar field
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // Output ke terminal
		zapcore.AddSync(os.Stdout),               // Kirim ke stdout
		zapcore.DebugLevel,                       // Level log minimal
	)

	// Logger
	logger := zap.New(core).Sugar()
	defer logger.Sync()

	// Main Database
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("database connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
	}

	// matrucs collected
	expvar.NewString("version").Set(version)
	expvar.Publish("database", expvar.Func(func() interface{} {
		return db.Stats()
	}))
	expvar.Publish("goroutines", expvar.Func(func() interface{} {
		return runtime.NumGoroutine()
	}))

	mux := app.mount()
	logger.Fatal(app.run(mux))
}
