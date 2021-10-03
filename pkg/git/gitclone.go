package git

/*
https://chromium.googlesource.com/external/github.com/src-d/go-git/+/8b0c2116cea2bbcc8d0075e762b887200a1898e1/example_test.go
https://github.com/weaveworks/libgitops
https://pkg.go.dev/github.com/go-git/go-git/v5#example-Clone
https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66
https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66

 */

import (
	"gopkg.in/src-d/go-git.v4"
	gkdir "gotkube/pkg/system"
)

// Clones the repository into the given dir, just as a normal git clone does
func GitClone(err error, dir string, url string) error {
	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL: url,
	})

	gkdir.CheckIfError(err)
	return err
}