# serialite
Lightweight Golang Serial/UART interface for Linux/OS X written entirely in
native Go.

## Usage

The folder structure is design such that this package can be fetched using

```bash
go get github.com/uw-midsun/serialite
```

And will automatically be available as an external library for any project.

## Development

For development this project is using the [Bazel](https://bazel.build/versions/master/docs/install.html)
build system. 

### Building

To build invoke

```bash
bazel build path/to:target
```

To update `BUILD` files first install [Gazelle](https://github.com/bazelbuild/rules_go/tree/master/go/tools/gazelle)
then run

```bash
gazelle -go_prefix github.com/uw-midsun/serialite
```

From then on you can just use

```bash
gazelle
```

To update `BUILD` files automatically.

### Testing

Testing is just as easy as building

```bash
bazel test path/to:target
```

### Tips and Tricks

A good understanding of all the functions of `bazel` makes development much
easier. More can be found [here](https://bazel.build/versions/master/docs/bazel-user-manual.html).

Here are a couple simple ones to get started:

```bash
bazel build path/to/...
```

Will glob all target that start with `path/to` and build them this also works
with the `test` action.

```bash
bazel test path/to:all
```

Will glob all targets in `path/to` and test them this also works with the 
`build` action.

## Dependencies

This project depends on [golang.org/x/sys](https://godoc.org/golang.org/x/sys)
which is a set of low level OS interface commands for Golang. It is well
maintained and supports Windows, OS X and Linux on various architectures.

## Future Features

These features are under consideration:

*   Windows Support
*   Hardware and Software Flow Control  
 
