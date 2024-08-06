# httpstatcheck [![Go Reference](https://pkg.go.dev/badge/github.com/minizilla/httpstatcheck.svg)](https://pkg.go.dev/github.com/minizilla/httpstatcheck) [![main](https://github.com/minizilla/httpstatcheck/actions/workflows/main.yaml/badge.svg)](https://github.com/minizilla/httpstatcheck/actions/workflows/main.yaml)

Package httpstatcheck provides wildcard HTTP Status Code checking, e.g. 200 == 2XX.

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
