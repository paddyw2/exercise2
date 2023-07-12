#!/bin/bash

# create test files outside of go tests to more accurately
# profile memory usage
touch "emptyFile.shrd"
dd if=/dev/urandom bs=1M count=1 | base64 > "stringFile.shrd"
dd if=/dev/urandom bs=1M count=1 | base64 > "stringFileIterations.shrd"
dd if=/dev/urandom bs=1M count=1 | base64 > "stringContents.shrd"
dd if=/dev/urandom bs=1M count=1 > "binaryFile.shrd"
dd if=/dev/urandom bs=1M count=100 > "binaryFileLarge.shrd"
# unwritable file
dd if=/dev/urandom bs=1M count=1 | base64 > "unwritable.shrd"
chmod -w "unwritable.shrd"
