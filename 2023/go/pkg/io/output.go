package io

import (
	"fmt"
	"os"
	"os/exec"
)

func PrintLines(lines []string) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}
