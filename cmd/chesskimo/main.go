package main

import (
	"fmt"

	"github.com/dbriemann/chesskimo"
)

var version = "undefined"

func main() {
	fmt.Println("Chesskimo", version)

	uci := &chesskimo.UCI{}
	engine := chesskimo.NewEngine("Chesskimo "+version+" 2018", "David Linus Briemann", uci)

	// Input/output runs until exit.
	uci.RunInputOutputLoop(engine)
}
