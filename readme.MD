## Introduction
Attempted to make Unix 'wc' like command in go 

## How to install
Run `go install go-wc` to create an executable file at $GOPATH/bin

## Usages

### Supported flags
`-l : For line counts`
`-w : For word counts`
`-c : For char counts`

### Examples
`go-wc fileName` - Will print line counts, word counts and char counts ending with fileName

`go-wc -l fileName` - Will print line counts ending with fileName

`go-wc -l -c fileName` - Will print line counts, char counts ending with fileName

`go-wc -l -c fileName1 fileName2` - Will print line counts, char counts ending with fileName for each file and then total counts
