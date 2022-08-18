import { CircularProgress, Grid, Typography } from "@mui/material";
import React from "react";
import { CardLineChart } from "../components";
import { GeneralDiversification } from "../components/Cards/GeneralDiversification";
import { OverviewHeader } from "../components/Dashboard/OverviewHeader";
import { useNetworthOverview } from "../hooks/dashboard/useNetworth";

export const Dashboard = () => {
  const { liquidity, netWorth, loading, error, currentRecord, lastRecord } =
    useNetworthOverview();
  if (loading) {
    return <CircularProgress />;
  }
  if (error !== null) {
    return <Typography color="error">{error}</Typography>;
  }
  return (
    <>
      <Grid container spacing={3}>
        <Grid item xs={6}>
          <OverviewHeader
            currentRecord={currentRecord}
            lastRecord={lastRecord}
          />
        </Grid>
        <Grid item md={8} xs={12}>
          <CardLineChart
            data={netWorth}
            loading={loading}
            error={error}
            title="Net Worth"
          />
        </Grid>
        <Grid item md={4} xs={12}>
          <GeneralDiversification currentRecord={currentRecord} />
        </Grid>
        <Grid item xs={12}>
          <CardLineChart
            data={liquidity}
            loading={loading}
            error={error}
            title="Liquidity"
          />
        </Grid>
      </Grid>
    </>
  );
};
