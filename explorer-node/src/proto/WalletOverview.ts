// Original file: src/proto/explorer.proto

import type { AddressOverview as _AddressOverview, AddressOverview__Output as _AddressOverview__Output } from './AddressOverview';

export interface WalletOverview {
  'Holdings'?: (_AddressOverview)[];
}

export interface WalletOverview__Output {
  'Holdings': (_AddressOverview__Output)[];
}
