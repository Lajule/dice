# dice

Roll a dice from command line.

## Installation

```sh
go get github.com/Lajule/dice
```

## Syntax

`dice` uses the following syntax:
* `<draws>`**d**`<faces>`
* `<draws>`**d**`<faces>`**+**`<modifier>`
* `<draws>`**d**`<faces>`**-**`<modifier>`

## Examples

Consider following input file `draws`:

```
3d6
6d3+10
```

`dice` binary should be used like this:


```sh
dice draws
11
20
```

or using _Here Documents_ syntax :

```sh
dice <<EOF
> 3d6
> 6d3+10
EOF
11
20
```

or even piping command result :

```sh
{
> echo 3d6
> echo 6d3+10
} | dice
11
20
```
