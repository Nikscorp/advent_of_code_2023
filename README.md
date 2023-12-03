# Advent of code 2023

Here you can find my golang solutions for Advent Of Code 2023.

## Codegen

`days/day*` codegen is a slightly changed version of [alexchao26](https://github.com/alexchao26/advent-of-code-go) aoc repo.

## Workflow

To generate skeleton for another day run:

```bash
make gen DAY=3
```

Then manually:
1. Fill tests with example input and want values
2. Fill input.txt with real data
3. Code problem solution


Another commands:

```bash
make run DAY=03 PART=2
make test DAY=03
```
