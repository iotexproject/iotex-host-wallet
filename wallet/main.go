package main

import (
	"bufio"
	"log"
	"os"

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
}
