import puppeteer from "puppeteer";
import { getErc20TokenHoldings } from "./webscrapingMethods/erc20Holdings";
import dontenv from "dotenv";
import { getBtcBalance } from "./webscrapingMethods/btcExplorer";
import { getSolBalance } from "./webscrapingMethods/solExplorer";
dontenv.config();

const main = async () => {
  try {


  } catch(e) {
    console.log("error: ", e)
  }
};

main();
