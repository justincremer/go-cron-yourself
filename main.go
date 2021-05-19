package main

import (
	"fmt"
	"log"

	"github.com/justincremer/go-cron-yourself/task"
)

func main() {
	cal, err := task.Create("123456", "123456", "12")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Date: %d\nTime: %d\nWeekday: %d\n", cal.D, cal.T, cal.WD)

	// task.Test("git status --porcelain")
	// task.Test("sudo pacman -Q emacs")
	// task.Test("ls -lah /proc/1/attr")
}
