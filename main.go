package main

import (
	"fmt"
	"os"

	"github.com/sebastian009w/go-tests-tutorial/firsttest"
)

func main() {
	add := firsttest.Sum(2, 3)

	if add != 5 {
		fmt.Printf("Error: Se esperaba %d, Se obtuvo %d", 5, add)
		os.Exit(1)
	}
	fmt.Println("Succesfuly")
}
