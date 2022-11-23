import { ExpandMore } from "@mui/icons-material";
import { LoadingButton } from "@mui/lab";
import {
  Accordion,
  AccordionDetails,
  AccordionSummary,
  Box,
  Button,
  Card,
  CardContent,
  Divider,
  Paper,
  Typography,
} from "@mui/material";
import React from "react";
import { useForm } from "react-hook-form";
import { ControlledTextField, MyCard } from "../../components";
import { useLiveTrackingSettings } from "../../hooks/liveTracking/useLiveTrackingSettings";
import { IApiKeys } from "../../types/liveTracking";

export const LiveTrackingSettings = () => {
  return (
    <Box>
      <Paper sx={{ p: 2, minHeight: "80vh" }} variant="outlined">
        <Typography variant="h5" gutterBottom>
          Live Tracking Settings
        </Typography>
        <Typography maxWidth="600px">
          Manage your live tracking settings here. Add your binance api keys or
          crypto wallet addresses to get real time statistics and value of your
          portofolio
        </Typography>
        <Divider sx={{ my: 2 }} />
        <SettingsAccordion />
      </Paper>
    </Box>
  );
};

const SettingsAccordion = () => {
  return (
    <Box>
      <Accordion>
        <AccordionSummary expandIcon={<ExpandMore />}>
          <Box display="flex" gap={2} alignItems="center">
            <img src="/binanceLogo.svg" width="30px" height="30px" />
            <Typography>Binance API Keys</Typography>
          </Box>
        </AccordionSummary>
        <AccordionDetails>
          <BinanceSettings />
        </AccordionDetails>
      </Accordion>
      <Accordion>
        <AccordionSummary expandIcon={<ExpandMore />}>
          <Box display="flex" gap={2} alignItems="center">
            <img
              src="/wallet.svg"
              style={{ color: "white" }}
              width="30px"
              height="30px"
            />
            <Typography>Wallet Addresses</Typography>
          </Box>
        </AccordionSummary>
        <AccordionDetails>
          <BinanceSettings />
        </AccordionDetails>
      </Accordion>
    </Box>
  );
};

const BinanceSettings = () => {
  const { handleSubmit, reset, control } = useForm<IApiKeys>();
  const { binanceKeys, dataLoading, actionLoading, addBinanceKeys, error } =
    useLiveTrackingSettings();

  const onSubmit = (data: IApiKeys) => {
    addBinanceKeys(data);
  };

  if (dataLoading) {
    return <Typography>Loading...</Typography>;
  }

  return (
    <Box display="flex" gap={1}>
      <form
        style={{
          display: "flex",
          gap: "10px",
          flexDirection: "column",
          width: "100%",
          justifyContent: "center",
        }}
        onSubmit={handleSubmit(onSubmit)}
      >
        <ControlledTextField
          control={control}
          name="apiKey"
          label="API Key"
          defaultValue={binanceKeys?.apiKey}
        />
        <ControlledTextField
          control={control}
          name="secretKey"
          label="Secret Key"
          defaultValue={binanceKeys?.apiSecret}
        />
        <LoadingButton
          type="submit"
          variant="contained"
          loading={actionLoading}
          sx={{ maxWidth: "150px" }}
        >
          Save
        </LoadingButton>
      </form>
      {error && (
        <Typography color="error" variant="body2">
          {error}
        </Typography>
      )}
    </Box>
  );
};

const WalletAddresses = () => {
  const {walletAddresses} = useLiveTrackingSettings();

  return (
    
  )
};