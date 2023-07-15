package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// do Quake log parsing

	elapsed := time.Since(start)
	fmt.Printf("Done in %s\n", elapsed)
}
