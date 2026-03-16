# AGENTS.md

## Code Generation

Hi Kilo Code agent! Whatever model you are using, make sure you adhere to the guide on this file to generate the code.

### Rules of engagement

1. You're allowed to create functions, structs and impl blocks to generate the code.
2. With the exception of import statements and Cargo macros (allow and deny), explain the flow of the program using comments. Make sure that you write not only what the program is doing, but why. It will help me to judge your work.
3. You have to add documentation in the form of Godoc for functions (public and private) and structs. The documentation should contain what is the function for (basically the description), a brief summary of the steps, and input and output parameters. If your function has the ability to throw panic, please state it in the documentation as well.

### Debugging

In this repository, we don't have Go language installed in local machine. Instead, `go` executable is managed by Pixi, which is why in [pixi.toml](pixi.toml) under `[dependencies]` section you will have `go` listed as dependency.

The consequence is that every Zig command you want to run should be run under `pixi run` command. I will list some of the examples here:
- Check Go version: Instead of `go version`, you should run `pixi run go version`.
- Run all unit tests: Instead of `go test ./... -v`, you should run `pixi run go test ./... -v`.
- Run the program: Instead of `go run main.go`, you should run `pixi run go run main.go`.
  - In the case of running the program (or the main function), you should execute `pixi run start` instead, following the task defined in [pixi.toml](pixi.toml).

## Code Validation

Hi Kilo Code agent! Whatever model you are using, make sure you validate the code you generated.

### Steps to Validate

Run these commands in sequence:
1. `pixi run fmt`: format your code. This ensures that the code written by human and by you (coding agent) are consistent, following the style that Rust has provided.
2. `pixi run lint`: ensure no linter errors.
3. `pixi run lint-fix`: if there is any linter error, fix it with this command.
4. `pixi run lint`: recheck again, maybe there are linter errors that need manual fix.
5. `pixi run test`: build the code, then run all unit tests. This ensures that the code you generate pass all the unit tests.

## Update Documentation

After the validation is finished, update the project tree structure and file descriptions in README.md if needed. This is to ensure we always have updated documentation.
