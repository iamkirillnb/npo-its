package handlers

import (
	"github.com/iamkirillnb/Packages/pkg/logger"
	"github.com/labstack/echo/v4"
	"npo-its/internal"
	"npo-its/internal/repos"
	"time"

	"net/http"
)

const (
	defaultServerReadTimeout  = 15 * time.Second
	defaultServerWriteTimeout = 30 * time.Second
)

type Server struct {
	config     *internal.ServerConfig
	BaseServer *http.Server
	BaseRouter *echo.Echo

	logger *logger.Logger
	db     *repos.DBRepo
}

func NewHandler(config *internal.ServerConfig, logger *logger.Logger, d *repos.DBRepo) *Server {

	defaultRouter := echo.New()

	srv := &Server{
		config:     config,
		BaseServer: nil,
		BaseRouter: defaultRouter,
		logger:     logger,
		db:         d,
	}

	srv.BaseServer = &http.Server{
		Addr:         srv.config.Address(),
		ReadTimeout:  defaultServerReadTimeout,
		WriteTimeout: defaultServerWriteTimeout,
	}

	return srv
}

func (s *Server) Start() {
	s.BaseRouter.GET("/", s.getFiveMaxMetrics)

	s.logger.Fatal(s.BaseRouter.StartServer(s.BaseServer))
}


func (s *Server) getFiveMaxMetrics(ctx echo.Context) error {
	metr, err := s.db.GetFiveMaxMetrics()
	if err != nil {
		return err
	}

	return ctx.JSON(200, metr)
}
