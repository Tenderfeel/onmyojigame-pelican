install:
	go install

start:
	go run .

buildmac:
	GOOS=darwin GOARCH=amd64 go build -o build/pelican-mac 

buildwin:
	GOODS=windows GOARCH=amd64 go build -o build/pelican-win.exe