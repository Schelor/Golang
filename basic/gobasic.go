// gobasic.go
package basic

import (
	"fmt"
)

func main() {
	
	//var for global variables
	var a, b, c int
	var d int = 10
	var e, f, g int = 1, 2, 3
	var vname = "vname"
	
	//simple var only use in function inside
	vname1, vname2, vname3 := "string", 1.9, true
	
	
	//constant variable: number, boolean, string
	const constName = 10
	const Pi float32 = 3.14
	
	const i = 1000
	const b bool = false
	
	
	//build-in basic type 
	var isActive bool
	var enable, disabled = true, false
	
	//number rune, int8, int16, int32, int64, byte, unit8, uint16, uint32
	//float32, float64
	var i rune = 100
	var f float32 = 122.00
	
	//string
	var hello = "string"
	var empty = ""
	
	s := "ss"
	m := "workd"
	a := s + m
	
	n := `hell
	adfasdf
	asdfaf`
	
	
}

func test() {
	var available bool
	valid := false
	available = true
	
	//string in func
	no, yes, maby := "no", "yes", "maybe"
	
	frenchHello := "bonjour"
	
	
}

func arr() {
	var arr [10]rune
	arr[0] = 42
	arr[1] = 13
	
	a := [3]int{1,2,3}
	
}

func sli() {
	var fslice []int
	slice := []byte{'a', 'b'}
}










