package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var fruits = map[string]string{
	"1": "Pear",
	"2": "Apple",
	"3": "Banana",
}

func main() {
	fruitsCmd := flag.NewFlagSet("fruits", flag.ExitOnError)
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("You must specify a command fruits")
	}

	switch flag.Arg(0) {
	case "fruits":
		ID := fruitsCmd.String("id", "", "find by ID")
		_ = fruitsCmd.Parse(os.Args[2:])

		if *ID != "" {
			fmt.Println(fruits[*ID])
		} else {
			fmt.Println(fruits)
		}
	}
}
