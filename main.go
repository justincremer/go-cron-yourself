package main

import (
	"github.com/justincremer/go-cron-yourself/task"
)

func main() {
	task.Test("git status --porcelain")
	task.Test("sudo pacman -Q emacs")
	task.Test("ls -lah /proc/1/attr")
}
