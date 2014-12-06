package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	
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
	
	for i := 1; i < len(branchList.Branches); i++ {
		fmt.Printf("[%d] %s\n", i, branchList.Branches[i])
	}
}
