package main

import "os"
import "fmt"
import "strings"
import "os/exec"

func main() {

	version := "0.1"

	if len(os.Args) < 2 {
		fmt.Printf("Dollarsign v%s\n", version)
		fmt.Println("Brian P. Hogan")
		fmt.Println("Source code: http://github.com/napcs/dollarsign")
		fmt.Println("\nExecutes commands pasted in with leading dollar signs.")
		os.Exit(0)
	}

	program := os.Args[1]
	args := os.Args[2:]

	run(program, args)

	fmt.Println("***** You ran this command with a $ prefixed:")
	fmt.Printf("%s %s\n", program, strings.Join(args[:], ""))
}

func run(program string, args []string) {
	cmd := exec.Command(program, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		fmt.Printf("failed due to :%v\n", err)
	}
}
