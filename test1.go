package main

import (
	"fmt"
	"strings"
)
type student struct {
	a int
	b int
}

func main(){
	a:="a b c d e"
	b := strings.Split(a," ")
	fmt.Print(b)
	fmt.Println(b[1])
	A := &student{
		a :1,
		b:2,
	}
	fmt.Println(A.a)
	fmt.Println((*A).a)



}

