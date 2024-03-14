package challange

import "fmt"

type Account struct {
	FirstName string
	LastName  string
}

func (a *Account) ChangeName(first, last string) {
	a.FirstName = first
	a.LastName = last
}

type Employee struct {
	Account
	credits float64
}

func (e *Employee) AddCredits(balance float64) float64 {
	e.credits += balance
	return e.credits
}
func (e *Employee) RemoveCredits(balance float64) float64 {
	e.credits -= balance
	return e.credits
}
func (e *Employee) GetCredits() float64 {
	fmt.Println(e.FirstName, e.LastName, "credit is:", e.credits)
	return e.credits
}
