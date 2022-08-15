import { Divider, Grid, Typography } from "@mui/material";
import { Box } from "@mui/system";
import React from "react";
import { IRecord } from "../../types/record";
import AttachMoneyIcon from "@mui/icons-material/AttachMoney";
import { MyCard } from "../Cards/MyCard";
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
}
const ItemContent = ({
  title,
  upByPercent,
  upBy,
  value,
  icon,
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
      <Grid container item xs={12}>
        <Grid item>
          <Typography color="textPrimary">{upByPercent}%</Typography>
        </Grid>
        <Grid item>
          <Typography color="textPrimary">{upBy}</Typography>
        </Grid>
      </Grid>
    </Grid>
  );
};

export const OverviewHeader = ({ lastRecord, currentRecord }: IProps) => {
  return (
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
        upByPercent={0}
        upBy={0}
        value={0}
      />
    </MyCard>
  );
};
