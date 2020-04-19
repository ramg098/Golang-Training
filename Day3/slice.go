package main

import (
	"fmt"
	"os"
)

func main(){
	a:=[]int {1,2,3,4,5,6}
	fmt.Println(len(a))
	update(a,5,19)
	var b []int
	b=a
	b[5]=17
	fmt.Print(a)
	fmt.Print(b)
	for x:=0;x< len(a);x++{
		if a[x]!=b[x]{
			fmt.Println("not equal")
			os.Exit(1)
		}
		fmt.Println("equal")
	}

}
func update(a []int,b int,c int) {
	a[b-1]=c
}
