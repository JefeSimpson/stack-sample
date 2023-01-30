package main

import (
	"fmt"
	"package-includes/packageImports"
)

func main() {
	//fmt.Println("Hello World!")
	st := &pkgimpt.Stack{}
	st.Insert(3)
	st.Insert(2)
	st.Insert(1)
	//st.Clear()

	fmt.Printf("stack size: %v\n", st.Size())
	fmt.Printf("stack top el: %v\n", st.Top())
	fmt.Println(st.Contains(4))
	fmt.Printf("stack size: %v\n", st.Size())
	st.Increment()
	fmt.Printf("stack top el: %v\n", st.Top())
	st.Print()
	st.PrintReverse()
	fmt.Printf("stack deleted el: %v\n", st.Pop())
	fmt.Printf("stack size: %v\n", st.Size())
	fmt.Printf("stack top el: %v\n", st.Top())
	fmt.Printf("stack deleted el: %v\n", st.Pop())
	fmt.Printf("stack size: %v\n", st.Size())
	fmt.Printf("stack top el: %v\n", st.Top())
	fmt.Printf("stack deleted el: %v\n", st.Pop())
	fmt.Printf("stack size: %v\n", st.Size())
	fmt.Printf("stack deleted el: %v\n", st.Pop())
	fmt.Printf("stack size: %v\n", st.Size())
}
