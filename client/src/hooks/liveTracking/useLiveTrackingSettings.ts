
import React from 'react'
import { IApiKeys, IWalletAddress } from '../../types/liveTracking';

export const useLiveTrackingSettings = () => {
  const [binanceKeys, setBinanceKeys] = React.useState<IApiKeys | null>(null)
  const [walletAddresses, setWalletAddresses] = React.useState<IWalletAddress[]>([]);

  const [dataLoading, setDataLoading] = React.useState<boolean>(false);
  const [actionLoading, setActionLoading] = React.useState<boolean>(false); // for updates / deletes
  const [error, setError] = React.useState<string | null>(null);
  
  const addWalletAddress = async (address: IWalletAddress) => {
    setWalletAddresses([...walletAddresses, address]);
  }

  const removeWalletAddress = async (address: IWalletAddress) => {
    setWalletAddresses(walletAddresses.filter((walletAddress) => walletAddress.address !== address.address));
  }

  const addBinanceKeys = async (keys: IApiKeys) => {
    setBinanceKeys(keys);
  }

  const removeBinanceKeys = async () => {
    setBinanceKeys(null);
  }



  return {
      binanceKeys, dataLoading, actionLoading, walletAddresses, addWalletAddress, removeWalletAddress, addBinanceKeys, removeBinanceKeys, error
  }
}
