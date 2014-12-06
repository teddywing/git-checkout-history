package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
)

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func OpenRCFile() {
	filename := ".git-checkout-historyrc"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		_, err := os.Create(getHomeDir() + "/" + filename)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	args := os.Args[1:]

	OpenRCFile()
	
	if len(args) > 0 {
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
