README
==========================================================================================

Synopsis
------------------------------------------------------------------------------------------

This is a GO scraper based on [Ruby version](https://github.com/Krugloff/serpapi-code-challenge) that I create for [SerpApi code challenge](https://github.com/serpapi/code-challenge).

Motivation
------------------------------------------------------------------------------------------

+ For example, I want to check that 100000 accessories prices was not changed. I don't need to parse full html tree, I just need to find one+ prices on the each page and compare with the prices I have (hourly/daily).

+ How to check all possible variants like carousel, masonry, anyhting else quickly? We can use one complex sequential scanner/tokenizer that will check all known patterns (and skip most of the symbols ideally) or we can use one simple sequential scanner per variant. It will give us possibility to configure scanning based on user request (for example in case of GraphQL api). Although you can configure complex sequential scanner too..

+ I imagine it even can be splitted on multiple microservices/goroutines, each of one will scan his own block (one microservice for carousel, one microservice for masonry, one microservice for wikipedia card)

Installation
------------------------------------------------------------------------------------------

This is not a production ready package so you need clone the source code from github to your local machine.

Benchmark
------------------------------------------------------------------------------------------

### scanner based

```
> go test -bench=. -count=10 -benchtime=1000x  -benchmem -cpuprofile cpu.out ./scanner_based/knowledge_graph/

goos: darwin
goarch: arm64
pkg: serpapi-code-challenge-go/scanner_based/knowledge_graph

BenchmarkJSON-8         1000      257696 ns/op    553772 B/op      482 allocs/op
BenchmarkJSON-8         1000      256169 ns/op    553745 B/op      482 allocs/op
BenchmarkJSON-8         1000      256005 ns/op    553745 B/op      482 allocs/op
BenchmarkJSON-8         1000      255374 ns/op    553747 B/op      482 allocs/op
BenchmarkJSON-8         1000      255174 ns/op    553746 B/op      482 allocs/op
BenchmarkJSON-8         1000      258956 ns/op    553744 B/op      482 allocs/op
BenchmarkJSON-8         1000      257402 ns/op    553747 B/op      482 allocs/op
BenchmarkJSON-8         1000      255842 ns/op    553747 B/op      482 allocs/op
BenchmarkJSON-8         1000      257016 ns/op    553744 B/op      482 allocs/op
BenchmarkJSON-8         1000      254809 ns/op    553747 B/op      482 allocs/op
PASS
ok    serpapi-code-challenge-go/scanner_based/knowledge_graph 2.798s
```

0.26s per 1000 pages vs 1.32s for Ruby

### regexp based

```
> go test -bench=. -count=10 -benchtime=1000x  -benchmem -cpuprofile cpu.out ./regexp_based/knowledge_graph/

goos: darwin
goarch: arm64
pkg: serpapi-code-challenge-go/regexp_based/knowledge_graph
BenchmarkJSON-8         1000    11444451 ns/op    576868 B/op      839 allocs/op
BenchmarkJSON-8         1000    11530694 ns/op    576204 B/op      838 allocs/op
BenchmarkJSON-8         1000    11544776 ns/op    575760 B/op      839 allocs/op
BenchmarkJSON-8         1000    11477662 ns/op    575325 B/op      837 allocs/op
BenchmarkJSON-8         1000    11619159 ns/op    576982 B/op      838 allocs/op
BenchmarkJSON-8         1000    11499895 ns/op    575888 B/op      838 allocs/op
BenchmarkJSON-8         1000    11472286 ns/op    576546 B/op      836 allocs/op
BenchmarkJSON-8         1000    11565138 ns/op    575352 B/op      837 allocs/op
BenchmarkJSON-8         1000    11492218 ns/op    575757 B/op      837 allocs/op
BenchmarkJSON-8         1000    11520657 ns/op    575991 B/op      836 allocs/op
PASS
ok    serpapi-code-challenge-go/regexp_based/knowledge_graph  115.824s
```

Profiling
------------------------------------------------------------------------------------------

You can use

```
> go tool pprof knowledge_graph.test cpu.out
> web
```

Testing
------------------------------------------------------------------------------------------

Tests will compare expected result (from the initial serp api challenge files) and parsed result.

```
> go test ./regexp_based/knowledge_graph/
> go test ./scanner_based/knowledge_graph/
```

TODO
------------------------------------------------------------------------------------------

+ try to use alternative [regexp](https://pkg.go.dev/github.com/flier/gohs/hyperscan) libraries
+ standard JSON marshalling working slow because based on reflection. Try to use alternatives.
+ completely repeat Ruby realization and all tests. Especially "nokogiri" based variant and "element not found" cases.
+ CI