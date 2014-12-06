package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"

	"gopkg.in/yaml.v2"
)

type BranchList struct {
	Branches []string
}

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func OpenRCFile() (f *os.File, err error) {
	file_path := getHomeDir() + "/.git-checkout-historyrc"
	if _, err := os.Stat(file_path); os.IsNotExist(err) {
		return os.Create(file_path)
	} else {
		return os.Open(file_path)
	}
}

func store(branch string) {
	branchList := BranchList{}
	rcfile, err := OpenRCFile()
	if err != nil {
		log.Fatal(err)
	}
	rcfile.Close()
	
	file_path := getHomeDir() + "/.git-checkout-historyrc"
	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Fatal(err)
	}
	
	err = yaml.Unmarshal(data, &branchList)
	if err != nil {
		log.Fatal(err)
	}
	
	branchList.Branches = append(branchList.Branches, branch)
	
	data, err = yaml.Marshal(&branchList)
	if err != nil {
		log.Fatal(err)
	}
	
	err = ioutil.WriteFile(file_path, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	args := os.Args[1:]
	
	if len(args) > 0 {
		store(args[0])
		
		cmd := exec.Command("git", "checkout", args[0])
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
		fmt.Println(out.String())
	}
}
