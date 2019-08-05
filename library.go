package main
import "fmt"
import "C"
func main() {}

//export Add
func Add(a:int, b:int) {
	return a+b
}
