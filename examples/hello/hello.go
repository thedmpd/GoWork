package main

import "fmt"

func main(){
	fmt.Printf("hello, world.\n")
	fmt.Println("Hello Word"[1])
	fmt.Printf("hi there Diogo\n")
	fmt.Println(42)
	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	for i := 0; i < 1000; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}

