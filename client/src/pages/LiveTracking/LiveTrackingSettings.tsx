import {
  Box,
  Card,
  CardContent,
  Divider,
  Paper,
  Typography,
} from "@mui/material";
import React from "react";
import { MyCard } from "../../components";

export const LiveTrackingSettings = () => {
  return (
    <Box>
      <Paper sx={{ p: 2, minHeight: "75vh" }} variant="outlined">
        <Typography variant="h5" gutterBottom>
          Live Tracking Settings
        </Typography>
        <Typography maxWidth="600px">
          Manage your live tracking settings here. Add your binance api keys or
          crypto wallet addresses to get real time statistics and value of your
          portofolio
        </Typography>
        <Divider sx={{ my: 2 }} />
        <MyCard>
          <Box display="flex" gap={2} alignItems="center">
            <img src="/binanceLogo.svg" width="30px" height="30px" />
            <Typography variant="h6">Binance API Keys</Typography>
          </Box>
        </MyCard>
      </Paper>
    </Box>
  );
};
