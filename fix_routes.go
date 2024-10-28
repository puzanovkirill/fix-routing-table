package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Starting route modification...")

	if runCommand("sudo", "route", "-n", "delete", "192.168.122.0/24", "192.168.122.1") {
		fmt.Println("Route to 192.168.122.0/24 successfully deleted.")
	} else {
		fmt.Println("Failed to delete route to 192.168.122.0/24. It may not exist.")
	}

	if runCommand("sudo", "route", "-n", "add", "192.168.122.0/24", "192.168.122.1") {
		fmt.Println("Route to 192.168.122.0/24 successfully added.")
	} else {
		fmt.Println("Failed to add route to 192.168.122.0/24.")
	}

	if runCommand("sudo", "route", "-n", "delete", "default", "192.168.122.1") {
		fmt.Println("Default route successfully deleted.")
	} else {
		fmt.Println("Failed to delete default route. It may not exist.")
	}

	if runCommand("sudo", "route", "-n", "add", "default", "192.168.122.1") {
		fmt.Println("Default route successfully added.")
	} else {
		fmt.Println("Failed to add default route.")
	}

	fmt.Println("Route modification process completed.")
}

func runCommand(command string, args ...string) bool {
	cmd := exec.Command(command, args...)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing command '%s %v': %v\n", command, args, err)
		return false
	}
	return true
}
