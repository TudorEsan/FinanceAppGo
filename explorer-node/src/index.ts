import puppeteer from "puppeteer";
import { getErc20TokenHoldings } from "./webscrapingMethods/erc20Holdings";
import dontenv from "dotenv";
import { getBtcBalance } from "./webscrapingMethods/btcExplorer";
import { getSolBalance } from "./webscrapingMethods/solExplorer";
import { getAdaBalance } from "./webscrapingMethods/adaExplorer";
import { getDotBalance } from "./webscrapingMethods/dotExplorer";
dontenv.config();

const main = async () => {
  try {

    console.log(await getDotBalance("1jZqirgn6ECrqwYNGUT1No9QSBWTatCkCe2nRzxFw2ufbyN"))
  } catch(e) {
    console.log("error: ", e)
  }
};

main();
