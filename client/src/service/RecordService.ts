import axios from "axios";
import { IRecord, IRecordForm } from "../types/record";

export const addRecordReq = async (record: IRecordForm) => {
  console.log(record.date);
  return axios.post("/api/networth/record", record);
};

export const getRecordsReq = async (): Promise<IRecord[]> => {
  const resp = await axios.get("/api/networth/");
  return resp.data.netWorth.records as IRecord[];
};

export const getRecordReq = async (id: string): Promise<IRecord> => {
  const resp = await axios.get(`/api/networth/record/${id}`);
  return resp.data.record as IRecord;
};

export const deleteRecordReq = async (id: string) => {
  return axios.delete(`/api/networth/record/${id}`);
};

export const updateRecordReq = async (id: string, data: IRecordForm) => {
  return axios.put(`/api/networth/record/${id}`, data);
};
