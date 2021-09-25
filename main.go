package main

/*
https://chromium.googlesource.com/external/github.com/src-d/go-git/+/8b0c2116cea2bbcc8d0075e762b887200a1898e1/example_test.go
https://github.com/weaveworks/libgitops
https://pkg.go.dev/github.com/go-git/go-git/v5#example-Clone
https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66
https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66

 */

import (
	"github.com/spf13/viper"
	"gopkg.in/src-d/go-git.v4"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {

	dir, err := createRepoDirectory()
	CheckIfError(err)
	defer os.RemoveAll(dir) // clean up

	url := getEnvVariableFromEnvFile("REPO_URL")

	err = gitClone(err, dir, url)
	changelog, err := os.Open(filepath.Join(dir, "CHANGELOG"))
	CheckIfError(err)

	io.Copy(os.Stdout, changelog)
}


func CheckIfError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

// Clones the repository into the given dir, just as a normal git clone does
func gitClone(err error, dir string, url string) error {
	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL: url,
	})

	CheckIfError(err)
	return err
}

func createRepoDirectory() (string, error) {
	// Tempdir to clone the repository
	dir, err := ioutil.TempDir("", "clone-example")
	CheckIfError(err)
	return dir, err
}



// use viper package to read .env file
// return the value of the key
func getEnvVariableFromEnvFile(key string) string {

	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not CheckIfError any of the config paths.
	// .env - It will search for the .env file in the current directory
	viper.SetConfigFile(".env")

	// Find and read the config file
	err := viper.ReadInConfig()

	CheckIfError(err)
	
	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, ok := viper.Get(key).(string)

	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}