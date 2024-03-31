package calculate

import (
	"fmt"
	"time"
)

func SwitchStatement() {
	today := time.Now()
	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th. Clean your house.")
		break
	case 8:
		fmt.Println("Today is 10th. Buy some wine.")
		fallthrough                                  // Case 15 will also execute
	case 15:
		fmt.Println("Today is 15th. Visit a doctor.")
	case 25:
		fmt.Println("Today is 25th. Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}
func DeferImplementation(){

	defer fmt.Println("world")
	
	fmt.Println("hello")
}
