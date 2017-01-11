package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

const prefix = "\\033]0;"
const postfix = "\\007"

func main() {
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
