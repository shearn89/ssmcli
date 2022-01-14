.PHONY: all

test:
	go test -v ./...

build:
	go build -v ./cmd/ssmcli

cross-compile: clean build-darwin-arm64 build-linux-amd64 build-linux-arm64 build-windows-amd64

build-darwin-arm64:
	env GOOS=darwin GOARCH=arm64 go build -o bin/darwin-arm64/ssmcli ./cmd/ssmcli

build-linux-amd64:
	env GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/ssmcli ./cmd/ssmcli

build-linux-arm64:
	env GOOS=linux GOARCH=arm64 go build -o bin/linux-arm64/ssmcli ./cmd/ssmcli

build-windows-amd64:
	env GOOS=windows GOARCH=amd64 go build -o bin/windows-amd64/ssmcli ./cmd/ssmcli

package: cross-compile
	tar -C bin/ -cvzf bin/ssmcli.linux-amd64.tar.gz --strip-components=1 linux-amd64/ssmcli
	tar -C bin/ -cvzf bin/ssmcli.linux-arm64.tar.gz --strip-components=1 linux-arm64/ssmcli
	tar -C bin/ -cvzf bin/ssmcli.darwin-arm64.tar.gz --strip-components=1 darwin-arm64/ssmcli
	zip -D bin/ssmcli.windows-amd64.zip bin/windows-amd64/ssmcli

clean:
	rm -rf bin/ ssmcli
