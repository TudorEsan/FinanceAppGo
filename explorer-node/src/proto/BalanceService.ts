import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { BalanceServiceClient as _BalanceServiceClient, BalanceServiceDefinition as _BalanceServiceDefinition } from './BalanceService';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  Address: MessageTypeDefinition
  AddressOverview: MessageTypeDefinition
  BalanceService: SubtypeConstructor<typeof grpc.Client, _BalanceServiceClient> & { service: _BalanceServiceDefinition }
  Coins: MessageTypeDefinition
  WalletOverview: MessageTypeDefinition
}

ument: _Address, metadata: grpc.Metadata, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  GetAdaBalance(argument: _Address, options: grpc.CallOptions, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  GetAdaBalance(argument: _Address, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  getAdaBalance(argument: _Address, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  getAdaBalance(argument: _Address, metadata: grpc.Metadata, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  getAdaBalance(argument: _Address, options: grpc.CallOptions, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  getAdaBalance(argument: _Address, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  
  GetBtcBalance(argument: _Address, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  GetBtcBalance(argument: _Address, metadata: grpc.Metadata, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  GetBtcBalance(argument: _Address, options: grpc.CallOptions, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  GetBtcBalance(argument: _Address, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  getBtcBalance(argument: _Address, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  getBtcBalance(argument: _Address, metadata: grpc.Metadata, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  getBtcBalance(argument: _Address, options: grpc.CallOptions, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  getBtcBalance(argument: _Address, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  
  GetDotBalance(argument: _Address, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  GetDotBalance(argument: _Address, metadata: grpc.Metadata, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  GetDotBalance(argument: _Address, options: grpc.CallOptions, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  GetDotBalance(argument: _Address, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  getDotBalance(argument: _Address, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  getDotBalance(argument: _Address, metadata: grpc.Metadata, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  getDotBalance(argument: _Address, options: grpc.CallOptions, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  getDotBalance(argument: _Address, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  
  GetErc20Balance(argument: _Address, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_Coins__Output>): grpc.ClientUnaryCall;
  GetErc20Balance(argument: _Address, metadata: grpc.Metadata, callback: grpc.requestCallback<_Coins__Output>): grpc.ClientUnaryCall;
  GetErc20Balance(argument: _Address, options: grpc.CallOptions, callback: grpc.requestCallback<_Coins__Output>): grpc.ClientUnaryCall;
  GetErc20Balance(argument: _Address, callback: grpc.requestCallback<_Coins__Output>): grpc.ClientUnaryCall;
  getErc20Balance(argument: _Address, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_Coins__Output>): grpc.ClientUnaryCall;
  getErc20Balance(argument: _Address, metadata: grpc.Metadata, callback: grpc.requestCallback<_Coins__Output>): grpc.ClientUnaryCall;
  getErc20Balance(argument: _Address, options: grpc.CallOptions, callback: grpc.requestCallback<_Coins__Output>): grpc.ClientUnaryCall;
  getErc20Balance(argument: _Address, callback: grpc.requestCallback<_Coins__Output>): grpc.ClientUnaryCall;
  
  GetSolBalance(argument: _Address, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  GetSolBalance(argument: _Address, metadata: grpc.Metadata, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  GetSolBalance(argument: _Address, options: grpc.CallOptions, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  GetSolBalance(argument: _Address, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  getSolBalance(argument: _Address, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  getSolBalance(argument: _Address, metadata: grpc.Metadata, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  getSolBalance(argument: _Address, options: grpc.CallOptions, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  getSolBalance(argument: _Address, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  
}

export interface BalanceServiceHandlers extends grpc.UntypedServiceImplementation {
  GetAdaBalance: grpc.handleUnaryCall<_Address__Output, _AddressOverview>;
  
  GetBtcBalance: grpc.handleUnaryCall<_Address__Output, _AddressOverview>;
  
  GetDotBalance: grpc.handleUnaryCall<_Address__Output, _AddressOverview>;
  
  GetErc20Balance: grpc.handleUnaryCall<_Address__Output, _Coins>;
  
  GetSolBalance: grpc.handleUnaryCall<_Address__Output, _AddressOverview>;
  
}

export interface BalanceServiceDefinition extends grpc.ServiceDefinition {
  GetAdaBalance: MethodDefinition<_Address, _AddressOverview, _Address__Output, _AddressOverview__Output>
  GetBtcBalance: MethodDefinition<_Address, _AddressOverview, _Address__Output, _AddressOverview__Output>
  GetDotBalance: MethodDefinition<_Address, _AddressOverview, _Address__Output, _AddressOverview__Output>
  GetErc20Balance: MethodDefinition<_Address, _Coins, _Address__Output, _Coins__Output>
  GetSolBalance: MethodDefinition<_Address, _AddressOverview, _Address__Output, _AddressOverview__Output>
}
