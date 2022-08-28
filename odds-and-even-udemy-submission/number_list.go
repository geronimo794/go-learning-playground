package main

import "fmt"

type numberList []int

func newNumberList() numberList {
	nL := numberList{}
	for i := 0; i < 11; i++ {
		nL = append(nL, i)
	}
	return nL
}

func (nL numberList) printOddEven() {
	for _, item := range nL {
		if item%2 == 0 {
			fmt.Println(item, " is even")
		} else {
			fmt.Println(item, " is odd")
		}
	}
}
