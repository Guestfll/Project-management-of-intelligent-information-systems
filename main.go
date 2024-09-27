package main

import (
	"fmt"
	"os"
	"project-management-practice-1/packages"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: main [|uppertask|beststudents]")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "uppertask":
		packages.InitToUpperTask()
	case "beststudents":
		packages.GetBestStudents()
	default:
		fmt.Println("Unknown command. Usage: main [|uppertask|beststudents]")
		os.Exit(1)
	}
}
