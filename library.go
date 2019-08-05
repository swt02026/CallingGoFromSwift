package main
import "fmt"
import "C"
func main() {}

//export Add
func Add(a, b:int) int {
	return a+b
}
