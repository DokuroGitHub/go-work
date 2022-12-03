# init mod
```go
go mod init github.com/DokuroGitHub/go-work
```

# create files
```bash
code main.go
```
# run
go run .

# test
go test
go test -fuzz=Fuzz -fuzztime=30s

# clone
git clone https://go.googlesource.com/example
git clone https://github.com/DokuroGitHub/go-crawl

# Initialize the workspace
go work init ./hello
go work use ./example
go work use ./go-crawl

# run module
go run example.com/hello