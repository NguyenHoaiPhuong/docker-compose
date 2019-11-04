package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Hello world")
		return
	}

	name := flag.String("name", "Anymous", "name of the user")
	age := flag.Int64("age", 100, "age of the user")
	flag.Parse()

	fmt.Println("Hello", *name, ". You are", *age, "years old.")
}
