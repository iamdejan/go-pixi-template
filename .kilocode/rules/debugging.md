# Debugging

In this repository, we don't have Go language installed in local machine. Instead, `go` executable is managed by Pixi, which is why in [pixi.toml](../../pixi.toml) under `[dependencies]` section you will have `go` listed as dependency.

The consequence is that every Zig command you want to run should be run under `pixi run` command. I will list some of the examples here:
- Check Go version: Instead of `go version`, you should run `pixi run go version`.
- Run all unit tests: Instead of `go test ./... -v`, you should run `pixi run go test ./... -v`.
  - In the case of running all unit tests, you should follow the rule under [code_validation.md](./code_validation.md).
- Run the program: Instead of `go run main.go`, you should run `pixi run go run main.go`.
  - In the case of running the program (or the main function), you should execute `pixi run start` instead, following the task defined in [pixi.toml](../../pixi.toml).
