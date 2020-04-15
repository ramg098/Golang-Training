package main
import "fmt"

func main() {

	fmt.Println("enter no. of rows ")
	var i int
	fmt.Scanln(&i)
	k := 0
	for j := 1; j < i; j++ {
		if i%j == 0 {
			k++;
		}
	}
	if k > 1 {
		fmt.Println("not prime")
	} else
	{
		fmt.Println("prime")
	}
}
