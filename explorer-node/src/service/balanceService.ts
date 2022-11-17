import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";
import { BalanceServiceHandlers } from "../proto/BalanceService";
import { ProtoGrpcType } from "../proto/explorer";
import wrapServerWithReflection from "grpc-node-server-reflection";
import { getAdaBalance } from "../webscrapingMethods/adaExplorer";
import { getBtcBalance } from "../webscrapingMethods/btcExplorer";
import { getDotBalance } from "../webscrapingMethods/dotExplorer";
import { getSolBalance } from "../webscrapingMethods/solExplorer";

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
        const balance = await getAdaBalance(address);
        balance.balance += balance.balance;
        balance.usdBalance += balance.usdBalance;
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
        const balance = await getBtcBalance(address);
        balance.balance += balance.balance;
        balance.usdBalance += balance.usdBalance;
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
        const balance = await getDotBalance(address);
        balance.balance += balance.balance;
        balance.usdBalance += balance.usdBalance;
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
        const balance = await getSolBalance(address);
        balance.balance += balance.balance;
        balance.usdBalance += balance.usdBalance;
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
