package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ququzone/go-common/crypto"
)

func main() {
	out := flag.String("o", "chain.db", "out file name")
	password := flag.String("p", "123456", "file password")
	flag.Parse()

	aes := crypto.NewAES(*password, *out)

	of, err := os.Create(*out)
	if err != nil {
		log.Fatalf("create out file error:%v\n", err)
	}
	or := bufio.NewWriter(of)
	defer func() {
		or.Flush()
		of.Close()
	}()

	seed := make([]byte, 256)
	_, err = rand.Read(seed)
	if err != nil {
		log.Fatalf("generate seed error:%v\n", err)
	}

	k, err := aes.Encrypt(seed)
	if err != nil {
		log.Fatalf("encrypt seed error:%v\n", err)
	}
	or.Write(k)

	fmt.Println("encrypt seed successful.")
}
