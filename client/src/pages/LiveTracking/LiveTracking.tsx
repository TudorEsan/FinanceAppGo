import { Button, Typography } from "@mui/material";
import { Box } from "@mui/system";
import React from "react";
import { useNavigate } from "react-router-dom";

export const NoKeys = () => {
  const navigate = useNavigate();
  return (
    <Box>
      <Typography variant="h5" gutterBottom>
        Live Tracking
      </Typography>
      <Typography maxWidth="500px">
        Here you can track your investments in real time! Add your binance api
        keys or crypto wallet addresses to get real time statistics and value of
        your portofolio
      </Typography>
      <Button
        variant="outlined"
        size="small"
        sx={{ mt: 2 }}
        onClick={() => navigate("/live-tracking-settings")}
      >
        Settings
      </Button>
    </Box>
  );
};

export const LiveTracking = () => {
  return <NoKeys />;
};
