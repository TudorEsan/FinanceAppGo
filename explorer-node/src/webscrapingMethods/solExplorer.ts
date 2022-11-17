import  Puppeteer  from "puppeteer";
import { AddressExplorerResponse } from "../types/explorerTypes";

export const getSolBalance = async (address: string): Promise<AddressExplorerResponse> => {
  const browser = await Puppeteer.launch({ headless: false, waitForInitialPage: true });
  const page = await browser.newPage();
  console.log("page: ", 'https://explorer.solana.com/address/' + address);
  await page.goto(`https://solscan.io/account/${address}`, {
    waitUntil: 'networkidle0'
  });
  page.screenshot({ path: "solscan.png" });
  const balanceText = await page.evaluate(() => {
    return document.querySelector("div.ant-card-body div.ant-col.ant-col-24.ant-col-md-16").textContent;
  });
  const [solBalanceStr, usdBalanceStr] = balanceText.split(" SOL ")
  const solBalance = Number(solBalanceStr.replace(/,/g, ""));
  const usdBalance = Number(usdBalanceStr.replace(/,/g, "").replace("$", "").replace("(", "").replace(")", "").trim());
  return {
    token: 'SOL',
    balance: solBalance,
    usdBalance: usdBalance
  }
};