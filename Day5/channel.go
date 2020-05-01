package main

import (
	"fmt"
	"net/http"
)


func main(){

	links:=[]string{
		"https://www.google.com/",
		"https://www.facebook.com",
		"https://www.yahoo.com",
	}
	c := make( chan int ,3)


	go hitURL(links,c)
	for i:=0;i<3;i++ {
		b := <-c
		fmt.Println("status code is ",b)
	}

}
func hitURL(l []string,ch chan int){
	for i:= range l {
		res, err := http.Get(l[i])
		if err != nil {
			fmt.Println(err)
			ch <- 0
		} else {
			ch <- res.StatusCode
		}
	}
}
