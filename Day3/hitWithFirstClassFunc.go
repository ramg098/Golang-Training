package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main(){

	a:=func (s string) int{
		res, err:=http.Get(s)
		if err!=nil{
			fmt.Println(err)
			os.Exit(1)
		}
		r,e:=ioutil.ReadAll(res.Body)
		if e!=nil{
			fmt.Println(e)
			os.Exit(1)
		}
		fmt.Print(r)
		return res.StatusCode
	}("http://www.facebook.com")


	fmt.Println(a)

}
