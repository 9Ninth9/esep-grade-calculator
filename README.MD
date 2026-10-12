# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```

You need to download goland 1.24.4 exactly.
Then in your command line where all the files are located run: go test ./...

README.MD: Has information about the project and instructions for using it
go.mod: Defines the Go module and gives the Go version.
grade_calculator.go: Calculates grades, averages, weights, and letter grades.
grade_calculator_test.go: Has the unit tests to test if the code works.
