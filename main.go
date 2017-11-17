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
	if config.Environment == "-debug" {
		fmt.Println("Running in Debug...")
	} else if config.Environment == "-dev" {
		fmt.Println("Running in Development...")
	} else if config.Environment == "-prod" {
		fmt.Println("Running in Production...")
	} else {
		fmt.Println("Add -debug, -dev, or -prod as a paramter")
		os.Exit(0)
	}
}