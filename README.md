# sbt-compile-warn
Aggregate sbt compile warnings.

```
 $ sbt-compile-warn
[info] Loading project definition from ...
...
[warn] 8 warnings found
[info] Done compiling.
[success] Total time: 4 s, completed Oct 7, 2018 4:29:37 PM
4: Unused import
  main.scala:1:22
  main.scala:2:13
  main.scala:3:13
  main.scala:4:13

1: local val x in method main is never used
  main.scala:10:9

1: local var y in method main is never used
  main.scala:11:9

1: local val y in method oldMethod is never used
  main.scala:16:9

1: method oldMethod in object TestSbtCompileWarn is deprecated (since Sample 2.0): this method will be removed
  main.scala:11:13

Total: 8 warnings
```

## Installation
```sh
go get -u github.com/itchyny/sbt-compile-warn
```

## Author
itchyny (https://github.com/itchyny)

## License
This software is released under the MIT License, see LICENSE.
