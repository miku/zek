package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/miku/zek"
)

func main() {
	flag.Parse()

	root := new(zek.Node)
	if _, err := root.ReadFrom(os.Stdin); err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(root)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
