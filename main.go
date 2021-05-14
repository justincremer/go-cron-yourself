package main

import (
	"log"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "status", "--porcelain")
	fmt.Printf("Command: %s\n", cmd.String())
	
	res, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("result: %s\n", res)
}
