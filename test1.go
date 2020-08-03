package main

import (
	"fmt"
	"strings"
)

func main(){
	a:="a b c d e"
	b := strings.Split(a," ")
	fmt.Print(b)
	fmt.Println(b[1])



}

