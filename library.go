package main
import "C"
import "fmt"
func main() {}

//export Add
func Add(a, b int) int {
	fmt.Println("Add from go")
	return a+b
}
