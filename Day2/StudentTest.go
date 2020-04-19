package main

import "testing"

func TestAge(t *testing.T) {
	D:=Student{"ramesh","suresh","11",15,Contact{"98766","ramsh@suresh.com"}}

	if  D.age< 17 {
		t.Errorf("age is : %d. minimum age %d", D.age, 17)
	}


}
func TestConatctNotNIL(t *testing.T){
	D:=Student{"ramesh","suresh","11",15,Contact{"","ramsh@suresh.com"}}

	if  D.mobile == "" || D.email==""{
		t.Errorf("mobile and email cannot be nil")
	}
}
func TestUpdate(t *testing.T){
	D:=Student{"ramesh","suresh","11",15,Contact{"","ramsh@suresh.com"}}
	D.UpdateClassAge("9",13)
	if  D.class !="9" || D.age != 13 {
		t.Errorf("not updates correctly")
	}
}
