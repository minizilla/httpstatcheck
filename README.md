# httpstatcheck [![Go Reference](https://pkg.go.dev/badge/github.com/minizilla/httpstatcheck.svg)](https://pkg.go.dev/github.com/minizilla/httpstatcheck) [![main](https://github.com/minizilla/httpstatcheck/actions/workflows/main.yaml/badge.svg)](https://github.com/minizilla/httpstatcheck/actions/workflows/main.yaml)

Package httpstatcheck provides wildcard HTTP Status Code checking, e.g. 200 == 2XX.

## Example

```go
var checker httpstatcheck.Checker // The zero value for Checker is an empty rule checker ready to use.
checker.Insert("2XX", "400", "500", "3X1")
fmt.Println(checker.Check(http.StatusOK) // true
fmt.Println(checker.Check(http.StatusCreated)) // true
fmt.Println(checker.Check(http.StatusBadRequest)) // true
fmt.Println(checker.Check(http.StatusUnauthorized)) // false
fmt.Println(checker.Check(http.StatusInternalServerError)) // true
fmt.Println(checker.Check(http.StatusNotImplemented)) // false
fmt.Println(checker.Check(http.StatusMultipleChoices)) // true 3X1 will be considered as 3XX
```

## Benchmark with regex

```sh
$ go test -bench=.
goos: linux
goarch: arm64
pkg: github.com/minizilla/httpstatcheck
BenchmarkHTTPStatCheck-4   	13465224	        84.36 ns/op
BenchmarkRegex-4           	10013356	       120.2 ns/op
PASS
ok  	github.com/minizilla/httpstatcheck	2.559s
```
