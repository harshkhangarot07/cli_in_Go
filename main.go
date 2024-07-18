package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	name := flag.String("name", "world", "a name to say hello")
	age := flag.Int("age", 0, "your age")
	help := flag.Bool("help", false, "display help")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	args := flag.Args()

	if len(args) > 0 {
		fmt.Println("positional arguments ", args)
	}

	fmt.Printf("Hello, %s! You are %d years old.\n", *name, *age)

}
