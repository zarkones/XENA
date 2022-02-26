rm -r build

CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="-extldflags=-static" -o build/main_linux_64
strip build/main_linux_64

CGO_ENABLED=0 GOARCH=386 GOOS=linux go build -ldflags="-extldflags=-static" -o build/main_linux_32
strip build/main_linux_32

CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -ldflags="-extldflags=-static" -o build/main_mac_64
