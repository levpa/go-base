package main

import (
	"fmt"
	"github.com/levpa/go-base/internal/auth"
  "github.com/levpa/go-base/algo"
)

func main() {
	fmt.Printf("Random string: %s\n", algo.RandStr(20))
	fmt.Printf("Random string: %s\n", auth.RandStr(20))
}
