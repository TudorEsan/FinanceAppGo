import { Datum, Serie } from "@nivo/line";
import React, { useEffect } from "react";
import { formatDate } from "../../helpers/date";
import { getErrorMessage } from "../../helpers/errors";
import { getNetWorthOverviewReq } from "../../service/OverviewService";
import { ILinear, ILiquidity, INetWorth } from "../../types/overview";
import { format } from "date-fns";
import { IRecord } from "../../types/record";

export const useNetworthOverview = () => {
  const [netWorth, setNetWorth] = React.useState<Datum[]>([]);
  const [liquidity, setLiquidity] = React.useState<Datum[]>([]);
  const [loading, setLoading] = React.useState(true);
  const [error, setError] = React.useState<string | null>(null);
  const [currentRecord, setCurrentRecord] = React.useState<IRecord | null>(null);
  const [lastRecord, setLastRecord] = React.useState<IRecord | null>(null);


  const formatOverview = (overview: INetWorth[]) => {
    return overview.map((item: INetWorth, index) => {
      return {
        x: format(new Date(item.date), "dd/MM/yy"),
        y: item.total,
      } as Datum;
    });
  };
  const formatLiquidity = (liquidityOverv: ILiquidity[]) => {
    return liquidityOverv.map((item, index) => {
      return {
        x: format(new Date(item.date), "dd/MM/yy"),
        y: item.liquidity,
      } as Datum;
    });
  };

  const getNetWorth = async () => {
    setLoading(true);
    try {
      const overview = await getNetWorthOverviewReq();
      setNetWorth(formatOverview(overview.networthOverview));
      setLiquidity(formatLiquidity(overview.liquidityOverview))
      setCurrentRecord(overview.currentRecord);
      setLastRecord(overview.lastRecord);
    } catch (error) {
      setError(getErrorMessage(error));
    }
    setLoading(false);
  };
  useEffect(() => {
    getNetWorth();
  }, []);

  return { netWorth, loading, error, liquidity, currentRecord, lastRecord};
};
