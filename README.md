# dice

Roll a dice in command line

## Installation

```sh
go get github.com/Lajule/dice
```

## Usage

Type the following command `dice -h` to display this help message:

```
qsfc [flags] file
  -d	Displays draws details
```

## Syntax

`dice` use the following syntax:
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
dice -d draws
3d6:(2+4+5)=11
6d3+10:(1+1+3+2+2+1)+10=20
```

or using _Here Documents_ syntax :

```sh
dice <<EOF
3d6
6d3+10
EOF
3d6:11
6d3+10:20
```

or piping command result :

```sh
echo '3d6\n6d3+10' | dice
3d6:11
6d3+10:20
```
