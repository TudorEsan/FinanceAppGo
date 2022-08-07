import React from "react";
import axios from "../axiosConfig";
import { INetWorth } from "../types/overview";

export const getNetWorthOverviewReq = async (year: number) => {
  const resp = await axios.get(`/overview/networth`, {
    params: {
      year: year,
    },
  });
  return resp.data.records as INetWorth[];
};
