package main

import (
	"strings"
	"figApi/datastore"
	"figApi/util"
	"math/rand"
	"time"
)


type Company struct {
	CompanyName     string `json:"companyName"`
	OrgNum          string `json:"orgNum"`
	VatCode         string `json:"vatCode"`
	BeneficialOwner string `json:"beneficialOwner"`
}

func fetchCompanies(amount int) []Company {

	var companies []Company
		for i := 1; i <= amount; i++ {
			companies = append(companies, generateCompany())
		}

	return companies
}

func generateCompany() Company {
	seed := time.Now().UnixNano()

	company := Company{
		CompanyName:     getCompanyname(),
		OrgNum:          getFormattedOrgNum(seed),
		VatCode:         getVatNumForOrgNum(getOrgNum(seed)),
		BeneficialOwner: getFullName(),
	}

	return company
}


func getCompanyname() string {
	rand.Seed(time.Now().UnixNano())
	numOfWords := util.RangeIn(2, 5)
	var companyname strings.Builder

	for i := 0; i < numOfWords; i++ {
		companyname.WriteString(datastore.GetRandomLine("companynameparts"))
		companyname.WriteString(" ")
	}

	finishedCompanyname := strings.TrimSpace(companyname.String())

	return finishedCompanyname
}

func getVatNumForOrgNum(orgNum string) string {
	prefix := "SE"
	suffix := "01"

	parts := []string{prefix, orgNum, suffix}

	vatcode := strings.Join(parts, "")

	return vatcode
}
