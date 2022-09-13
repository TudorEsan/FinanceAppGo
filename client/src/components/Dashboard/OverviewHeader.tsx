import { Divider, Grid, Typography } from "@mui/material";
import { Box } from "@mui/system";
import React from "react";
import { IRecord } from "../../types/record";
import AttachMoneyIcon from "@mui/icons-material/AttachMoney";
import { MyCard } from "../Cards/MyCard";
import TrendingUpIcon from "@mui/icons-material/TrendingUp";
import TrendingDownIcon from "@mui/icons-material/TrendingDown";
import { formatDate } from "../../helpers/date";
import { getTotalNetWorth } from "../../helpers/recordHelper";
import { round } from "../../helpers/generalHelpers";
import { Money } from "@mui/icons-material";

interface IProps {
  lastRecord?: IRecord | null;
  currentRecord?: IRecord | null;
}
interface IItemContentProps {
  title: string;
  upByPercent: number;
  upBy: number;
  value: number;
  icon: React.ReactNode;
  since: Date;
}
const ItemContent = ({
  title,
  upByPercent,
  upBy,
  value,
  icon,
  since,
}: IItemContentProps) => {
  return (
    <Grid container spacing={0.5} padding={1}>
      <Grid item xs={12} mb={2}>
        <Grid item>{icon}</Grid>
      </Grid>
      <Grid item xs={12}>
        <Typography variant="body1" color="gray">
          {title}
        </Typography>
      </Grid>
      <Grid item xs={12}>
        <Typography fontSize="1.5rem">$ {value}</Typography>
      </Grid>
      <Grid container item xs={12} spacing={2}>
        <Grid item>
          <Box sx={{ display: "flex", alignItems: "center", gap: 1 }}>
            {upByPercent >= 0 ? (
              <TrendingUpIcon color="success" />
            ) : (
              <TrendingDownIcon color="error" />
            )}
            <Typography color="textPrimary">
              {(upBy >= 0 ? "+" : "") + upByPercent}%
            </Typography>
          </Box>
        </Grid>
        <Grid item>
          <Typography color="gray">
            {(upBy >= 0 ? "+" : "") + `${String(upBy)} $`}
          </Typography>
        </Grid>
      </Grid>
      <Grid item>
        <Typography color="gray">Since {formatDate(since)}</Typography>
      </Grid>
    </Grid>
  );
};

export const OverviewHeader = ({ lastRecord, currentRecord }: IProps) => {
  return (
    <Grid item container xs={12} spacing={2}>
      <Grid item md={6} xs={12}>
        <MyCard>
          <ItemContent
            title="Net Worth"
            icon={
              <AttachMoneyIcon
                color="primary"
                sx={{
                  fontSize: "50px",
                  background: "#252525",
                  borderRadius: "50%",
                  padding: 1,
                }}
              />
            }
            since={lastRecord!.date}
            upByPercent={round(
              (getTotalNetWorth(currentRecord) * 100) /
                getTotalNetWorth(lastRecord) -
                100
            )}
            upBy={round(
              getTotalNetWorth(currentRecord) - getTotalNetWorth(lastRecord)
            )}
            value={round(getTotalNetWorth(currentRecord))}
          />
        </MyCard>
      </Grid>
      <Grid item xs={12} md={6}>
        <MyCard>
          <ItemContent
            title="Liquidity"
            icon={
              <Money
                color="primary"
                sx={{
                  fontSize: "50px",
                  background: "#252525",
                  borderRadius: "50%",
                  padding: 1,
                }}
              />
            }
            since={lastRecord!.date}
            upByPercent={round(
              (currentRecord!.liquidity * 100) / lastRecord!.liquidity - 100
            )}
            upBy={currentRecord!.liquidity - lastRecord!.liquidity}
            value={currentRecord!.liquidity}
          />
        </MyCard>
      </Grid>
    </Grid>
  );
};
