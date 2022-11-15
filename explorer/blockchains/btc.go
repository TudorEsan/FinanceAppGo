package blockchains

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/TudorEsan/FinanceAppGo/explorer/models"
)

func GetBtcFromAddress(address string) (models.AddressOverview, error) {
	explorerDomain := "https://explorer.btc.com/btc/address/" + address
	fmt.Print("explorerDomain: ", explorerDomain)
	
	site, err := http.Get(explorerDomain)
	if err != nil {
		return models.AddressOverview{}, err
	}
	defer site.Body.Close()

	doc, err := goquery.NewDocumentFromReader(site.Body)
	if err != nil {
		return models.AddressOverview{}, err
	}

	var btcAmount, usdAmount float64
	var e error
	doc.Find("span.jsx-3465105434.font-size-sm ").Each(func(i int, s *goquery.Selection) {
		if (i == 1) {
			// remove " BTC"
			balance := s.Text()
			fmt.Print("balance: ", balance)
			i := strings.Index(balance, "BTC")
			fmt.Println(i)			
			balance = balance[:i-1]

			btcAmount, err = strconv.ParseFloat(balance, 64)
			if err != nil {
				e = err
				return
			}
		}
		if (i == 2) {
			amount := s.Text()
			amount = strings.ReplaceAll(amount, "$", "")
			amount = strings.ReplaceAll(amount, " ", "")
			amount = strings.ReplaceAll(amount, ",", "")
			fmt.Println("amount: ", amount)
			usdAmount, err = strconv.ParseFloat(amount, 64)
			if err != nil {
				e = err
				return
			}
			
		}
		
	})

	if e != nil {
		fmt.Println("error: ", e)
		return models.AddressOverview{}, e
	}
	fmt.Println("btcAmount: ", btcAmount)
	fmt.Printf("usdAmount: %f \n", usdAmount)

	return models.AddressOverview{
		Blockchain: "BTC",
		Amount:     btcAmount,
		USDAmount:  usdAmount,
	}, nil
}
