# dice
Roll a dice in command line

```sh
cat draws
3d6
6d3+10
```

```sh
dice draws
3d6:(2+4+5)=11
6d3+10:(1+1+3+2+2+1)+10=20
```

```sh
dice <<EOF
3d6
6d3+10
EOF
3d6:(2+4+5)=11
6d3+10:(1+1+3+2+2+1)+10=20
```


```sh
dice -e 3d6 -e 6d3+10
3d6:(2+4+5)=11
6d3+10:(1+1+3+2+2+1)+10=20
```
