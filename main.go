package main

import (
	"fmt"
	"os"
	"os/exec"
)

func run(source string) {
	cmd := exec.Command("go", "run", ".")
	cmd.Dir = source
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		out, _ := cmd.CombinedOutput()
		fmt.Println(string(out))
		return
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	command := args[0]
	args = args[1:]

	switch command {
	case "run":
		if len(args) == 0 {
			fmt.Println("provide source")
			return
		}

		source := args[0]
		run(source)
		return
	default:
		fmt.Println("unknown command")
		return
	}
}
