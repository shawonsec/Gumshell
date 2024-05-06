package main

import (
	"fmt"
	"os/exec"
	"os/user"
	"strings"
)

func main() {
	// Display a welcome message.
	fmt.Println("Welcome to GumShell!")

	for {
		// Get the current user.
		currentUser, err := user.Current()
		if err != nil {
			panic(err)
		}

		// Use $ for normal users and # for root.
		prompt := "$"
		if currentUser.Uid == "0" {
			prompt = "#"
		}

		fmt.Printf("%s ", prompt)
		var input string
		fmt.Scanln(&input)
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		cmd := exec.Command("/bin/sh", "-c", input)
		output, err := cmd.Output()

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(string(output))
		}
	}
}
