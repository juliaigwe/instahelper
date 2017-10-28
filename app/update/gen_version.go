// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Masterminds/semver"
)

var tmpl = `package update

// VERSION of instahelper
var VERSION = "%s"`

func main() {
	path, _ := filepath.Abs(".")
	path = filepath.Join(path, "app", "update", "version.go")

	out := fmt.Sprintf(tmpl, os.Args[1])

	// Check if version is proper semver
	if os.Args[1] != "nightly" {
		semver.MustParse(os.Args[1])
	}

	err := ioutil.WriteFile(path, []byte(out), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
