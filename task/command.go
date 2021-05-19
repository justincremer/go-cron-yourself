package task

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Task struct {
	Command  string
	Calendar Calendar
	Error    error
}

func Test(i string) {
	cmd, err := CreateCmd(i)
	must(err)
	fmt.Printf("Command: %s\n", cmd.String())

	res, err := cmd.Output()
	must(err)
	fmt.Printf("result: %s\n", res)
}

func CreateCmd(i string) (*exec.Cmd, error) {
	switch argv := strings.Fields(i); len(argv) {
	case 0:
		return nil, errors.New("Error: input must not be nil")
	case 1:
		return exec.Command(argv[0]), nil
	default:
		return exec.Command(argv[0], argv[1:]...), nil
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
