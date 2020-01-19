//program to print flaged and unflaged arguments
package main

import (
	"flag"
	"fmt"
)

func main() {

	fmt.Println("Program Started")
	flag1 := flag.String("name", "self", "name of the person")
	flag.Parse()
	fmt.Println(*flag1)

	//to print non flaged arguments
	fmt.Println("Non flaged arguments", flag.Args())
}
