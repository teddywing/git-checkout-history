package utils

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"strconv"

	"gopkg.in/yaml.v2"
)

var history_file string = ".git-checkout-history"

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func currentGitDir() string {
	// Check git version
	// If below 1.7.0 then exit 1. `--show-toplevel` is not available in prior
	// versions.
	// /git version (\d+\.\d+).*/
	// Run `git rev-parse --show-toplevel`
	// Return output from the command
	cmd := exec.Command("git", "--version")
	var version_string bytes.Buffer
	cmd.Stdout = &version_string
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	
	r, _ := regexp.Compile(`git version (\d+\.\d+).*`)
	matches := r.FindStringSubmatch(version_string.String())
	
	version_number, _ := strconv.ParseFloat(matches[1], 64)
	if version_number >= 1.7 {
		cmd = exec.Command("git", "rev-parse", "--show-toplevel")
		var git_directory bytes.Buffer
		cmd.Stdout = &git_directory
		if err := cmd.Run(); err != nil {
			panic(err)
		}
		
		dir_trimmed := git_directory.String()
		
		// Trim newline at the end of the directory string
		return dir_trimmed[:len(dir_trimmed) - 1]
	}
	
	return ""
}

func OpenHistoryFile() (f *os.File, err error) {
	file_path := getHomeDir() + "/" + history_file
	if _, err := os.Stat(file_path); os.IsNotExist(err) {
		return os.Create(file_path)
	} else {
		return os.Open(file_path)
	}
}

func Store(branch string) {
	rcfile, err := OpenHistoryFile()
	if err != nil {
		log.Fatal(err)
	}
	rcfile.Close()
	
	file_path := getHomeDir() + "/" + history_file
	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Fatal(err)
	}
	
	branchList := make(map[string][]string)
	
	err = yaml.Unmarshal(data, &branchList)
	if err != nil {
		log.Fatal(err)
	}
	
	current_git_dir := currentGitDir()
	branchList[current_git_dir] = append([]string{branch}, branchList[current_git_dir]...)
	
	data, err = yaml.Marshal(&branchList)
	if err != nil {
		log.Fatal(err)
	}
	
	err = ioutil.WriteFile(file_path, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func Branches() []string {
	branchList := make(map[string][]string)
	
	file_path := getHomeDir() + "/" + history_file
	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		log.Fatal(err)
	}
	
	err = yaml.Unmarshal(data, &branchList)
	if err != nil {
		log.Fatal(err)
	}
	
	return branchList[currentGitDir()]
}
