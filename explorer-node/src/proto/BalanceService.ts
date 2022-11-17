// Original file: src/proto/balanceService.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { Address as _Address, Address__Output as _Address__Output } from './Address';
import type { AddressOverview as _AddressOverview, AddressOverview__Output as _AddressOverview__Output } from './AddressOverview';

export interface BalanceServiceClient extends grpc.Client {
  GetAdaBalance(argument: _Address, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
  GetAdaBalance(argument: _Address, metadata: grpc.Metadata, callback: grpc.requestCallback<_AddressOverview__Output>): grpc.ClientUnaryCall;
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
  
  GetSolBalance: grpc.handleUnaryCall<_Address__Output, _AddressOverview>;
  
}

export interface BalanceServiceDefinition extends grpc.ServiceDefinition {
  GetAdaBalance: MethodDefinition<_Address, _AddressOverview, _Address__Output, _AddressOverview__Output>
  GetBtcBalance: MethodDefinition<_Address, _AddressOverview, _Address__Output, _AddressOverview__Output>
  GetDotBalance: MethodDefinition<_Address, _AddressOverview, _Address__Output, _AddressOverview__Output>
  GetSolBalance: MethodDefinition<_Address, _AddressOverview, _Address__Output, _AddressOverview__Output>
}
