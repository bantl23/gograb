package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"golang.org/x/mod/modfile"
)

const (
	gomod = "go.mod"
)

var (
	version = "v0.0.0"
	commit  = "unknown"
)

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "-v" || os.Args[1] == "--version" {
			fmt.Printf("%s (%s)\n", version, commit)
			os.Exit(0)
		}
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Println("Downloads go modules one at a time")
			os.Exit(0)
		}
	}

	_, err := os.Stat(gomod)
	if os.IsNotExist(err) {
		fmt.Println(err)
		os.Exit(1)
	}

	contents, err := ioutil.ReadFile(gomod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := modfile.Parse(gomod, contents, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for j := range file.Require {
		modvers := fmt.Sprintf("%+v@%+v", file.Require[j].Mod.Path, file.Require[j].Mod.Version)
		cmdGoGet := exec.Command("go", "get", "-d", modvers)
		fmt.Printf("downloading %s\n", modvers)
		resultGoGet, err := cmdGoGet.CombinedOutput()
		if err != nil {
			fmt.Printf("%s [%+v]\n", resultGoGet, err)
			os.Exit(1)
		}
	}
}
