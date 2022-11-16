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
  WalletOverview: MessageTypeDefinition
}

