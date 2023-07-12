test:
	@echo "Creating test files via bash script..."
	bash create_test_files.sh
	@echo "Running tests and profiling using /bin/time..."
	/bin/time -v go test
	@echo "Checking for open .shrd files..."
	@lsof | grep *.shrd || echo "No open *.shrd test files found - success"
	@echo "Force removing unwritable test case files..."
	rm -f unwritable.shrd
