package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://www.gundamhobby.ca/")

	if err != nil {
		return
	}

	if err != nil {
		return
	}

	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return
	}

	doc.Find("main>.grid>div.grid__item>div>div.grid__item").Each(func(i int, s *goquery.Selection) {
		soldOut := s.HasClass("sold-out")
		productName := s.Find("a>p.grid-link__title").Text()
		log.Println(productName, soldOut)
	})
}
