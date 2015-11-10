package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

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
	case "export":
		export()
		return
	}
	printHelp()
	return
}

func get() {
	t := strings.Join(os.Args[2:], " ")
	r := searchStorage(t)

	cmd := exec.Command("/usr/bin/pbcopy")
	cmd.Stdin = strings.NewReader(r)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func add() {
	out, _ := exec.Command("/usr/bin/pbpaste").Output()
	paste := string(out)
	paste = paste + " " + strings.Join(os.Args[2:], " ") + "\n"
	appendToStorage(paste)
}

func export() {
	f, err := os.OpenFile(storagePath(), os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	md := ""
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		p := strings.Split(scanner.Text(), " ")
		text := strings.Join(p[1:], " ")
		md += "[" + text + "](" + p[0] + ")\n"
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(os.Args[2], []byte(md), 0644)
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("gif add description")
	fmt.Println("gif get searchtext")
}

func searchStorage(searchtext string) string {
	f, err := os.OpenFile(storagePath(), os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		p := strings.Split(scanner.Text(), " ")
		text := strings.Join(p[1:], " ")
		if strings.Contains(text, searchtext) {
			return p[0]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ""
}

func appendToStorage(v string) {
	f, err := os.OpenFile(storagePath(), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(v); err != nil {
		panic(err)
	}
}

func storagePath() string {
	var filename = "gifs.txt"
	usr, _ := user.Current()
	dir := usr.HomeDir
	return dir + "/" + filename
}
