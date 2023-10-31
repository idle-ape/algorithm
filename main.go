package main

import (
	"fmt"
	"os"

	"github.com/idle-ape/algorithm/exec"
	_ "github.com/idle-ape/algorithm/linked-list"
)

func main() {
	argc := len(os.Args)
	if argc != 2 {
		fmt.Printf("Usage: %s <num>\n", os.Args[0])
		return
	}

	exec.Exec(os.Args[1])
}
