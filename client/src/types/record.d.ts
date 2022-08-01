export interface IRecord {
  id: string;
  date: Date;
  liquidity: number;
  stocks: IStock[];
  cryptos: ICrypto[];
  cryptoDiversification: IDiversification[];
  stockDiversification: IDiversification[];
  cryptosValue: number;
  stocksValue: number;
  investedAmount: number;
}

export interface IDiversification {
  symbol: string;
  percent: number;
}

export interface IRecordForm {
  date: Date | string;
  liquidity: number;
  stocks: IStock[];
  cryptos: ICrypto[];
}

export interface IStock {
  name: string;
  symbol: string;
  shares: number;
  valuedAt: number;
}

export interface ICrypto {
  name: string;
  symbol: string;
  coins: number;
  valuedAt: number;
}
