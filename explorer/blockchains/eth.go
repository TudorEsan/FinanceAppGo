package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/TudorEsan/FinanceAppGo/explorer/models"
)




func getEthFromAddress(address string) (models.AddressOverview, error) {
	explorerDomain := "https://etherscan.io/address/"

	res, err := http.Get(explorerDomain + address)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	card := doc.Find("div.card-body").First()
	var etherBallance, usdBallance float64
	var e error
	card.Find("div.row.align-items-center").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			etherBallance, err = strconv.ParseFloat(strings.Replace(s.Find("div.col-md-8").Text(), " Ether", "", -1), 64)
			if err != nil {
				e = err
				return
			}

		} else if i == 1 {
			strUsdBallance := strings.Replace(s.Find("div.col-md-8").Text(), " USD", "", -1)
			bracketIndex := strings.Index(strUsdBallance, "(")
			strUsdBallance = strUsdBallance[:bracketIndex-2] // remove " ("
			strUsdBallance = strings.Replace(strUsdBallance, "$", "", -1)
			strUsdBallance = strings.Replace(strUsdBallance, ",", "", -1)
			usdBallance, err = strconv.ParseFloat(strUsdBallance, 64)
			if err != nil {
				e = err
			}
		}

	})

	if e != nil {
		return models.AddressOverview{}, e
	}

	return models.AddressOverview{
		Blockchain: "ETH",
		Amount:     etherBallance,
		USDAmount:  usdBallance,
	}, nil

}
