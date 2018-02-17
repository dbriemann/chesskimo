package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dbriemann/chesskimo"
)

var version = "undefined"

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Chesskimo", version)

	uci := &chesskimo.UCI{}
	engine := chesskimo.NewEngine("Chesskimo "+version+" 2018", "David Linus Briemann", uci)

	// Input/output runs until exit.
	uci.RunInputOutputLoop(engine)
}
