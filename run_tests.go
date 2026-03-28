package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Start test cases...")

	cmd := exec.Command("go", "test", "-v", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Tests are wrong:", err)
		os.Exit(1)
	} else {
		fmt.Println("Tests succeeded")
	}
}
