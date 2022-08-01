import React, { useEffect } from "react";
import { getErrorMessage } from "../helpers/errors";
import { getRecordsReq } from "../service/RecordService";
import { IRecord } from "../types/record";

export const useRecords = () => {
  const [records, setRecords] = React.useState<IRecord[]>([]);
  const [loading, setLoading] = React.useState(true);
  const [error, setError] = React.useState<null | string>(null);

  const getRecords = async () => {
		setLoading(true);
		setError(null);
    try {
      const rec = await getRecordsReq();
      setRecords(rec);
    } catch (e) {
      console.error(e);
      setError(getErrorMessage(e));
    }
    setLoading(false);
  };

  useEffect(() => {
    getRecords();
  }, []);

  return { records, loading, error };
};
