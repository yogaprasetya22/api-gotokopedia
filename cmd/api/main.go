package main

import (
	"expvar"
	"os"
	"runtime"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/yogaprasetya22/api-gotokopedia/internal/auth"
	"github.com/yogaprasetya22/api-gotokopedia/internal/db"
	"github.com/yogaprasetya22/api-gotokopedia/internal/env"
	"github.com/yogaprasetya22/api-gotokopedia/internal/mailer"
	"github.com/yogaprasetya22/api-gotokopedia/internal/ratelimiter"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store/cache"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const version = "0.1"

//	@title			Tokopedia API
//	@description	API ini menggunakan mekanisme rate limiter untuk membatasi jumlah permintaan yang dapat dilakukan oleh setiap klien dalam jangka waktu tertentu.
//	@description	Hal ini bertujuan untuk mencegah penyalahgunaan API, memastikan kinerja yang stabil, dan melindungi server dari potensi serangan seperti brute force atau DDoS.
//	@description	Rate limiter ini didukung oleh Redis sebagai penyimpanan sementara yang cepat dan andal untuk melacak jumlah permintaan setiap klien berdasarkan API key atau alamat IP mereka.
//	@description	Redis digunakan karena kemampuannya yang tinggi dalam menangani data secara real-time dengan latensi rendah.
//	@description	Jika batas permintaan terlampaui, klien akan menerima respons dengan kode status 429 (Too Many Requests).

//	@termsOfService	http://www.myogaprasetya.my.id

//	@contact.name	API Support
//	@contact.url	http://www.myogaprasetya.my.id
//	@contact.email	mochammad.yogaprasetya112@gmail.com

// @BasePath					/api/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	cfg := config{
		addr:        env.GetString("ADDR", ":8080"),
		apiURL:      env.GetString("EXTERNAL_URL", "localhost:8080"),
		frontendURL: env.GetString("frontendURL", "http://localhost:3000"), // email verification
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgresql://jagres:Jagres112.@localhost:5432/socialjagres?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		redisCfg: redisConfig{
			addr:    env.GetString("REDIS_ADDR", "localhost:6379"),
			pw:      env.GetString("REDIS_PW", ""),
			db:      env.GetInt("REDIS_DB", 0),
			enabled: env.GetBool("REDIS_ENABLED", false),
		},
		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp:       time.Hour * 24 * 3, // 3 days
			fromEmail: env.GetString("FROM_EMAIL", ""),
			host:      env.GetString("MAIL_HOST", "email-smtp.ap-southeast-1.amazonaws.com"),
			port:      env.GetInt("MAIL_PORT", 587),
			username:  env.GetString("MAIL_USERNAME", "AKIAUJ3VUEHXL7D4SEVL"),
			password:  env.GetString("MAIL_PASSWORD", "BN69fKzVY1wIx1yV8TtIeOIt/P25MLmHzjzsMUMH5B7d"),
			timeout:   time.Second * 5,
			sender:    env.GetString("MAIL_SENDER", ""),
		},
		rateLimiter: ratelimiter.Config{
			RequestsPerTimeFrame: env.GetInt("RATELIMITER_REQUESTS_COUNT", 20),
			TimeFrame:            time.Second * 5,
			Enabled:              env.GetBool("RATE_LIMITER_ENABLED", true),
		},
		auth: authConfig{
			basic: basicConfig{
				usrname: env.GetString("AUTH_BASIC_USRNAME", "jagresuye"),
				pass:    env.GetString("AUTH_BASIC_PASS", "asdasdasd"),
			},
			token: tokenConfig{
				secret: env.GetString("AUTH_TOKEN_SECRET", "jagreskuy112"),
				exp:    time.Hour * 24 * 3, // 3 days
				iss:    "jagressocial",
			},
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

	// Mailer
	config := mailer.MailerConfig{
		Timeout:      cfg.mail.timeout,
		Host:         cfg.mail.host,
		Port:         cfg.mail.port,
		Username:     cfg.mail.username,
		Password:     cfg.mail.password,
		Sender:       cfg.mail.sender,
		TemplatePath: "internal/mailer/template/",
	}
	mailer := mailer.New(config)

	// Redis Cache
	var rdb *redis.Client
	if cfg.redisCfg.enabled {
		rdb = cache.NewRedisClient(cfg.redisCfg.addr, cfg.redisCfg.pw, cfg.redisCfg.db)
		logger.Info("redis cache connection established")
		defer rdb.Close()
	}

	// Rate limiter
	rateLimiter := ratelimiter.NewFixedWindowLimiter(
		cfg.rateLimiter.RequestsPerTimeFrame,
		cfg.rateLimiter.TimeFrame,
	)

	// JWT Authenticator
	jwtAuthenticator := auth.NewJWTAuthenticator(cfg.auth.token.secret, cfg.auth.token.iss, cfg.auth.token.iss)

	// Store
	cacheStorage := cache.NewRedisStore(rdb)
	store := store.NewStorage(db)

	app := &application{
		config:        cfg,
		store:         store,
		cacheStorage:  cacheStorage,
		logger:        logger,
		mailer:        mailer,
		authenticator: jwtAuthenticator,
		rateLimiter:   rateLimiter,
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
