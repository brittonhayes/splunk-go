<div align="center">
  <img alt="Banner Logo" src="https://user-images.githubusercontent.com/46035482/89118313-a134ee80-d459-11ea-952c-b4fc3c88aaf5.png" />
</div>

<h1 align="center">Splunk Golang REST API Client</h1>

> `splunk-go` is a Splunk REST API Client written in GO designed to allow common tasks for Splunk usage and administration to be conducted through an easy-to-use command line tool. This package assumes that you have a running instance of Splunk Enterprise with user credentials to access this instance. For all information related to the Splunk software, please visit https://www.splunk.com/en_us/software/splunk-enterprise.html

<p>
  <a href="https://pkg.go.dev/github.com/brittonhayes/splunk-go?tab=overview" target="_blank">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" />
  </a>
</p>

### 🏠 [Homepage](https://github.com/brittonhayes/splunk-go)

### ✨ [Examples](https://github.com/brittonhayes/splunk-go/tree/master/examples)

### :computer: [CLI Docs](https://pkg.go.dev/github.com/brittonhayes/splunk-go/cmd?tab=doc)

### :gift: [Package Docs](https://pkg.go.dev/github.com/brittonhayes/splunk-go/pkg?tab=doc)

## Install

```sh
# Clone the repository
git clone https://github.com/brittonhayes/splunk-go

# Install dependencies
make dependencies
```

## Usage

```sh
# From entrypoint
go run main.go --help

# To run with CI
make all
```

## Run tests

```sh
make security
```

## Using the binary

Once the binary is compiled and added to your path, you can utilize the CLI to perform Splunk operations.

1. Build the binary

```shell
# Build your OS's binary
make build

or

go build -o ./bin/splunk-go main.go

# Cross-compile for all systems
make cross-compile

```

2. Run the help command to get a list of possible actions

```shell
# Using the entrypoint
go run main.go --help

or

# Using the Makefile
make run
```

This will output a message like the following: [Help Output](https://github.com/brittonhayes/splunk-go/tree/master/docs/splunk-go.md)

## Adding features to the project

1. Navigate to the `pkg` directory of the repository and add a new `*.go` file with your added functionality.

2. Run `cobra add [command_name]` and the Cobra CLI tool will add a new `*.go` file for your command.

3. After creating your command, run `make docs` to automatically update the documentation of all CLI commands.

## Demo

[![asciicast](https://asciinema.org/a/QEHBxxaB45kzrKZm1QUS78rPu.svg)](https://asciinema.org/a/QEHBxxaB45kzrKZm1QUS78rPu)

---

## Author

👤 **Britton Hayes**

- Github: [@brittonhayes](https://github.com/brittonhayes)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check the [issues page](https://github.com/brittonhayes/splunk-go/issues). You can also take a look at the [contributing guide](https://github.com/brittonhayes/splunk-go/blob/master/contributing.md).

## Acknowledgements

Social image from [Ashley Mcnamara](https://github.com/ashleymcnamara/gophers)

Original Project written by [drewrm](https://github.com/drewrm/splunk-golang)

## Show your support

Give a ⭐️ if this project helped you!

---

_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
