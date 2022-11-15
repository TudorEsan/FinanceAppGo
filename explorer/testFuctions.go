package main

import (
	"fmt"

	"github.com/TudorEsan/FinanceAppGo/explorer/blockchains"
)

func main() {
	fmt.Println(blockchains.GetEthFromAddress("0x9C9d497FCF0566cF308516A5B8ed9C6991FAd049"))
}