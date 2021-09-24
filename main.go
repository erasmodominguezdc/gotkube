package gotkube

/*
https://chromium.googlesource.com/external/github.com/src-d/go-git/+/8b0c2116cea2bbcc8d0075e762b887200a1898e1/example_test.go
https://github.com/weaveworks/libgitops
https://pkg.go.dev/github.com/go-git/go-git/v5#example-Clone
 */
import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)


func main() {
	fmt.Println("hello world")
	// Tempdir to clone the repository
	dir, err := ioutil.TempDir("", "clone-example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir) // clean up

	// Clones the repository into the given dir, just as a normal git clone does
	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL: "https://github.com/git-fixtures/basic.git",
	})

	if err != nil {
		log.Fatal(err)
	}

	// Prints the content of the CHANGELOG file from the cloned repository
	changelog, err := os.Open(filepath.Join(dir, "CHANGELOG"))
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, changelog)
}
