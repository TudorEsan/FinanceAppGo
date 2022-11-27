
export interface AddressExplorerResponse {
  token: string;
  balance: number;
  usdBalance: number;
}

export interface ICoinsObject {
  [key: string]: AddressExplorerResponse;
}

