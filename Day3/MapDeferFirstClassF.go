package main

import "fmt"

func main(){
	defer fmt.Println("so this was the use of map and defer")
	b:=make(map[int ] string)
	b[1]="raju"
	b[2]="ramesh"
	b[3]="suresh"
	b[4]="prakash"
	b[5]="akash"
	defer fmt.Println(b)
	defer fmt.Println("this was the old map values but since map is refrence type it has new values only")
	print(b)
	func (b map[int]string , c int){
		b[c]="rakesh" //for updation using first class fucntion
	}(b,4)
	print(b)
	func (b map[int]string,c  int){
		delete(b,c) // for deletion using first class function
	}(b,4)
	print(b)

}
func print(b map[int]string) {
	for x,y:=range b{
		fmt.Println("roll no",x ,"is ",y )
	}
}

