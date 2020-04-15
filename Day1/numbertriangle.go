package main
import "fmt"

func main(){

	fmt.Println("enter no. of rows ")
	var i int
	fmt.Scanln(&i)
	l:=0
	for j:=0 ; j < i ; j++ {
		for k:=0;k<=j;k++ {
			l++
			fmt.Print(l," ")
		}
		fmt.Println("")
	}


}

