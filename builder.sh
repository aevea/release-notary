for GOOS in darwin linux windows; do
    for GOARCH in 386 amd64; do
        go build -v -o build/release-notary-$GOOS-$GOARCH
    done
done