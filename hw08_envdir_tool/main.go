package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args

	env, err := ReadDir(args[1])
	if err != nil {
		log.Fatal(err)
	}

	exitCodeCmd := RunCmd(args[2:], env)
	os.Exit(exitCodeCmd)
}
