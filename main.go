package main

/*
https://chromium.googlesource.com/external/github.com/src-d/go-git/+/8b0c2116cea2bbcc8d0075e762b887200a1898e1/example_test.go
https://github.com/weaveworks/libgitops
https://pkg.go.dev/github.com/go-git/go-git/v5#example-Clone
https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66
https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66

 */

import (
	gkconf "gotkube/pkg/config"
	gkgit "gotkube/pkg/git"
	gkdir "gotkube/pkg/system"
	"io"
	"log"
	"os"
	"path/filepath"

)

func main() {

	dir, err := gkdir.CreateRepoDirectory()
	CheckIfError(err)
	defer os.RemoveAll(dir) // clean up

	url := gkconf.GetEnvVariableFromEnvFile("REPO_URL")

	err = gkgit.GitClone(err, dir, url)
	changelog, err := os.Open(filepath.Join(dir, "README.md"))
	CheckIfError(err)

	io.Copy(os.Stdout, changelog)
}


func CheckIfError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}