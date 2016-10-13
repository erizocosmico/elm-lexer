# elm-lexer

Elm language lexer in Go.

### Install

```
go get github.com/mvader/elm-lexer
```

### Usage 

```go
l := elmlexer.New(os.Stdin)
// Run needs to be run in a goroutine, generating tokens in parallel
go l.Run()

for {
    token, ok := l.Next()
    if !ok {
        break
    }

    // do stuff with token
}
```

### `elmlex` cmd

There is also a test command that lexes an Elm file and prints the generated tokes. You can use it like this:

```
cd $GOPATH/github.com/mvader/elm-lexer
go build cmd/elmlex.go
./elmlex < /path/to/file.elm
```

### Note

This library is, for now, just an experiment, but I may be doing an Elm-to-Go compiler in the future as a personal project.