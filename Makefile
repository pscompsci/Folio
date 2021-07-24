all: dist/folio_macos_darwin_amd64 dist/folio_linux_amd64 dist/folio_windows_X86_64

clean:
	rm -rf dist/

dist/folio_macos_darwin_amd64: *.go
	GOOS=darwin GOARCH=amd64 go build -o dist/folio_macos_darwin_amd64

dist/folio_linux_amd64: *.go
	GOOS=linux GOARH=amd64 go build -o dist/linux_amd64

dist/folio_windows_X86_64: *.go
	GOOS=windows GOARCH=amd64 go build -o dist/folio_windows_X86_64.exe
