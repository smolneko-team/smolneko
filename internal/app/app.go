package app

import (
	"fmt"

	"github.com/smolneko-team/smolneko/config"
	v1 "github.com/smolneko-team/smolneko/internal/controller/http/v1"
	"github.com/smolneko-team/smolneko/internal/usecase"
	"github.com/smolneko-team/smolneko/internal/usecase/repo"
	"github.com/smolneko-team/smolneko/pkg/httpserver"
	"github.com/smolneko-team/smolneko/pkg/logger"
	"github.com/smolneko-team/smolneko/pkg/postgres"

	"github.com/gofiber/fiber/v2"
)

func Run(cfg *config.Config) {

	l := logger.New(cfg.Log.Level)

	url := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.SSLMode,
	)

	pg, err := postgres.New(url, postgres.MaxPoolSize(cfg.DB.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// TODO Use cases here
	figuresUseCase := usecase.NewFigures(repo.NewFiguresRepo(pg))
	charactersUseCase := usecase.NewCharacters(repo.NewCharactersRepo(pg))

	// TODO HTTP Server
	handler := fiber.New(httpserver.FiberConfig())
	v1.NewRouter(handler, l, figuresUseCase, charactersUseCase)
	httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
}
