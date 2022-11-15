package blockchains

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/TudorEsan/FinanceAppGo/explorer/models"
)

func GetErc20TokenHoldings(address string) ([]models.AddressOverview, error) {
	explorerLink := "https://etherscan.io/tokenholdings?a=" + address
	fmt.Println("explorerLink: ", explorerLink)

	res, err := http.Get(explorerLink)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	fmt.Println("res: ", res.Body)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(doc.Text())
	doc.Find("tbody").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})

	return nil, nil
}
