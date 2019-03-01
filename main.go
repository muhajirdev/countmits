package main

import (
	"log"
	"os"
	"os/exec"
)

func getBranchName() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	// If branch name is not specified, get commits for all branch
	return "--all"

}

func main() {
	branch := getBranchName()

	cmd := exec.Command("git", "rev-list", "--count", branch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("%v\nPlease install git \n", err.Error())
	}
}
