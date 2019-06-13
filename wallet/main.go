package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/iotexproject/iotex-host-wallet/wallet/config"
	"github.com/iotexproject/iotex-host-wallet/wallet/controller"
	"github.com/iotexproject/iotex-host-wallet/wallet/key"
	"github.com/iotexproject/iotex-host-wallet/wallet/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	pwd, _ := reader.ReadString('\n')
	pwd = pwd[:len(pwd)-1]

	seed, err := key.ReadSeed("chain.db", pwd)
	if err != nil {
		log.Fatalf("read master seed error: %v\n", err)
	}
	key.NewMasterKey(seed)

	e := echo.New()
	e.Use(middleware.Logger())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ip := c.RealIP()
			if !utils.Contain(config.C.IPs, ip) {
				return echo.NewHTTPError(http.StatusForbidden, fmt.Sprintf("request ip %s forbidden", ip))
			}
			return next(c)
		}
	})

	controller.AddressRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
