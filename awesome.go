package main

import "C"

import (
	"fmt"
	"math"
	"os"
	"sort"
	"sync"
)

var count int
var mtx sync.Mutex

//export Add
func Add(a, b int) int {
	return a + b
}

//export Cosine
func Cosine(x float64) float64 {
	return math.Cos(x)
}

//export Sort
func Sort(vals []int) {
	sort.Ints(vals)
}

//export SortPtr
func SortPtr(vals *[]int) {
	Sort(*vals)
}

//export Log
func Log(msg string) int {
	mtx.Lock()
	defer mtx.Unlock()
	fmt.Println(msg)
	count++
	return count
}

//export LogPtr
func LogPtr(msg *string) int {
	return Log(*msg)
}

//export PrintProgramArguments
func PrintProgramArguments() int {
	argv := os.Args
	fmt.Printf("Program has %d arguments\n", len(argv))
	for i, arg := range argv {
		fmt.Printf("arg %d: %s\n", i, arg)
	}

	return len(argv)
}

//export PrintProgramEnvironment
func PrintProgramEnvironment() int {
	envv := os.Environ()
	fmt.Printf("Program has %d environment variables\n", len(envv))
	for i, env := range envv {
		fmt.Printf("env %d: %s\n", i, env)
	}

	return len(envv)
}

func main() {}
