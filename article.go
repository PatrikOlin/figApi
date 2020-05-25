package main

import (
	"strconv"
	"math/rand"
	"time"	 
	"figApi/datastore"
	"figApi/util"
)			 

type Article struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	VatCode     int `json:"vatCode"`
	Price       string `json:"price"`
	Unit        string `json:"unit"`
}

func fetchArticles(amount int) []Article {
	
	var articles []Article
		for i := 1; i <= amount; i++ {
			articles = append(articles, generateArticle())
		}

	return articles
}

func generateArticle() Article {
	rand.Seed(time.Now().UnixNano())

	article := Article{
		Id:          strconv.Itoa(util.RangeIn(1, 9999)),
		Description: getArticleName(),
		VatCode:     util.RangeIn(0, 3),
		Price:       strconv.Itoa(util.RangeIn(1, 99999)),
		Unit:        "st",
	}
				
	return article
}				

func getArticleName() string {
	article := datastore.GetRandomLine("articles")
	return article
}
