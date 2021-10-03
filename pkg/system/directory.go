package system

/*
https://chromium.googlesource.com/external/github.com/src-d/go-git/+/8b0c2116cea2bbcc8d0075e762b887200a1898e1/example_test.go
https://github.com/weaveworks/libgitops
https://pkg.go.dev/github.com/go-git/go-git/v5#example-Clone
https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66
https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66

 */

import (
	"io/ioutil"
	"log"
)


func CheckIfError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}


func CreateRepoDirectory() (string, error) {
	// Tempdir to clone the repository
	dir, err := ioutil.TempDir("", "clone-example")
	CheckIfError(err)
	return dir, err
}