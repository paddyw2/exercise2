package shred

import (
	"fmt"
	"io"
	"log"
	"os"
)

const LOG_PREFIX = "[SHRED] "
const ENABLE_LOGGING = true

func GetLogger() *log.Logger {
	logger := log.Default()
	logger.SetPrefix(LOG_PREFIX)
	if !ENABLE_LOGGING {
		logger.SetOutput(io.Discard)
	}
	return logger
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func FileExists(path string) bool {
	_, statErr := os.Stat(path)
	if os.IsNotExist(statErr) {
		return false
	} else {
		return true
	}
}

// ShredFile
// Overwrites a file `shredCount` times with random data
// Deletes the file if `removeFile` is set to true
// Writes file in chunks of size `chunkSize` bytes
func ShredFile(path string, shredCount int, removeFile bool, chunkSize int) {
	logger := GetLogger()

	if !FileExists(path) {
		return
	}

	logger.Print("Running shred on ", path)
	for shredIter := 0; shredIter < shredCount; shredIter++ {
		logger.Print("---> Shred iteration: ", shredIter)
		// open fd for read/write
		fileToShred, err := os.OpenFile(path, os.O_RDWR, 0)
		check(err)
		defer fileToShred.Close()

		// get current fd information
		fileInfo, err := fileToShred.Stat()
		check(err)

		// store current fd size
		shredFileSize := fileInfo.Size()

		// write file in chunks to avoid bringing entire
		// contents into memory
		chunks := shredFileSize / int64(chunkSize)
		remainder := int(shredFileSize % int64(chunkSize))

		// open /dev/random to sample random data
		randomFile, err := os.Open("/dev/urandom")
		check(err)
		defer randomFile.Close()

		for chunkIter := int64(0); chunkIter < chunks; chunkIter++ {
			// read chunkSize random data from /dev/random
			randomBytes := make([]byte, chunkSize)
			_, err := io.ReadAtLeast(randomFile, randomBytes, chunkSize)
			check(err)

			// write to file
			fileToShred.Write(randomBytes)
		}

		// write remainder of bytes required from /dev/random
		randomBytes := make([]byte, remainder)
		_, err = io.ReadAtLeast(randomFile, randomBytes, remainder)
		check(err)

		// write to file
		fileToShred.Write(randomBytes)
	}

	if removeFile {
		logger.Print("Deleting file: ", path)
		deleteErr := os.Remove(path)
		check(deleteErr)
	}
}

// Shred
// Overwrites a file 3x with random data and deletes it after
func Shred(path string) {
	// defaults
	chunkSize := 1000
	shredCount := 3
	removeFile := true
	ShredFile(path, shredCount, removeFile, chunkSize)
}
