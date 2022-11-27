import Puppeteer from "puppeteer";
import { AddressExplorerResponse } from "../types/explorerTypes";

export const getErc20Holdings = async (
  address: string
): Promise<AddressExplorerResponse[]> => {
  const browser = await Puppeteer.launch({
    headless: true,
    args: ["--window-size=2160,3840"],
    waitForInitialPage: true,
  });
  const page = await browser.newPage();
  page.setViewport({ width: 3840, height: 2160 });
  await page.goto(`https://ethplorer.io/address/${address}`);


  const balances = await page.evaluate(() => {
    // get id address-token-balances
    const balances = document.querySelector("#address-token-balances");

    // iterate over all tr elements
    // get the token name
    // get the token balance
    const trs = balances.querySelectorAll("tr");
    const balancesArray = [];
    for (let i = 0; i < trs.length; i++) {
      const coin = trs[i].querySelector("td:nth-child(1)").textContent;
      const balanceArea = trs[i].querySelector("td:nth-child(2)");
      const splited = balanceArea.textContent.split(/\s+/);
      const balance = splited[0];
      const price = splited[2];

      balancesArray.push({
        token: coin,
        balance: Number(balance.replace(/,/g, "")),
        usdBalance: Number(price.replace(/,/g, "")),
      });
    }

    return balancesArray;
  });
  return balances;
};
