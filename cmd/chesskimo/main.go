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
	engine := chesskimo.NewEngine("Chesskimo "+version+" 2022", "David Linus Briemann", uci, chesskimo.SimpleMCSearch)
	// engine := chesskimo.NewEngine("Chesskimo "+version+" 2022", "David Linus Briemann", uci, chesskimo.AlphaBetaSearch)

	// Input/output runs until exit.
	engine.Run()
}
