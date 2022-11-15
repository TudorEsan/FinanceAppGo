package blockchains

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/TudorEsan/FinanceAppGo/explorer/models"
)

func GetSolFromAddress(address string) (models.AddressOverview, error) {
	explorerDomain := "https://solanabeach.io/address/" + address
	fmt.Println("explorerDomain: ", explorerDomain)

	site, err := http.Get(explorerDomain)
	if err != nil {
		return models.AddressOverview{}, err
	}
	defer site.Body.Close()

	doc, err := goquery.NewDocumentFromReader(site.Body)
	if err != nil {
		return models.AddressOverview{}, err
	}

	
	// var solAmount, usdAmount float64
	// var e error
fmt.Println(doc.Html())


	// details := doc.Find("/html/body/div/section/main/div/div[2]/div/div[1]/div/div[2]/div[1]/div[2]").Text()

	return models.AddressOverview{}, nil

}
