package main

import (
	"fmt"
	"github.com/levpa/go-base/internal/auth"
  "github.com/levpa/go-base/algo"
)

func main() {
	fmt.Printf("Pull first: %s\n", algo.RandStr(20))
	fmt.Printf("Pull second: %s\n", auth.RandStr(20))
}
