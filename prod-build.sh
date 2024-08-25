rm -r export
mkdir export
# mkdir export/builder
mkdir export/c2
mkdir common/builder/static/agents

# Build Agent.
echo "Building Linux Agents"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -extldflags='-static'" -tags "netgo,linux,purego" -o common/builder/static/agents/linux_amd64 ./agent
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags="-s -w -extldflags='-static'" -tags "netgo,linux,purego" -o common/builder/static/agents/linux_386 ./agent
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="-s -w -extldflags='-static'" -tags "netgo,linux,purego" -o common/builder/static/agents/linux_arm ./agent
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w -extldflags='-static'" -tags "netgo,linux,purego" -o common/builder/static/agents/linux_arm64 ./agent
echo "Building Windows Agents"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -extldflags='-static' -H=windowsgui" -tags "netgo,windows,purego" -o common/builder/static/agents/windows_amd64 ./agent
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-s -w -extldflags='-static' -H=windowsgui" -tags "netgo,windows,purego" -o common/builder/static/agents/windows_386 ./agent

# Build Agent Builder.
# echo "Building Linux Agent Builder"
# CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -tags "netgo,linux" -o export/builder/linux_amd64 ./agent-builder
# CGO_ENABLED=1 GOOS=linux GOARCH=386 go build -ldflags="-s -w" -tags "netgo,linux" -o export/builder/linux_386 ./agent-builder
# CGO_ENABLED=1 GOOS=linux GOARCH=arm go build -ldflags="-s -w" -tags "netgo,linux" -o export/builder/linux_arm ./agent-builder
# CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -tags "netgo,linux" -o export/builder/linux_arm64 ./agent-builder
# echo "Building Windows Agent Builder"
# CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -tags "netgo,windows" -o export/builder/windows_amd64.exe ./agent-builder
# CGO_ENABLED=1 GOOS=windows GOARCH=386 go build -ldflags="-s -w" -tags "netgo,windows" -o export/builder/windows_386.exe ./agent-builder

# Build C2.
echo "Building Linux C2"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -extldflags='-static'" -tags "netgo,linux,purego" -o export/c2/linux_amd64 ./c2
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags="-s -w -extldflags='-static'" -tags "netgo,linux,purego" -o export/c2/linux_386 ./c2
CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="-s -w -extldflags='-static'" -tags "netgo,linux,purego" -o export/c2/linux_arm ./c2
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w -extldflags='-static'" -tags "netgo,linux,purego" -o export/c2/linux_arm64 ./c2
echo "Building Windows C2"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -extldflags='-static'" -tags "netgo,windows,purego" -o export/c2/windows_amd64.exe ./c2
# CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-s -w -extldflags='-static'" -tags "netgo,windows,purego" -o export/c2/windows_386.exe ./c2
echo "Building MacOS C2"
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w -extldflags='-static'" -tags "netgo,darwin,purego" -o export/c2/darwin_amd64 ./c2

# Build User Interface.
echo "Building Linux UI"
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -tags "netgo,linux" -o export/XENA_Linux_amd64 ./ui
# CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -tags "netgo,linux" -o export/XENA_Linux_arm64 ./ui
cp -r ui/xena-tools export/
# echo "Building Windows UI"
# CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -tags "netgo,windows"  -o export/XENA.exe ./ui
