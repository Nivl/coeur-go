package httpsrv

import (
	"errors"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

// HTTPServer represents
type HTTPServer struct {
	Server *fiber.App
	cfg    *Config
}

// Config represents
type Config struct {
	CORS *cors.Config
	// Default: os.Getenv("HTTP_PORT")
	Port string
	// Default: os.Getenv("PPROF_PORT")
	PprofPort string
	// RateLimitMax is the amount of requests a client can do during
	// $RateLimitExpiration
	// Default:  20
	// Set to -1 to disable rate limiting
	RateLimitMax int
	// RateLimitExpiration is the time interval in which the amount of
	// requests is limited to RateLimitMax. If set to 10s, a client
	// will be limited to $RateLimitMax request per 10 seconds
	// Default: 5 seconds
	RateLimitExpiration time.Duration
	// RateLimitKey is a function that returns a unique key per client
	// this is used to know if 2 requests have been made by the same 2
	// clients.
	// Default:  c.IP()
	RateLimitKey func(c *fiber.Ctx) string
	// Fiber lets you control fiber more deeply
	Fiber fiber.Config
}

func NewServer(cfg *Config) (*fiber.App, error) {
	app := fiber.New()

	if cfg.Port == "" {
		cfg.Port = os.Getenv("HTTP_PORT")
	}
	if cfg.PprofPort == "" {
		cfg.PprofPort = os.Getenv("PPROF_PORT")
	}
	if cfg.Port == "" {
		return nil, errors.New("no port set")
	}

	// Rate Limiting
	if cfg.RateLimitMax > -1 {
		if cfg.RateLimitMax == 0 {
			cfg.RateLimitMax = 20
		}
		if cfg.RateLimitExpiration == 0 {
			cfg.RateLimitExpiration = 5
		}
		app.Use(limiter.New(limiter.Config{
			Max:          cfg.RateLimitMax,
			Expiration:   cfg.RateLimitExpiration,
			KeyGenerator: cfg.RateLimitKey,
		}))
	}

	// Pprof
	if cfg.PprofPort == cfg.Port {
		app.Use(pprof.New())
	}

	if cfg.CORS != nil {
		app.Use(cors.New(*cfg.CORS))
	}

	return app, nil
}
