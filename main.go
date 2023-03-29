package main

import (
	"awesomeProject/floor"
	"awesomeProject/idWorker"
	"fmt"
)

func main() {
	fmt.Print("WELCOME to our company\nWhich floor do tou need?")
	var a int
	fmt.Scan(&a)
	floor.OurFloor(a)
	fmt.Print("Please, Enter the number your ID\nHite\nYour ID number must consist of 6 number ")
	var id int
	fmt.Scan(&id)
	idWorker.IdOfPerson(id)
}

//fmt.Println("CEKCK YOUR ID")
//var id int
//fmt.Scan(&id)
//if id >= 100000 && id < 200000 {
//fmt.Println("You Can go only to first floor")
//fmt.Println("SECOND floor STOP")
//fmt.Println("THIRD floor STOP")
//} else if id >= 200000 && id < 300000 {
//fmt.Println("YOU CAN GO TO FIRST")
//fmt.Println("YOU CAN GO TO SECOND")
//fmt.Println("YOU CAN'T GO TO THIRD")
//} else if id >= 300000 && id < 400000 {
//fmt.Println("YOU CAN GO TO FIRST")
//fmt.Println("YOU CAN GO TO SECOND")
//fmt.Println("YOU CAN GO TO THIRD")
//} else {
//fmt.Println("YOUR ID NOT CORRECT FO OUR COMPANY")
//}

//var age int
//switch age {
//case 20:
//fmt.Println("Pay")
//case 40:
//fmt.Println("Pay")
//case 50:
//fmt.Println("Pay")
//case 60:
//fmt.Println("Pay")
//default
//}

//var age int
//var gender string
//if age  >= 18 && gender == "female" {
//fmt.Println("Ok")
//} else {
//fmt.Println("No")
//}

//var age int
//var gender string
//if age >= 60 || gender == "female" {
//fmt.Println("Ok")
//} else {
//fmt.Println("No")
//}

//var gender string
//if  gender != "female" {
//fmt.Println("Ok")
//} else {
//fmt.Println("No")
//}

//var age string
//if  age < 18 {
//fmt.Println("Ok")
//} else if age == 20 {
//fmt.Println("No")
//}
