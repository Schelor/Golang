// gorefer.go
package gorefer

import (
	"fmt"
)

func main() {
	//sli()
	//mapp()
	Invoke()
}

func sli() {
	Array_a := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	Slice_a := Array_a[2:4]
	fmt.Println(len(Slice_a))
	fmt.Println(cap(Slice_a))
	
	fmt.Println(Array_a)
}

func mapp() {
	//var numbers map[string] int
	numbers := make(map[string]int)
	
	numbers["one"] = 1
	numbers["two"] = 2
	
	fmt.Println("Number 2 is:", numbers )
}

func SumAndProduct(A, B int) (int, int) {
	return A+B, A*B
}

func Invoke() {
	x := 3
	y := 4
	
	
	xPlus, xTime := SumAndProduct(x, y)
	fmt.Println(xPlus, xTime)
}








