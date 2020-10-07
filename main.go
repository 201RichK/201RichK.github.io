package main

import (
	"bufio"
	"log"
	"os"
)

var a uintptr = 23
var line []string

func contestResponse() {
	file, err := os.Open("test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	os.Stdin = file

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
}

func main() {
	println("Hello ")
	contestResponse()
}
