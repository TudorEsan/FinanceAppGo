import puppeteer from "puppeteer";
import { getErc20TokenHoldings } from "./webscrapingMethods/erc20Holdings";
import dontenv from "dotenv";
import { getBtcBalance } from "./webscrapingMethods/btcExplorer";
dontenv.config();

const main = async () => {
  try {

    // console.log(await getErc20TokenHoldings("0x9C9d497FCF0566cF308516A5B8ed9C6991FAd049", "link"));
    // console.log(await getErc20TokenHoldings("0x9C9d497FCF0566cF308516A5B8ed9C6991FAd049", "omi"));
    getBtcBalance("bc1qmgp6k2cjjnpp7922nahfda7ufxlwmwlxmu5wva")
  } catch(e) {
    console.log("error: ", e)
  }
};

main();
