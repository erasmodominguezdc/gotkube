package config

/*
https://chromium.googlesource.com/external/github.com/src-d/go-git/+/8b0c2116cea2bbcc8d0075e762b887200a1898e1/example_test.go
https://github.com/weaveworks/libgitops
https://pkg.go.dev/github.com/go-git/go-git/v5#example-Clone
https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66
https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66

 */

import (
	viper "github.com/spf13/viper"
	gkdir "gotkube/pkg/system"
	"log"
)

// use viper package to read .env file
// return the value of the key
func GetEnvVariableFromEnvFile(key string) string {

	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not CheckIfError any of the config paths.
	// .env - It will search for the .env file in the current directory
	viper.SetConfigFile(".env")

	// Find and read the config file
	err := viper.ReadInConfig()

	gkdir.CheckIfError(err)
	
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