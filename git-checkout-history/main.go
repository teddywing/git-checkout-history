package main

import (
	"fmt"
	"github.com/teddywing/git-checkout-history/utils"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	
	"gopkg.in/yaml.v2"
)

// TODO: move to package
type BranchList struct {
	Branches []string
}

// TODO: move to package
func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func main() {
	branchList := BranchList{}
	
	file_path := getHomeDir() + "/.git-checkout-history"
	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Fatal(err)
	}
	
	err = yaml.Unmarshal(data, &branchList)
	if err != nil {
		log.Fatal(err)
	}
	
	args := os.Args[1:]
	
	if len(args) > 0 {
		branchIndex, _ := strconv.Atoi(args[0])
		
		cmd := exec.Command("git", "checkout", branchList.Branches[branchIndex])
		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
		
		utils.Store(branchList.Branches[branchIndex])
	} else {
		// List branches in history
		for i := 1; i < len(branchList.Branches); i++ {
			fmt.Printf("[%d] %s\n", i, branchList.Branches[i])
		}
	}
}
