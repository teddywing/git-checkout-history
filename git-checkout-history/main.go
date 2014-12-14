package main

import (
	"fmt"
	"github.com/teddywing/git-checkout-history/utils"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	branches := utils.Branches()
	
	args := os.Args[1:]
	
	if len(args) > 0 {
		branchIndex, _ := strconv.Atoi(args[0])
		
		cmd := exec.Command("git", "checkout", branches[branchIndex])
		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
		
		utils.Store(branches[branchIndex])
	} else {
		// List branches in history
		for i := 1; i < len(branches); i++ {
			fmt.Printf("[%d] %s\n", i, branches[i])
		}
	}
}
