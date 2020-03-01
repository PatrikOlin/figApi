package main

import (
	"strings"
	"figApi/datastore"
	"figApi/util"
	"strconv"
	"math/rand"
	"time"										  
)

type Person struct {
	Name     string
	Pin      string
	Email    string
	Address  string
	Phone    string
	Password string
}

func fetchPeople(amount int) []Person {
	
	var people []Person
		for i := 1; i <= amount; i++ {
			people = append(people, generatePerson())
		}

	return people
}

func generatePerson() Person {
	fullname := getFullName()
	email := getEmailForName(fullname)
	person := Person{
		Name: fullname,
		Pin: getPIN(),
		Address: getFullAddress(),				  
		Phone: getPhoneNumber(),				  
		Email: email,
		Password: getPassword(),
	}

	return person								  
}

func getFullName() string {
	var fullName strings.Builder
	fullName.WriteString(datastore.GetRandomLine("firstnames"))
	fullName.WriteString(" ")
	fullName.WriteString(datastore.GetRandomLine("surnames"))

	return fullName.String()
}
												  

func getEmailForName(fullname string) string {	  
	names := strings.Fields(fullname)
	var emailAddress strings.Builder
	emailAddress.WriteString(names[0])
	emailAddress.WriteString(".")
	emailAddress.WriteString(names[1])
	emailAddress.WriteString(datastore.GetRandomLine("emaildomains"))

	return strings.ToLower(emailAddress.String())
}

func getPhoneNumber() string {
	rand.Seed(time.Now().UnixNano())
	numbers := util.RangeIn(10000000, 99999999)		  
	initial := util.RangeIn(0, 9)				  
	phoneNum := "07" + strconv.Itoa(initial) + "-" + strconv.Itoa(numbers)

	return phoneNum
}												  

func getPassword() string {
	return datastore.GetRandomLine("passwords")	  
}
