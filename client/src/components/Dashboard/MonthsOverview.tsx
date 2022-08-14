import React from "react";
import { useNetworthOverview } from "../../hooks/dashboard/useNetworth";
import { Datum, ResponsiveLine } from "@nivo/line";
import { MyCard } from "../Cards/MyCard";
import { LineChart } from "../Charts/LineChart";
import { Typography } from "@mui/material";
import { INetWorth } from "../../types/overview";

interface IProps {
  data: Datum[];
  loading: boolean;
  error: string | null;
  title: string;
}
export const CardLineChart = ({ data, title, loading, error }: IProps) => {
  if (loading) return <div>Loading...</div>;
  if (error) return <Typography color="error">{error}</Typography>;
  return (
    <>
      <MyCard>
        <Typography variant="h5" textAlign="center">
          {title}
        </Typography>
        <LineChart id="months-overview" data={data} />
      </MyCard>
    </>
  );
};
