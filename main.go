package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
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
	readStorage()
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

func readStorage() {
	usr, _ := user.Current()
	dir := usr.HomeDir
	f, err := os.OpenFile(dir+"/"+filename, os.O_RDONLY, 0600)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func appendToStorage(v string) {
	usr, _ := user.Current()
	dir := usr.HomeDir
	f, err := os.OpenFile(dir+"/"+filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(v); err != nil {
		panic(err)
	}
}
