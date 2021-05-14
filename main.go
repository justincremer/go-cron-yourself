package main

import (
	"log"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := CreateCmd("git status --porcelain")
	fmt.Printf("Command: %s\n", cmd.String())

	res, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("result: %s\n", res)
}

func CreateCmd(i string) *exec.Cmd {
	switch argv := strings.Fields(i); len(argv) {
	case 0:
		log.Fatal("Command cannot be nil")
		return nil
	case 1:
		return exec.Command(argv[0])
	default:
		return exec.Command(argv[0], argv[1:]...)
	}
}
