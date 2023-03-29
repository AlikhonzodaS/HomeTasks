package idWorker

import "fmt"

func IdOfPerson(a int) {
	if a >= 10000 && a < 20000 {
		fmt.Println("You can go to first floor")
	} else if a >= 10000 && a < 30000 {
		fmt.Println("You can go to second floor")
	} else if a >= 10000 && a < 40000 {
		fmt.Println("You can go to thirth floor")
	} else {
		fmt.Println("Dont Come")
	}
}
