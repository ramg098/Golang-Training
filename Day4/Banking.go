package main

import "fmt"

type Account interface {
	withdraw(float32)
	deposit(float32)
	checkBalance()
	addInterest()
}

type acc struct {
	accNo int
	accBalance float32
	accType bool
	interest float32
}

func (ac *acc)withdraw (a float32){
	if a > ac.accBalance{
		fmt.Println("try again with amount less than or equal to your account balance")
	}else{
	ac.accBalance=ac.accBalance - a
	}
}

func (ac *acc)deposit( a float32){
	ac.accBalance=ac.accBalance+a
}

func (ac *acc) checkBalance(){
	fmt.Println("your balance is ",ac.accBalance)
}

func (ac *acc) addInterest(){
	if ac.accType==true {
		ac.accBalance =ac.accBalance +(ac.accBalance * ac.interest/100)
	}else{
		fmt.Println("no interest for current account")
	}
}

func main(){
	var ac Account=&acc{1,15000,true,4}
	var ac1 Account=&acc{2,12000,false,0}
	ac.checkBalance()
	ac.deposit(4100)
	ac.checkBalance()
	ac.addInterest()
	ac.checkBalance()
	ac.withdraw(2100)
	ac.checkBalance()
	ac1.checkBalance()
	ac1.deposit(2500)
	ac1.addInterest()
	ac1.checkBalance()
	ac1.withdraw(2000)
	ac1.checkBalance()
}
