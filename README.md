# K
I'm making a language!
```
Lexeme List

identifier : [A-Za-z][A-za-z0-9]*
separator : ( | )
literal integer : [0-9]* 
literal string : "anything in quotes"
literal decimal : [0-9][.][0-9]* | [.][0-9][0-9]*
```

```
file.txt => 
0.42 42 .42 42. "42"
    (    forty2        4two)
```

```
$ go run *.go file.txt
```
```
0.42 - literal decimal
42 - literal integer
.42 - literal decimal
42. - literal decimal
"42" - literal string
( - separator
forty2 - identifier
4 - literal integer
two - identifier
) - separator
```
