# SIGNALS
This package tries to implement signal handlers for various operating systems.

It uses file name suffixs per Go's conditional build rules.
I first learned about that from Cheney's
[How to use conditional compilation with the go build tool](https://dave.cheney.net/2013/10/12/how-to-use-conditional-compilation-with-the-go-build-tool)
blog.

To see what will be built, run the following command from the root of the repository.

```shell
go list -f '{{.GoFiles}}' ./internal/signals
```

There's a catch-all file, `signal.go`, that will be built for unknown systems.
It supplies an implementation that panics with a useful message.