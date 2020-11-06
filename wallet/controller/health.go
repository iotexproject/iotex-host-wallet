package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/iotexproject/iotex-host-wallet/wallet/config"
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
	err := config.C.Container.MongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		status.Status = "DOWN"
		status.Error = err.Error()
	}

	return c.JSON(http.StatusOK, status)
}
