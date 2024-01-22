$env.GOOS = windows;
$env.GOARCH = amd64;
go build -trimpath -ldflags="-s -w" -o bin\main.exe main

$env.GOOS = linux;
go build -trimpath -ldflags="-s -w" -o bin\main main
