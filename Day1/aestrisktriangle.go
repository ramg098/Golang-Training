package main
import "fmt"

func main(){

	fmt.Println("enter no. of rows ")
	var i int
	fmt.Scanln(&i)

	for j:=0 ; j < i ; j++ {
		for k:=0;k<=j;k++ {

			fmt.Print("* ")
		}
		fmt.Println("")
	}


}

