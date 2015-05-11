package main

import (
	"fmt"
	"os"

	"github.com/arunjit/brainfuck-go"
)

func main() {
	fmt.Println(brainfuck.EvalToString(os.Args[1]))
}
