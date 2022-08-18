import { IRecord } from "../types/record";

const roundPercent = (value: number, decimals: number) => {
  return (
    Math.floor(value * Math.pow(10, decimals)) / Math.pow(10, decimals - 2)
  );
};

export const getTotalNetWorth = (record?: IRecord | null) => {
  if (!record) {
    return 0;
  }
  return record.liquidity + record.investedAmount;
};

export const getDiversification = (record: IRecord) => {
  const diversification = [];
  const total = record.stocksValue + record.cryptosValue + record.liquidity;
  diversification.push({
    symbol: "Liquidity",
    percent: roundPercent(record.liquidity / total, 2),
  });
  diversification.push({
    symbol: "Stocks",
    percent: roundPercent(record.stocksValue / total, 2),
  });
  diversification.push({
    symbol: "Cryptos",
    percent: roundPercent(record.cryptosValue / total, 2),
  });
  return diversification;
};
