env CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -a -o csplistenerLinux
env CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -a -o csplistenerMac

# to run the program
# ./cspListenerMac