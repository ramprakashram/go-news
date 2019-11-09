package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/martinlindhe/notify"
)

type ArticleSource struct {
	Id   string
	Name string
}

type Articles struct {
	Source      ArticleSource
	Author      string
	Title       string
	Description string
	Url         string
	UrlToImage  string
	PublishedAt string
	Content     string
}

// Response Structure
type NewsResponse struct {
	Status       string
	TotalResults string
	Articles     []Articles
}

const APP_ID = "364aa37e19b347d39d795a42f8eb7dab"
const API = "https://newsapi.org/v2/top-headlines"

func getHeadlines(countryCode string, category string, pageSize string) []Articles {
	var newsData NewsResponse
	apiCall := API + "?country=" + countryCode + "&category=" + category + "&pageSize=" + pageSize + "&apiKey=" + APP_ID
	res, err := http.Get(apiCall)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	respData, respErr := ioutil.ReadAll(res.Body)

	if respErr != nil {
		fmt.Fprintf(os.Stderr, "%v\n", respErr)
	}
	// fmt.Printf("%s", respData)
	json.Unmarshal([]byte(respData), &newsData)
	return newsData.Articles
}

func main() {
	articles := getHeadlines("in", "technology", "3")
	if len(articles) > 0 {
		news := "\n" + articles[0].Title
		// news = news + "\n2. " + articles[1].Title
		// news = news + "\n3. " + articles[2].Title
		notify.Notify("Headlines", "Top Technology News", news, "")
	} else {
		fmt.Printf("%v", articles)
	}

}
