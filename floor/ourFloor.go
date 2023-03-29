package floor

import "fmt"

func OurFloor(a int) {
	switch a {
	case 1:
		fmt.Println("Welcome to first floor")
	case 2:
		fmt.Println("Welcome to second floor")
	case 3:
		fmt.Println("Welcome to thrid")
	default:
		fmt.Println("We have not this floor")
	}
}
