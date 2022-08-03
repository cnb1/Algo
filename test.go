package main

import (
	"container/list"
	"fmt"
)

type Test struct {
	name string
	age  int
}

func main2() {
	// create a list of struct
	l := list.New()

	// add values
	p1 := Test{name: "connor", age: 20}
	p2 := Test{name: "john", age: 22}
	p3 := Test{name: "mike", age: 24}

	l.PushBack(p1)
	l.PushBack(p2)
	l.PushBack(p3)

	fmt.Println("l address: ", &l)
	fmt.Println("l dereferenced", l)
	fmt.Println()
	fmt.Println("l front value", l.Front())
	fmt.Println()
	fmt.Println("item2 front value", l.Front().Next())
	fmt.Println()

	item := l.Front()
	fmt.Println(*item)
	fmt.Println(item)

	// // print them before
	// fmt.Println("printing before")

	// for e := l.Front(); e != nil; e = e.Next() {
	// 	item := Test(e.Value.(Test))
	// 	fmt.Println(item)
	// }

	// fmt.Println()

	// // remove one
	// remove(*l)

	// // print again
	// fmt.Println("printing after")

	// for e := l.Front(); e != nil; e = e.Next() {
	// 	item := Test(e.Value.(Test))
	// 	fmt.Println(item)
	// }

	// then do the removal in a separate function
}

func remove(l list.List) {

	etemp := l.Front()
	etemp = etemp.Next()

	l.Remove(etemp)
}
