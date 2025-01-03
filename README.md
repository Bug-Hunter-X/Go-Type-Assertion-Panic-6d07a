# Go Type Assertion Panic

This repository demonstrates a common error in Go: panics caused by incorrect type assertions.  Type assertions are powerful but require careful handling to avoid runtime crashes. The `bug.go` file contains code that will cause a panic.  The `bugSolution.go` file shows how to safely handle type assertions.

## Bug Description

In Go, type assertions are used to check the concrete type of an interface value. If the type assertion fails, a runtime panic occurs. This is a common source of unexpected errors in Go programs.

## Solution

The best practice is to always check the type of an interface value before performing a type assertion using a type switch.  This allows for graceful handling of cases where the type assertion might fail.
