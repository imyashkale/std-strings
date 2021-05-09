package models

import (
	"fmt"
	"strings"
	"time"
)

// Net worth
type NetWorth struct {
	Stocks []string
	Bonds  []string
	Cash   int32
}

// Profile Information
type ProfileInfo struct {
	FirstName  string
	LastName   string
	DOB        string
	City       string
	FavActress string
	Age        int
	Tax        int32
	CreatedAt  time.Month
	NetWorth
}

// Get the Full Name
func (p *ProfileInfo) GetFullName() string {
	return fmt.Sprintf("%v %v", p.FirstName, p.LastName)
}

// Change First Name
func (p *ProfileInfo) ChangeFirstName(NewFirstName string) {
	p.FirstName = NewFirstName
}

func (p *ProfileInfo) UpperCaseName() (string, string) {
	return strings.ToUpper(p.FirstName), strings.ToUpper(p.LastName)
}

// Change the city name
func ModifyCityName(p *ProfileInfo, newCity string) {
	fmt.Printf("City changed from %v to %v \n", p.City, newCity)
	p.City = newCity

}

func CalculateTaxOnCash(cash *ProfileInfo) {
	cash.Tax = (cash.NetWorth.Cash * 30 / 100)
}
