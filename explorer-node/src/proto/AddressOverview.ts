// Original file: src/proto/balanceService.proto


export interface AddressOverview {
  'Token'?: (string);
  'Balance'?: (number | string);
  'UsdBallance'?: (number | string);
}

export interface AddressOverview__Output {
  'Token': (string);
  'Balance': (number);
  'UsdBallance': (number);
}
