import { Typography } from "@mui/material";
import { Box } from "@mui/system";
import React from "react";
import { round } from "../../helpers/generalHelpers";
import { IDiversification, IRecord } from "../../types/record";
import { MyPie } from "../Charts/PieChart";
import { MyCard } from "./MyCard";

interface IProps {
  currentRecord?: IRecord | null;
}

export const GeneralDiversification = ({ currentRecord }: IProps) => {
  const total = currentRecord
    ? currentRecord.stocksValue +
      currentRecord.liquidity +
      currentRecord.cryptosValue
    : 1;
  if (!currentRecord) return null;
  console.log(currentRecord);
  return (
    <MyCard>
      <Typography gutterBottom variant="h5" textAlign="center">
        Diversification
      </Typography>
      <MyPie
        showLegend
        enableArcLinkLabels={false}
        data={
          [
            {
              symbol: "Stocks",
              percent: round((currentRecord.stocksValue / total) * 100),
            },
            {
              symbol: "Crypto",
              percent: round((currentRecord.cryptosValue / total) * 100),
            },
            {
              symbol: "Liquidity",
              percent: round((currentRecord.liquidity / total) * 100),
            },
          ] as IDiversification[]
        }
      />
    </MyCard>
  );
};
