import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";
import { BalanceServiceHandlers } from "../proto/BalanceService";
import { ProtoGrpcType } from "../proto/explorer";
import wrapServerWithReflection from "grpc-node-server-reflection";
import { getAdaBalance } from "../webscrapingMethods/adaExplorer";
import { getBtcBalance } from "../webscrapingMethods/btcExplorer";
import { getDotBalance } from "../webscrapingMethods/dotExplorer";
import { getSolBalance } from "../webscrapingMethods/solExplorer";
import { getErc20Holdings } from "../webscrapingMethods/erc20Explorer";
import { ICoinsObject } from "../types/explorerTypes";
import { Coins } from "../proto/Coins";

const PROTO_PATH = "src/proto/balanceService.proto";

const options = {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
};

const balanceServer: BalanceServiceHandlers = {
  GetAdaBalance: async (call, callback) => {
    console.log("GetAdaBalance: ", call.request);
    const addresses = call.request.Addresses;
    const balance = {
      token: "ADA",
      balance: 0,
      usdBalance: 0,
    };
    for (let address of addresses) {
      try {
        const res = await getAdaBalance(address);
        balance.balance += res.balance;
        balance.usdBalance += res.usdBalance;
      } catch (e) {
        console.error(e);
      }
    }
    callback(null, {
      Balance: balance.balance,
      UsdBallance: balance.usdBalance,
      Token: "ADA",
    });
  },
  GetBtcBalance: async (call, callback) => {
    console.log("GetBtcBalance: ", call.request);
    const addresses = call.request.Addresses;
    const balance = {
      token: "BTC",
      balance: 0,
      usdBalance: 0,
    };
    for (let address of addresses) {
      try {
        const res = await getBtcBalance(address);
        console.log('balance', balance)
        balance.balance += res.balance;
        balance.usdBalance += res.usdBalance;
      } catch (e) {
        console.error(e);
      }
    }
    callback(null, {
      Balance: balance.balance,
      UsdBallance: balance.usdBalance,
      Token: "BTC",
    });
  },
  GetDotBalance: async (call, callback) => {
    console.log("GetDotBalance: ", call.request);
    const addresses = call.request.Addresses;
    const balance = {
      token: "DOT",
      balance: 0,
      usdBalance: 0,
    };
    for (let address of addresses) {
      try {
        const res = await getDotBalance(address);
        balance.balance += res.balance;
        balance.usdBalance += res.usdBalance;
      } catch (e) {
        console.error(e);
      }
    }

    callback(null, {
      Balance: balance.balance,
      UsdBallance: balance.usdBalance,
      Token: "DOT",
    });
  },
  GetSolBalance: async (call, callback) => {
    console.log("GetSolBalance: ", call.request);
    const addresses = call.request.Addresses;
    const balance = {
      token: "SOL",
      balance: 0,
      usdBalance: 0,
    };
    for (let address of addresses) {
      try {
        const res = await getSolBalance(address);
        balance.balance += res.balance;
        balance.usdBalance += res.usdBalance;
      } catch (e) {
        console.error(e);
      }
    }
    callback(null, {
      Balance: balance.balance,
      UsdBallance: balance.usdBalance,
      Token: "SOL",
    });
  },
  GetErc20Balance: async (call, callback) => {
    console.log("GetEtc20Balance: ", call.request);
    const addresses = call.request.Addresses;
    const coins: ICoinsObject = {};
    for (let address of addresses) {
      try {
        const res = await getErc20Holdings(address);
        for (let coin of res) {
          if (coins[coin.token]) {
            coins[coin.token].balance += coin.balance;
            coins[coin.token].usdBalance += coin.usdBalance;
          } else {
            coins[coin.token] = coin;
          }
        }
      } catch (e) {
        console.error(e);
      }
    }
    // convert object to array
    const coinsArray: Coins = {
      coins: Object.keys(coins).map((key) => {
        return {
          Balance: coins[key].balance,
          UsdBallance: coins[key].usdBalance,
          Token: coins[key].token,
        };
      })
    }
    callback(null, coinsArray);
  }
};

export function getBalanceServer(): grpc.Server {
  const packageDefinition = protoLoader.loadSync(PROTO_PATH);
  const proto = grpc.loadPackageDefinition(
    packageDefinition
  ) as unknown as ProtoGrpcType;
  const server = wrapServerWithReflection(new grpc.Server());
  server.addService(proto.BalanceService.service, balanceServer);
  return server;
}
