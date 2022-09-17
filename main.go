package main

import (
	"fmt"

	"github.com/google/uuid"

	"rsc.io/quote"
)

func main() {
	fmt.Println(uuid.New().String())
	fmt.Println(quote.Go())
}
