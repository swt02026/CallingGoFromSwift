package main
import "C"
import "fmt"
func main() {}

//export Add
func Add(a, b int) int {
	fmt.println("Add from go")
	return a+b
}
