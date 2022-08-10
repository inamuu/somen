# somen

SOMEN is a my BENRI(convenient) infra tool.

## Usage

```sh
Somen is a infra tool

Usage:
  somen [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  myip        Print my global ip address
  version     Print th version number of somen

Flags:
  -h, --help   help for somen

Use "somen [command] --help" for more information about a command.

```

## How to develop
### Setup

```sh
$ go mod init somen
$ go install github.com/spf13/cobra-cli@latest
$ go get -u github.com/spf13/cobra@latest
$ cobra-cli init --license MIT --viper=false
```

### Add a sub command

```sh
$ cobra-cli add TYPE_SUB_COMMAND_NAME
```
