rm -r build

CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="-extldflags=-static" -o build/setup_linux_64bit
strip build/setup_linux_64bit

CGO_ENABLED=0 GOARCH=386 GOOS=linux go build -ldflags="-extldflags=-static" -o build/setup_linux_32bit
strip build/setup_linux_32bit

CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -ldflags="-extldflags=-static" -o build/setup_mac_64bit
