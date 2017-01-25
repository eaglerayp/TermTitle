package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

const prefix = "\\033]0;"
const postfix = "\\007"
const version = "v0.1.0"

func helpMsg() {
	fmt.Printf("TermTitle\nVersion:%s\nChange the terminal title in toolbar.\n\nUsage:\n    TermTitle [title]\n", version)
}

func main() {
	if len(os.Args) < 2 {
		helpMsg()
		return
	}
	title := os.Args[1]
	arg := prefix + title + postfix
	cmd := exec.Command("echo", "-ne", arg)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println(out.String())
}
