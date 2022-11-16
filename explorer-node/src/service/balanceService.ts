import * as grpc from "@grpc/grpc-js";
import * as protoLoader from '@grpc/proto-loader';
import { BalanceServiceHandlers } from "../proto/BalanceService";
import { ProtoGrpcType } from "../proto/explorer";
const PROTO_PATH = "src/proto/balanceService.proto";

const options = {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
};

const balanceServer: BalanceServiceHandlers = {
  GetAdaBalance: (call, callback) => {
    console.log("GetAdaBalance: ", call.request);
    callback(null, { Balance: 0, UsdBallance: 0, Token: "ADA" });
  },
  GetBtcBalance: (call, callback) => {
    console.log("GetBtcBalance: ", call.request);
    callback(null, { Balance: 0, UsdBallance: 0, Token: "BTC" });
  },
  GetDotBalance: (call, callback) => {
    console.log("GetDotBalance: ", call.request);
    callback(null, { Balance: 0, UsdBallance: 0, Token: "DOT" });
  },
  GetSolBalance: (call, callback) => {
    console.log("GetSolBalance: ", call.request);
    callback(null, { Balance: 0, UsdBallance: 0, Token: "SOL" });
  },
};

export function getBalanceServer(): grpc.Server {
  const packageDefinition = protoLoader.loadSync(PROTO_PATH);
  const proto = (grpc.loadPackageDefinition(
    packageDefinition
  ) as unknown) as ProtoGrpcType;
  const server = new grpc.Server();
  server.addService(proto.BalanceService.service, balanceServer);
  return server;
}

