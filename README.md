# Advent of Code 2022

This repository contains my solutions for the [Advent of Code 2022][1], which I have tried to solve using [Go][2]. The repository contains a basic setup using `go mod` and provides a [Nix][3] flake to develop.

## Usage

To install run the module:

```shell
nix develop
go build .
./aoc -day={1-25}
```


[1]: https://adventofcode.com/2022
[2]: https://go.dev
[3]: https://nixos.org