import React from "react";
import axios from "../axiosConfig";
import { INetWorth, IOverview } from "../types/overview";

export const getNetWorthOverviewReq = async () => {
  console.log("help");
  const resp = await axios.get(`/overview/networth`);
  return resp.data.overview as IOverview;
};
