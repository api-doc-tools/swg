.PHONY: test
test:
	go test -coverprofile=covprofile
	go tool cover -html=covprofile -o coverage.html