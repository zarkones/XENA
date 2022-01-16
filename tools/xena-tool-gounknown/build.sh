go build -ldflags="-w -s" -o build/gounknown_linux_64
strip build/gounknown_linux_64

GOOS=windows GOARCH=386 go build -ldflags="-w -s" -o build/gounknown_win_32
strip build/gounknown_win_32
