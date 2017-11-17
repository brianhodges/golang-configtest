package main

import (
	"fmt"
	"os"
)

type Config struct {
	Environment string
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("You must specify an environment parameter")
		os.Exit(0)
	}
	
	config := Config{Environment: os.Args[1]}
	
	//command prompt params
	switch config.Environment {
		case "-debug": fmt.Println("Running in Debug...")
		case "-dev"  : fmt.Println("Running in Development...")
		case "-prod" : fmt.Println("Running in Production...")
		default:
			fmt.Println("Add -debug, -dev, or -prod as a parameter")
			os.Exit(0)
	}
}