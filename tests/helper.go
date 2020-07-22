package test

import (
	"go/build"
	"math/rand"
	"os"
	"strings"
	"time"
)

const changeIt string = "src/change_it/s3-log-uploader/"

// Has some useful functions used in the tests.
// @author rnojiri

// GetScanPathRoot - build a scan path
func GetScanPathRoot(path string) string {

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	return gopath + changeIt + path
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandomString - generates random strings
func RandomString(n int) string {
	rand.Seed(int64(time.Now().Nanosecond()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return strings.ToLower(string(b))
}
