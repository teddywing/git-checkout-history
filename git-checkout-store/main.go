package main

import (
	"bytes"
	"fmt"
	"github.com/teddywing/git-checkout-history/utils"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]
	
	if len(args) > 0 {
		utils.Store(args[0])
		
		cmd := exec.Command("git", "checkout", args[0])
		var out bytes.Buffer
		cmd.Stderr = &out
		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
		fmt.Print(out.String())
	}
}
