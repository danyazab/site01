package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"site01/config"
	_ "site01/docs"
	"site01/internal"
	"time"
)

func main() {
	cfg := &config.Database{
		User:     "postgres",
		Password: "secret",
		Host:     "127.0.0.1",
		Port:     8088,
		Name:     "site01",
	}

	err := internal.Invoke(internal.RunServer, cfg.Cgf)
	if err != nil {
		log.Fatal(err.Error())
	}

	e := echo.New()
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:      "1; mode=block",
		ContentTypeNosniff: "nosniff",
		XFrameOptions:      "SAMEORIGIN",
		HSTSMaxAge:         3600,
	}))

	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(30)))

	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 30, Burst: 30, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}
	e.Use(middleware.RateLimiterWithConfig(config))
}

// @title           Site01 Service
// @version         1.0
// @description     It is Swagger of rest api Site01.

// @contact.name   Daniel Zab
// @contact.email  zabolotnijdanilo@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8088
