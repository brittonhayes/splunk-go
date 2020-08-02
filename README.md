<h1 align="center">Splunk GO REST API Client 🐹</h1>
<p>
  <a href="https://github.com/brittonhayes/splunk-golang/tree/master/docs/splunk-go.md" target="_blank">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" />
  </a>
</p>

> A Splunk REST API Client written in GO. This command line tool was designed to allow common tasks for Splunk usage and administration to be conducted through an easy-to-use command line tool.

### 🏠 [Homepage](https://github.com/brittonhayes/splunk-golang)

### ✨ [Examples](https://github.com/brittonhayes/splunk-golang/tree/master/examples)

## Install

```sh
git clone https://github.com/brittonhayes/splunk-golang
```

## Usage

```sh
go run main.go --help
```

## Run tests

```sh
go run main.go test
```

## Using the binary

Once the binary is compiled and added to your path, you can utilize the CLI to perform Splunk operations.

1. Build the binary

```shell
go build -o ./bin/splunk-go .

```

2. Run the help command to get a list of possible actions

```shell
splunk-go --help
```

This will output a message like the following: [Help Output](https://github.com/brittonhayes/splunk-golang/tree/master/docs/splunk-go.md)

## Adding features to the project

1. Navigate to the `pkg` directory of the repository and add a new `*.go` file with your added functionality.

2. Run `cobra add [command_name]` and the Cobra CLI tool will add a new `*.go` file for your command.

3. After creating your command, run `go docs` to automatically update the documentation of all CLI commands.

---

## Author

👤 **Britton Hayes**

- Github: [@brittonhayes](https://github.com/brittonhayes)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check the [issues page](https://github.com/brittonhayes/splunk-golang/issues). You can also take a look at the [contributing guide](https://github.com/brittonhayes/splunk-golang/tree/master/contributing.md).

## Acknowledgements

Social image from [Ashley Mcnamara](https://github.com/ashleymcnamara/gophers)

## Show your support

Give a ⭐️ if this project helped you!

---

_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
