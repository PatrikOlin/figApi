package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"							   
)									   

func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	amount, err := getAmountParam(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error:" "invalid datatype for parameter"}`))
		return
	}

	data := fetchArticles(amount)	   

	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error:" "error marshalling data"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	amount, err := getAmountParam(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error:" "invalid datatype for parameter"}`))
		return
	}

	data := fetchPeople(amount)	   	   

	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error:" "error marshalling data"}`))
		return						   
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func getCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	amount, err := getAmountParam(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error:" "invalid datatype for parameter"}`))
		return
	}

	data := fetchCompanies(amount)	   	   

	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error:" "error marshalling data"}`))
		return						   
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func getAmountParam(r *http.Request) (int, error) {
	amount := 1
	queryParams := r.URL.Query()
	fmt.Println(queryParams)
	a := queryParams.Get("amount")
	if a != "" {
		val, err := strconv.Atoi(a)
		if err != nil {
			return amount, err
		}
		amount = val
	}								   
	return amount, nil				   
}


									   
