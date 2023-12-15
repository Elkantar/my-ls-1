/*
	sys exec : https://gobyexample.com/execing-processes
*/
package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	args := os.Args[1:]

	// chemin absolu (/bin/ls)
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Environement actuel
	env := os.Environ()

	args = append([]string{"ls"}, args...)
	syscall.Exec(binary, args, env)
}
