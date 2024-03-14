package challange

import (
	"fmt"
	"testing"
)

func TestProgram(t *testing.T) {
	lele := Employee{
		Account{
			FirstName: "Lele",
			LastName: "Nihao",
		},
		0,
	}
	lele.ChangeName("Nihao", "Blyat")
	lele.AddCredits(90.25)
	lele.RemoveCredits(21)
	lele.GetCredits()
	fmt.Println(lele)
}