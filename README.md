# Advent of code 2023

<p align="center">
  <img src="./static/logo.svg" width="220" />
</p>

Here you can find my golang solutions for [Advent Of Code 2023](https://adventofcode.com).

## Current results

<!--- advent_readme_stars table --->

## Codegen

`days/day*` codegen is a slightly changed version of [alexchao26](https://github.com/alexchao26/advent-of-code-go) aoc repo.

### Workflow

To generate skeleton for another day run:

```console
$ make gen DAY=3
```

Then manually:
1. Fill tests with example input and want values
2. Fill input.txt with real data
3. Code problem solution

Another commands:

```console
$ make run DAY=03 PART=2
$ make test DAY=03
```
