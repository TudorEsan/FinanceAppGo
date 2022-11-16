import axios from "axios";
import dotenv from "dotenv";
dotenv.config();
const API_KEY = process.env.ETHERSCAN_API_KEY;

const contractAddresses = {
  link: "0x514910771AF9Ca656af840dff83E8264EcF986CA",
  omi: "0xeD35af169aF46a02eE13b9d79Eb57d6D68C1749e",
};

const WEI: BigInt = 1_000_000_000_000_000_000n;

interface ITokenBallance {
  status: string;
  message: string;
  result: string;
}

export const getErc20TokenHoldings = async (address: string, token: string): Promise<number> => {
  const contractAddress = contractAddresses[token.toLowerCase()];
  if (!contractAddress) {
    throw new Error("token not found / not supported");
  }
  const { data: resp } = await axios.get(
    `https://api.etherscan.io/api?module=account&action=tokenbalance&contractaddress=${contractAddress}&address=${address}&apikey=${API_KEY}`
  ) as { data: ITokenBallance };

  if (resp.status !== "1") {
    throw new Error(resp.message);
  }
  const amount = Number(resp.result);
  return amount / 10 ** 18;

};
