# Parser

Generated using ANTLR

- Run `go test -v -cover ./ast` for Test or `Ayi test`
- Run `antlr4 -Dlanguage=Go while.g4` in `parser` folder if you updated the `while.g4`
- Run `antlr4 -Dlanguage=Go -visitor while.g4` in `parser` folder if you need visitor pattern

## Ref

Heroic Grammar 

- https://github.com/spotify/heroic/blob/master/heroic-parser/src/main/antlr4/com/spotify/heroic/grammar/HeroicQuery.g4

InfluxDB's InfluxQL

- https://github.com/influxdata/influxdb/tree/master/influxql

Hand written SQL parsers 

- https://blog.gopheracademy.com/advent-2014/parsers-lexers/

