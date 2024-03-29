package main

import (
	"fmt"
	"os"

	_ "github.com/idle-ape/algorithm/array"
	_ "github.com/idle-ape/algorithm/binary-tree"
	_ "github.com/idle-ape/algorithm/dynamic_programming"
	"github.com/idle-ape/algorithm/exec"
	_ "github.com/idle-ape/algorithm/linked-list"
	_ "github.com/idle-ape/algorithm/string"
)

func main() {
	argc := len(os.Args)
	if argc != 2 {
		fmt.Printf("Usage: %s <num>\n", os.Args[0])
		return
	}

	exec.Exec(os.Args[1])
}
