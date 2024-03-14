package challange

import (
	"errors"
	"fmt"
)

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

func (e Employee) String() string {
	return fmt.Sprintf("Name: %s %s, Credits: %.2f\n", e.FirstName, e.LastName, e.credits)
}

func CreateEmployee(firstName, lastName string, credit float64) (*Employee, error) {
	return &Employee{Account{firstName, lastName}, credit,}, nil
}

func (e *Employee) AddCredits(balance float64) (float64, error) {
	if balance > 0.0 {
		e.credits += balance
		fmt.Printf("Deposit %.2f amount\n", balance)
		return e.credits, nil
	}
	return 0.0, errors.New("invalid credit amount")
}

func (e *Employee) RemoveCredits(balance float64) (float64, error) {
	if balance > 0.0 {
		e.credits -= balance
		fmt.Printf("Withdraw %.2f amount\n", balance)
		return e.credits, nil
	}
	return 0.0, errors.New("invalid credit amount")
}

func (e *Employee) GetCredits() float64 {
	return e.credits
}
