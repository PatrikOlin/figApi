package main

import (
	"figApi/datastore"
	"figApi/util"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Person struct {
	Name     string `json:"name"`
	Pin      string `json:"pin"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func fetchPeople(amount int) []Person {
	var wg sync.WaitGroup
	wg.Add(amount)
	var people []Person

	for i := 1; i <= amount; i++ {
		go func() {
			defer wg.Done()
			people = append(people, generatePerson())
		}()
	}
	wg.Wait()
	return people
}

func generatePerson() Person {
	fullname := getFullName()
	email := getEmailForName(fullname)
	person := Person{
		Name:     fullname,
		Pin:      datastore.GetRandomLine("safepins"),
		Address:  getFullAddress(),
		Phone:    getPhoneNumber(),
		Email:    email,
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
	var emailAddress strings.Builder
	rand.Seed(time.Now().UnixNano())
	num := util.RangeIn(0, 100)

	if num < 15 {
		emailAddress.WriteString(datastore.GetRandomLine("emailaddresses"))
		emailAddress.WriteString(datastore.GetRandomLine("emaildomains"))

		return strings.ToLower(emailAddress.String())
	} else {
		names := strings.Fields(fullname)
		emailAddress.WriteString(names[0])
		emailAddress.WriteString(".")
		emailAddress.WriteString(names[1])
		e := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
		result, _, _ := transform.String(e, emailAddress.String())

		return strings.ToLower(result + datastore.GetRandomLine("emaildomains"))
	}
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

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}
