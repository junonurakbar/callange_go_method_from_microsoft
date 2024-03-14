package challange

import (
	"fmt"
	"testing"
)

func TestProgram(t *testing.T) {
	lele, _ := CreateEmployee("Lele", "Nihao", 0.0)

	plusCredit, err := lele.AddCredits(90.25)
	if err != nil {
		fmt.Println("invalid credit amount")
	} else {
		fmt.Printf("Added %.2f amount\n", plusCredit)
	}

	minusCredit, err := lele.RemoveCredits(21)
	if err != nil {
		fmt.Println("invalid credit amount")
	} else {
		fmt.Printf("Balance: %.2f\n", minusCredit)
	}
	
	lele.ChangeName("Nihao", "Ma")

	fmt.Println(lele.String())
}