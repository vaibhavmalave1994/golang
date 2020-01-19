//This program prints its command line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	//started with index 1 as index 0 here is command itself
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("Value : %s \n", os.Args[i])
	}
}
