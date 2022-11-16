import Puppeteer from "puppeteer";
import { AddressExplorerResponse } from "../types/explorerTypes";

export const getDotBalance = async (
  address: string
): Promise<AddressExplorerResponse> => {
  const browser = await Puppeteer.launch({
    headless: true,
    waitForInitialPage: true,
  });
  const page = await browser.newPage();
  await page.goto(`https://polkadot.subscan.io/account/${address}`, {
  waitUntil: "networkidle0",
  });
  page.screenshot({ path: "polkadot.png" });
  const balance = await page.evaluate(() => {
    const balanceArea = document.querySelector(
      "div.balance-wrapper"
    );

    const balance = balanceArea.querySelector(
      "div.balance-text"
    );
    const usdBalance = balanceArea.querySelector(
      "div.token-value-text"
    );
    return {
      balance: balance.textContent,
      usdBalance: usdBalance.textContent,
    };
  });
  console.log("balance: ", balance);
  return {
    token: "DOT",
    balance: Number(balance.balance.replace(/,/g, "").trim()),
    usdBalance: Number(balance.usdBalance.replace(/,/g, "").replace("$", "").replace("≈", "").trim()),
  }
};
