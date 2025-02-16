# Add this new target to your existing Makefile
build_linux:
	GOOS=linux GOARCH=amd64 go build -o motd .

# Update .PHONY to include the new target
.PHONY: build_linux
