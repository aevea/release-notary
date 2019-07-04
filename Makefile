install_deps:
	- go mod download

build:
	- go build -o build/release-notary ./
build-all:
	- sh builder.sh