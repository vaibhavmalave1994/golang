//program to print flaged and unflaged arguments
package main

import (
	"flag"
	"fmt"
	"os"
)

type Person struct {
	Name   string
	Age    int
	IsGood bool
}

var person Person

func init() {
	flagSet := flag.NewFlagSet("person", flag.ContinueOnError)
	var name string
	var age int
	var isGood bool

	flagSet.StringVar(&name, "name", "", "name of the person")
	flagSet.IntVar(&age, "age", 1, "age of the person")
	flagSet.BoolVar(&isGood, "isgood", false, "is person good")

	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("error while parsing flags.")
		os.Exit(1)
	}
	person = Person{name, age, isGood}
}

func main() {
	fmt.Println("person", person)
	fmt.Println("ok")
}
