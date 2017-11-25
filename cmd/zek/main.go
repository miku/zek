package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/miku/zek"
)

var createExampleProgram = flag.Bool("p", false, "write out an example program")

func main() {
	flag.Parse()

	root := new(zek.Node)
	if _, err := root.ReadFrom(os.Stdin); err != nil {
		log.Fatal(err)
	}

	if *createExampleProgram {
		sw := zek.NewStructWriter(os.Stdout)
		if err := sw.WriteNode(root); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	b, err := json.Marshal(root)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
