import Puppeteer from "puppeteer";
import { AddressExplorerResponse } from "../types/explorerTypes";

export const getBtcBalance = async (
  address: string
): Promise<AddressExplorerResponse> => {
  const browser = await Puppeteer.launch({ headless: true });
  const page = await browser.newPage();
  await page.goto(`https://www.blockchain.com/btc/address/${address}`);
  const balance = await page.evaluate(() => {
    const balanceText = document.querySelector(
      "span.sc-1ryi78w-0.cILyoi.sc-16b9dsl-1.ZwupP.u3ufsr-0.eQTRKC"
    );
    // ballance text example:
    /**
     * This address has transacted 2 times on the Bitcoin blockchain.
     * It has received a total of 0.99872274 BTC ($16,716.21) and has sent a total of 0.99872274 BTC ($16,716.21).
     * The current value of this address is 0.00000000 BTC ($0.00).
     */
    let i = balanceText.textContent.indexOf("a total of ");
    console.log(i);
    let j = balanceText.textContent.indexOf(" BTC");
    console.log(j);

    const btcBalance = Number(balanceText.textContent.slice(i + 10, j).trim());
    i = balanceText.textContent.indexOf("(");
    j = balanceText.textContent.indexOf(")");
    let usdBalance = balanceText.textContent.slice(i + 2, j).trim(); // exclude the $ sign too
    // replace all commas with empty string
    const usdBallanceNumber = Number(usdBalance.replace(/,/g, ""));
    console.log(usdBalance);
    return { balance: btcBalance, usdBalance: usdBallanceNumber };
  });
  return {
    token: "BTC",
    balance: balance.balance,
    usdBalance: balance.usdBalance,
  };
};
