package utils

import (
	"io/ioutil"
	"log"
	"os"
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
	file_path := getHomeDir() + "/.git-checkout-history"
	if _, err := os.Stat(file_path); os.IsNotExist(err) {
		return os.Create(file_path)
	} else {
		return os.Open(file_path)
	}
}

func Store(branch string) {
	branchList := BranchList{}
	rcfile, err := OpenRCFile()
	if err != nil {
		log.Fatal(err)
	}
	rcfile.Close()
	
	file_path := getHomeDir() + "/.git-checkout-history"
	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Fatal(err)
	}
	
	err = yaml.Unmarshal(data, &branchList)
	if err != nil {
		log.Fatal(err)
	}
	
	branchList.Branches = append([]string{branch}, branchList.Branches...)
	
	data, err = yaml.Marshal(&branchList)
	if err != nil {
		log.Fatal(err)
	}
	
	err = ioutil.WriteFile(file_path, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
