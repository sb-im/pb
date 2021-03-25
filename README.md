# Protocol buffer definitions shared across projects

This is a central repo hosting protocol buffers. All packages contain the generated code for any desired languages.

# How to use the generated package.

For Go, follow these steps:

```bash
$ go env -w GOPRIVATE="github.com/SB-IM"
```

```bash
$ git config --global url."git@github.com:".insteadOf "https://github.com/"
```

Then, you can get it with (use signal package for example):

```bash
$ go get github.com/SB-IM/pb/signal
```
