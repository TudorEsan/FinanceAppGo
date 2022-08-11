import React from "react";
import { useNetworthOverview } from "../../hooks/dashboard/useNetworth";
import { ResponsiveLine } from "@nivo/line";
import { MyCard } from "../Cards/MyCard";
import { LineChart } from "../Charts/LineChart";
import { Typography } from "@mui/material";

export const RecordOverview = () => {
  const { netWorth, loading, error, setYear } = useNetworthOverview();
  if (loading) return <div>Loading...</div>;
  if (error) return <Typography color="error">{error}</Typography>;
  return (
    <>
      <MyCard>
        <LineChart id="months-overview" data={netWorth} />
      </MyCard>
    </>
  );
};
