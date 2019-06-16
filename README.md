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


parse tree

```
( let ( A 4 ) )

( let ( B ( + ( A 3 ) )))

( + ( A B ))

```
```
&{0xc420086600 0xc420086480  } &{0xc4200864b0 0xc420086540  }
&{0xc4200862d0 <nil>  } &{0xc420086300 0xc420086450  }
 &{0xc420086210 0xc4200862a0  }
&{<nil> <nil> let IDBind} &{0xc420086240 0xc420086270  }
let
&{<nil> <nil> A IDFree} &{<nil> <nil> 4 NumInt}
A
4
&{<nil> <nil> let IDBind} &{0xc420086330 0xc420086420  }
let
&{<nil> <nil> B IDFree} &{0xc420086360 0xc4200863f0  }
B
&{<nil> <nil> + OpSum} &{0xc420086390 0xc4200863c0  }
+
&{<nil> <nil> A IDFree} &{<nil> <nil> 3 NumInt}
A
3
&{<nil> <nil> + OpSum} &{0xc4200864e0 0xc420086510  }
+
&{<nil> <nil> A IDFree} &{<nil> <nil> B IDFree}
A
B
```
