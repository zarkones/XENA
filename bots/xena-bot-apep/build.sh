rm -r build

go build -ldflags="-w -s" -o build/main_linux_64
strip build/main_linux_64

GOOS=windows GOARCH=386 go build -ldflags="-w -s" -o build/main_win_32
strip build/main_win_32
