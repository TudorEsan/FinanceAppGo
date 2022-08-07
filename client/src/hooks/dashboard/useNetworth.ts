import { Datum, Serie } from "@nivo/line";
import React, { useEffect } from "react";
import { formatDate } from "../../helpers/date";
import { getErrorMessage } from "../../helpers/errors";
import { getNetWorthOverviewReq } from "../../service/OverviewService";
import { ILinear, INetWorth } from "../../types/overview";
import { format } from "date-fns";

export const useNetworthOverview = () => {
  const [netWorth, setNetWorth] = React.useState<Datum[]>([]);
  const [loading, setLoading] = React.useState(false);
  const [error, setError] = React.useState<string | null>(null);
  const [year, setYear] = React.useState(new Date().getFullYear());

  const formatOverview = (overview: INetWorth[]) => {
    return overview.map((item: INetWorth, index) => {
      return {
        x: format(new Date(item.date), "dd/MM/yy"),
        y: item.total,
      } as Datum;
    });
  };

  const getNetWorth = async (year: number) => {
    setLoading(true);
    try {
      const overview = await getNetWorthOverviewReq(year);
      setNetWorth(formatOverview(overview));
    } catch (error) {
      setError(getErrorMessage(error));
    }
    setLoading(false);
  };
  useEffect(() => {
    getNetWorth(year);
  }, [year]);

  return { netWorth, loading, error, setYear };
};
