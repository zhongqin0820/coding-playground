# CREATED
29, May, 2019

# Preface
Practice TDD in Golang. Includes:

1. Basics about `go test`. 
    - related commands

2. Advanced usage
    - unit test
    - benchmark test
    - race test
    - mock test

# `go test`
## basics
In `var t *testing.T`, `t` has many usages, one is logging! There are three categories and each has three functional of its way.

```go
t.Log
t.Logf
t.Error
t.Fatal
```

### Commands
- `-v`: verbose flag
- `-cover`: performs converage analysis
- `-race`: race detection analysis
- `-run=XXX`: XXX is the specific func name in unit test, `none` means ignoring all unit test
- `-bench=YYY`: YYY is the specific func name in benchmark test

### Frequency Usages
- `go test .`: files in current folder, not includes subfolder
- `go test ./..`: files in father folder
- `go test XXX_test.go`: specific files in current folder
- `go test -test.run TestXXX`: specific func in current folder
- `go test -run TestXXX`: specific func in current folder
- `go test ./YYY -run TestXXX`: specific func in specific folder

## unit test
Compare **the output** of the `func` with the **expected result** which needs predefined. `Fail` if not equal else `OK` in the final report.

- basic tests 
- table tests: the advanced version of basic tests with the help of slice!

> func name should start with `func TestXXX(t *testing.T) {...}`

## benchmark test

> func name should start with `func BenchmarkYYY(b *testing.B) {...}`
> `b.ResetTimer()`
> `for i:=0;i<b.N;i++{...}`


# TODO
- [ ] godoc: the comments in the test file
- [ ] race test
- [ ] mock test in details

# Reference
- __Go in Action__
