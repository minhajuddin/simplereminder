package main

import (
	"fmt"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("starting")
	httprouter.New()
}
