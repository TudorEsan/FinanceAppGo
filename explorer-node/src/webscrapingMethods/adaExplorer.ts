import Puppeteer from "puppeteer";
import { AddressExplorerResponse } from "../types/explorerTypes";

export const getAdaBalance = async (
  address: string
): Promise<AddressExplorerResponse> => {
  const browser = await Puppeteer.launch({
    headless: true,
    waitForInitialPage: true,
  });
  const page = await browser.newPage();
  await page.goto(`https://cardanoscan.io/address/${address}`);
  const balanceText = await page.evaluate(() => {
    const balanceArea = document.querySelector("div.flex.flex-wrap.gap-x-10.gap-y-5")
      
    const balance = balanceArea.querySelector("div.flex.items-baseline.font-mono")
    const usdBalance = balanceArea.querySelectorAll("span.text-base.font-semibold.textColor1")[1]
    return {
      balance: balance.textContent,
      usdBalance: usdBalance.textContent
    }
  });
  const usdBalance = Number(balanceText.usdBalance.replace(/,/g, "").replace("$", "").trim());
  const adaBalance = Number(balanceText.balance.replace(/,/g, "").trim());
  return {
    token: "ADA",
    balance: adaBalance,
    usdBalance,
  };
};
