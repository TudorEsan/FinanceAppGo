import { Typography } from "@mui/material";
import React from "react";
import { CardLineChart } from "../components";
import { GeneralDiversification } from "../components/Cards/GeneralDiversification";
import { useNetworthOverview } from "../hooks/dashboard/useNetworth";

export const Dashboard = () => {
  const {liquidity, netWorth, loading, error } = useNetworthOverview();
  return (
    <>
      <CardLineChart data={netWorth} loading={loading} error={error} title="Net Worth" />
      <CardLineChart data={liquidity} loading={loading} error={error} title="Liquidity" />
    </>
  );
};
