# Example project for replicating `go mod vendor` bug: https://github.com/golang/go/issues/31524

## Steps to reproduce:

1. use windows
2. `setx GOROOT "C:\go"`
3. `setx GOPATH "C:\go\workspace"`
4. clone repository
5. type in `go get`
6. type in `go mod vendor`
