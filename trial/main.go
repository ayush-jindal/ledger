package main

import (
	"fmt"
	"io/ioutil"
)

const (
	A = 1
	b = 2
	c = 3
	d = "sfdsdf"
)

type Transaction struct {
	from string
	to string
	amt int
}

type MyInt int
func (mi MyInt) Double() MyInt { return mi*2 }

func main() {
	p1 := Point{1, 2}
	fmt.Println(p1)
	var mi_1 MyInt = 1
	fmt.Println(mi_1.Double())
	a := [3][1]int{{1},{2},{3}}
	b := a
	b[0][0] = 0
	fmt.Println(a)
	trns := make([]Transaction, 10)
	data, err := ioutil.ReadFile("transactions.json")
	fmt.Println(data, trns, err)
	var t Transaction = Transaction{}
	fmt.Println(t.from, t.to, t.amt)
	fmt.Scan(&t.from)
	fmt.Scan(&t.to)
	fmt.Scan(&t.amt)
	fmt.Println(t)
	fmt.Println([]rune("Hello world"))
	var x int32 = 1
	var y int16 = 2
	x = int32(y)
	fmt.Println(x)
	fmt.Println(d)
	for i := 0; i<10; i++ {
		fmt.Println("hello, ")
	}
	var s string
	fmt.Scan(&s)
	fmt.Println(s)
	s1 := []int{1,2,3}
	fmt.Println(len(s1), cap(s1))
}
