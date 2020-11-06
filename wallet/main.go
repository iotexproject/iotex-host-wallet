package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/iotexproject/iotex-host-wallet/wallet/config"
	"github.com/iotexproject/iotex-host-wallet/wallet/controller"
	"github.com/iotexproject/iotex-host-wallet/wallet/key"
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

	controller.AddressRoute(e)
	controller.SignerRoute(e)
	controller.HealthRoute(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.C.Port)))
}
