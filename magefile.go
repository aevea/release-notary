//+build mage

package main

import (
	"github.com/aevea/magefiles"
	"github.com/magefile/mage/sh"
)

func Test() error {
	return magefiles.Test()
}

func SetupTest() error {
	return sh.RunV("sh", "./testdata/setup_test_repos.sh")

}

func GoModTidy() error {
	return magefiles.GoModTidy()
}
