import puppeteer from "puppeteer";
import { getErc20TokenHoldings } from "./webscrapingMethods/erc20Holdings";
import dontenv from "dotenv";
import { getBtcBalance } from "./webscrapingMethods/btcExplorer";
import { getSolBalance } from "./webscrapingMethods/solExplorer";
import { getAdaBalance } from "./webscrapingMethods/adaExplorer";
dontenv.config();

const main = async () => {
  try {

    console.log(await getAdaBalance("addr1qxyn28janhl5xq5xxn6n29pktsjqw74qljfuwd0nmgc827q9u8l8c3mttswgmh4c29g20dyv38zf73r7ccp3p9j94pfqg7rjzv"))
  } catch(e) {
    console.log("error: ", e)
  }
};

main();
