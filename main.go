package gotkube

/*
https://chromium.googlesource.com/external/github.com/src-d/go-git/+/8b0c2116cea2bbcc8d0075e762b887200a1898e1/example_test.go
https://github.com/weaveworks/libgitops

 */
import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"gopkg.in/src-d/go-billy.v4/memfs"
)


func main() {
	fmt.Println("hello world")
	ExampleClone()
}

func ExampleClone() {
	// Filesystem abstraction based on memory
	fs := memfs.New()
	// Git objects storer based on memory
	storer := memory.NewStorage()
	// Clones the repository into the worktree (fs) and storer all the .git
	// content into the storer
	_, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: "https://github.com/git-fixtures/basic.git",
	})
	if err != nil {
		log.Fatal(err)
	}
	// Prints the content of the CHANGELOG file from the cloned repository
	changelog, err := fs.Open("CHANGELOG")
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, changelog)
	// Output: Initial changelog
}

