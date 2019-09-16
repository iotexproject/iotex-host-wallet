package controller

import (
	"github.com/iotexproject/iotex-host-wallet/wallet/config"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type Status struct {
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
	Error     string `json:"error,omitempty"`
}

// HealthRoute address router
func HealthRoute(e *echo.Echo) {
	e.GET("/health", health)
}

func health(c echo.Context) error {
	status := &Status{
		Status:    "UP",
		Timestamp: time.Now().Unix(),
	}
	err := config.C.Container.MongoSession.Ping()
	if err != nil {
		status.Status = "DOWN"
		status.Error = err.Error()
	}

	return c.JSON(http.StatusOK, status)
}
