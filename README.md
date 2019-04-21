# Qvik's Golang Library

This is [Qvik's](https://qvik.com) open sourced Golang library. It aims to collect all the usual boiler plate code from projects into one place.

See the godoc documentation for the library contents.

## Usage

Enable the module support if planning to use Go1.11 modules:

```sh
export GO111MODULE=on
```

Install the library like so:

```sh
go get -u "github.com/qvik/go"
```

Or just a subpackage, eg:

```sh
go get -u "github.com/qvik/go/rest"
```

## License

The library is released under the [MIT license](LICENSE.md).

## Contributing

Pull requests are welcomed. Your code must pass through `gofmt+golint` with no issues.

Contact [Matti Dahlbom](mailto:matti@qvik.fi) if any questions arise.