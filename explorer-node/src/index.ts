import puppeteer from "puppeteer";
import { getErc20TokenHoldings } from "./webscrapingMethods/erc20Holdings";
import dontenv from "dotenv";
import { getBtcBalance } from "./webscrapingMethods/btcExplorer";
import { getSolBalance } from "./webscrapingMethods/solExplorer";
import { getAdaBalance } from "./webscrapingMethods/adaExplorer";
import { getDotBalance } from "./webscrapingMethods/dotExplorer";
import { getBalanceServer } from "./service/balanceService";
import * as grpc from "@grpc/grpc-js";

dontenv.config();

const main = async () => {
  const balanceServer = getBalanceServer();
  balanceServer.bindAsync(
    "localhost:8083",
    grpc.ServerCredentials.createInsecure(),
    (err) => {
      if (err) {
        console.error(err);
        return
      }
      console.log("Server running at http://localhost:8083");
      balanceServer.start();
      
    }
  );
};

main();
