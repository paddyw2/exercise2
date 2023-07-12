package shred

import (
	"fmt"
	"log"
	"os"
	"testing"
)

const CHUNKSIZE = 1000

func getTestLogger() *log.Logger {
	logger := log.Default()
	logger.SetPrefix("[TEST] ")
	return logger
}

func CheckFileExists(path string) {
	if !FileExists(path) {
		panic(fmt.Sprintf("Expected case file %s does not exist - run test file creation script", path))
	}
}

func TestShred(t *testing.T) {
	testCases := []string{"stringFile.shrd", "binaryFile.shrd"}

	for _, testFileName := range testCases {
		CheckFileExists(testFileName)
		Shred(testFileName)
		if FileExists(testFileName) {
			t.Errorf("File %s was not deleted", testFileName)
		}
	}
}

func TestShredLargeFile(t *testing.T) {
	binaryFileLarge := "binaryFileLarge.shrd"
	CheckFileExists(binaryFileLarge)
	Shred(binaryFileLarge)
}

func TestShredEmptyFile(t *testing.T) {
	emptyFile := "emptyFile.shrd"
	CheckFileExists(emptyFile)
	Shred(emptyFile)
}

func TestShredNoFile(t *testing.T) {
	Shred("doesnotexist.shrd")
}

func TestShredUnwritable(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected function to panic due to permission denied")
		}
	}()
	unwritable := "unwritable.shrd"
	CheckFileExists(unwritable)
	Shred(unwritable)
}

func TestShredFileContents(t *testing.T) {
	stringFile := "stringContents.shrd"
	CheckFileExists(stringFile)
	contents, err := os.ReadFile(stringFile)
	check(err)
	ShredFile(stringFile, 3, false, CHUNKSIZE)
	newContents, err := os.ReadFile(stringFile)
	check(err)
	if string(contents) == string(newContents) {
		t.Errorf("New file contents should not equal original contents")
	}
	deleteErr := os.Remove(stringFile)
	check(deleteErr)
}

func TestShredExtraIterations(t *testing.T) {
	stringFile := "stringFileIterations.shrd"
	CheckFileExists(stringFile)
	shredCount := 10
	ShredFile(stringFile, shredCount, true, CHUNKSIZE)
	if FileExists(stringFile) {
		t.Errorf("File %s was not deleted", stringFile)
	}
}
