all: build_windows_amd64 build_macos_amd64 build_macos_arm64 build_linux_amd64 build_linux_i386

build_windows_amd64: # Windows Support
	GOOS=windows GOARCH=amd64 go build -v -a -o release/windows/amd64/file-compare.exe

build_macos_amd64: # Apple Intel Support
    GOOS=darwin GOARCH=amd64 go build -v -a -o release/macos/arm64/file-compare

build_macos_arm64: # Apple Silicon Support
	GOOS=darwin GOARCH=arm64 go build -v -a -o release/macos/arm64/file-compare

build_linux_amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -o release/linux/amd64/file-compare

build_linux_i386:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -v -a -o release/linux/i386/file-compare
