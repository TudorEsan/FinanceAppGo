package main

import (
	"fmt"

	"github.com/TudorEsan/FinanceAppGo/explorer/blockchains"
)


func main() {
	fmt.Println(blockchains.GetBtcFromAddress("1AC4fMwgY8j9onSbXEWeH6Zan8QGMSdmtA"))
}