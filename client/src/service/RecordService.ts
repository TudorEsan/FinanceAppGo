import axios from "axios";
import { IRecord, IRecordForm } from "../types/record";

export const addRecordReq = async (record: IRecordForm) => {
  return axios.post("/api/networth/record", record);
};

export const getRecordsReq = async (): Promise<IRecord[]> => {
  const resp = await axios.get("/api/networth/");
  return resp.data.netWorth.records as IRecord[];
};

export const getRecordReq = async (id: string): Promise<IRecord> => {
  const resp = await axios.get(`/api/networth/record/${id}`);
  return resp.data.netWorth.record as IRecord;
};
