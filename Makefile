install_deps:
	- go mod download

build:
	- go build -o build/release-notary ./