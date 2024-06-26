package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/mirzahilmi/locavest-backend/internal/delivery/rest"
	"github.com/mirzahilmi/locavest-backend/internal/pkg/cfg"
	"github.com/mirzahilmi/locavest-backend/internal/repository"
	"github.com/mirzahilmi/locavest-backend/internal/usecase"
	"github.com/rs/zerolog"
)

var log = zerolog.
	New(os.Stderr).
	Level(zerolog.DebugLevel).
	With().
	Timestamp().
	Logger()

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Caller().Msg("")
	}
}

func main() {
	r := cfg.NewEcho(&log)
	api := r.Group("/api")
	db := cfg.NewMariaDB()

	rest.RegisterUtilHandler(api)
	cartRepo := repository.NewCartRepository(db)
	cartUsecase := usecase.NewcartUsecase(cartRepo)
	rest.RegisterCartHandler(api, cartUsecase)

	ctx, exit := signal.NotifyContext(context.Background(), os.Interrupt)
	defer exit()
	go func() {
		if err := r.Start(":" + os.Getenv("API_PORT")); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Caller().Msg("failed to start server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Caller().Msg("failed to close server")
	}
}
