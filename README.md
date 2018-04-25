# dice
Roll a dice in command line

## Usage

Consider following input file `draws`:

```
3d6
6d3+10
```

`dice` binary should be used like this:


```sh
dice draws
3d6:(2+4+5)=11
6d3+10:(1+1+3+2+2+1)+10=20
```

or using _Here Documents_ syntax :

```sh
dice <<EOF
3d6
6d3+10
EOF
3d6:(2+4+5)=11
6d3+10:(1+1+3+2+2+1)+10=20
```


```sh
echo '3d6\n6d3+10' | dice
3d6:(2+4+5)=11
6d3+10:(1+1+3+2+2+1)+10=20
```
