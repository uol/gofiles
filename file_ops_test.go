package gofiles_test

import (
	"go/build"
	"math/rand"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/uol/gofiles"
)

//
// Test all functions from file operations package
// @author rnojiri
//

// getScanPathRoot - build a scan path
func getScanPathRoot(path string) string {

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	return gopath + "/src/github.com/uol/gofiles/tests/" + path
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

func TestGzipFile(t *testing.T) {

	filePath := getScanPathRoot("gzip/large-text.log")

	fileInfo, err := os.Stat(filePath)
	err = gofiles.GzipFile(filePath, filePath+".gz")
	assert.NoError(t, err)

	fileInfoGziped, err := os.Stat(filePath + ".gz")
	assert.NoError(t, err)
	assert.True(t, fileInfo.Size() > fileInfoGziped.Size())

	gofiles.DeleteFile(filePath + ".gz")
}

func TestGzipDecompressFile(t *testing.T) {

	filePath := getScanPathRoot("gzip/gziped-large-text.log")

	fileInfoGziped, err := os.Stat(filePath + ".gz")
	err = gofiles.GzipDecompressFile(filePath+".gz", filePath)
	assert.NoError(t, err)

	fileInfo, err := os.Stat(filePath)
	assert.NoError(t, err)
	assert.True(t, fileInfo.Size() > fileInfoGziped.Size())

	gofiles.DeleteFile(filePath)
}
