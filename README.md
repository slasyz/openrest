# OpenREST

OpenREST generates Golang server boilerplate code based on OpenAPI document (or, to be more precise, its subset).

It is in pre-alpha stage, so it is not intended for production usage.

## Installation

```bash
go get github.com/slasyz/openrest
```

## Usage

```golang
//go:generate ./tools/openrest -outDir ./internal/server -srcFile ./openapi.yml
```

It will read `openapi.yml` file and create files in `./internal/server` directory such as:
 * `generated/dto.go` with structs;
 * `generated/handler.go` with `http.Handler` implementation;
 * `methods/methods.go` with methods implementation constructor;
 * `methods/error.go` with method example that will be called for error handling (logging and creating response);
 * set of files `methods/XXX.go` for each path where `XXX` is path converted to snake case; these files contain method implementation stubs.

All files in `generated` directory will be overwritten after each OpenREST call.  You are free to modify all files in `methods` directory.


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
