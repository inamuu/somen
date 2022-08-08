# somen

SOMEN is a my BENRI(convenient) infra tool.

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
