package main

import (
	"fmt"
	"log"
	"os"

	"github.com/syniol/lucysplains/internal"
)

func main() {
	ai, err := internal.NewLucySplains()
	if err != nil {
		log.Fatal(err.Error())
	}

	resp, err := ai.Splain(os.Args[0])
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(resp)
}
