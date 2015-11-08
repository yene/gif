package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var filename = "gifs.txt"

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 2 {
		printHelp()
		return
	}

	arg := os.Args[1]
	switch arg {
	case "get":
		get()
		return
	case "add":
		add()
		return
	}
	printHelp()
	return
}

func get() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 1 {
		printHelp()
		return
	}
}

func add() {
	out, _ := exec.Command("/usr/bin/pbpaste").Output()
	paste := string(out)
	paste = paste + ", " + strings.Join(os.Args[2:], " ") + "\n"
	appendToStorage(paste)
}

func printHelp() {
	fmt.Println("Commands:")
	fmt.Println("gif add description")
	fmt.Println("gif get searchtext")
}

func appendToStorage(v string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(v); err != nil {
		panic(err)
	}
}
