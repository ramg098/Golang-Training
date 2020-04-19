package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Contact struct{
	mobile string
	email string
}
type Student struct {
	fname string
	lname string
	class string
	age int
	Contact
}

func main(){

	s:=Student{"ramesh","suresh","11",17,Contact{"98766","ramsh@suresh.com"}}
	s.print()
	s.UpdateClassAge("12",18)
	s.print()
	err:=sendToFile(s)
	if err == nil {
		fmt.Println("filesent")
	}else{
		fmt.Print(err)
	}
	err,byt:=readFromFile()
	if err == nil {
		fmt.Println("filerecieved")
		fmt.Print(byt)
	}else{
		fmt.Print(err)
	}


}

func (s Student) print (){
	fmt.Printf("%v\n",s)
}

func (s *Student) UpdateClassAge(c string,a int){
	s.class=c
	s.age=a
}

func (s Student) toByteSlice()[]byte{
	var str []string
	var stri string

	str = append(str,s.fname,s.lname,s.class,string(s.age),s.mobile,s.email)
	strings.Join(str,stri)
	var byt []byte=[]byte(stri)
	fmt.Print(str)
	return byt
}

func sendToFile(s Student)error{
	err:=ioutil.WriteFile("studentRecord",s.toByteSlice(),0666)

	return err
}

func readFromFile()(error,[]byte){
	byt,err:=ioutil.ReadFile("studentRecord")
	return err,byt
}



