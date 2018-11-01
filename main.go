package main

import (
	"fmt"
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
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%v\nPlease install git \n", err.Error())
	}

	fmt.Println(string(output))
}
