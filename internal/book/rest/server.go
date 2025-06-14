package rest

import (
	"github.com/ngoctrng/calendarium/internal/book"
	"github.com/ngoctrng/calendarium/pkg/config"
	"log/slog"
	"net/http"
	"strings"

	sentrygo "github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Router *echo.Echo
	Config *config.Config

	// storage adapters
	BookStore book.Storage
}

func New(options ...Options) (*Server, error) {
	s := Server{
		Router: echo.New(),
		Config: config.Empty,
	}

	for _, fn := range options {
		if err := fn(&s); err != nil {
			return nil, err
		}
	}

	s.RegisterGlobalMiddlewares()

	s.RegisterHealthCheck(s.Router.Group(""))
	s.RegisterBookRoutes(s.Router.Group("/api/books"))

	return &s, nil
}

func (s *Server) RegisterGlobalMiddlewares() {
	s.Router.Use(middleware.Recover())
	s.Router.Use(middleware.Secure())
	s.Router.Use(middleware.RequestID())
	s.Router.Use(middleware.Gzip())
	s.Router.Use(sentryecho.New(sentryecho.Options{Repanic: true}))

	// CORS
	if s.Config.AllowOrigins != "" {
		aos := strings.Split(s.Config.AllowOrigins, ",")
		s.Router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: aos,
		}))
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) RegisterHealthCheck(router *echo.Group) {
	router.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  http.StatusText(http.StatusOK),
			"message": "Service is up and running",
		})
	})
}

func (s *Server) handleError(c echo.Context, err error, status int) error {
	slog.Error(err.Error(), "request_id", s.requestID(c))

	if status >= http.StatusInternalServerError {
		if hub := sentryecho.GetHubFromContext(c); hub != nil {
			hub.WithScope(func(scope *sentrygo.Scope) {
				hub.CaptureException(err)
				scope.SetExtra("request_id", s.requestID(c))
				scope.SetExtra("status", status)
				scope.SetExtra("method", c.Request().Method)
				scope.SetExtra("path", c.Request().URL.Path)
				scope.SetExtra("query", c.Request().URL.RawQuery)
			})
		}
	}

	return c.JSON(status, map[string]string{
		"message": http.StatusText(status),
	})
}

func (s *Server) requestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}
