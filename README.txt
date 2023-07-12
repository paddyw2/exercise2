Shred

Usage:

`Shred("yourfilename")`

Testing:

Run `make run` or `bash create_test_files.sh` followed by `go test`

Note: test files are created outside the go tests to allow the memory consumption of the go tests to be profile easier without being influenced by memory used during the creation of test cases

Configuration:
 - ENABLE_LOGGING can be changed within shred.go

Test cases handled:
 - handles cases where the file does not exist
 - throws an error when the file cannot be written to
 - handles both text and binary formats
 - handles files of varying sizes
 - deletes the file after shredding
 - ensures file contents before deletion is different to the original contents
 - records memory usage to check that it is capped (i.e. shredding 100MB file does not use 100MB of memory)
 - checking for open file handles after completion
 - empty files

Test cases not handled:
 - behaviour with "non-standard" files, such as virtual filesystems, network drives, devices, etc.
 - low memory devices (although basic memory usage tracked)
 - simulating devices low on storage and handling "No space left on device" errors
