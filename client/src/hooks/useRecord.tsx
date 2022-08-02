import React from "react";
import { useParams } from "react-router-dom";
import { getErrorMessage } from "../helpers/errors";
import { getRecordReq } from "../service/RecordService";
import { IRecord } from "../types/record";

export const useRecord = () => {
  const [record, setRecord] = React.useState<IRecord | null>(null);
  const [loading, setLoading] = React.useState(true);
  const [error, setError] = React.useState<string | null>(null);
  const { recordId: id } = useParams();

  const getRecord = async () => {
    setLoading(true);
    setError(null);
    try {
      const record = await getRecordReq(id!);
      setRecord(record);
    } catch (e) {
      console.error(e);
      setError(getErrorMessage(e));
    }
    setLoading(false);
  };

  React.useEffect(() => {
    getRecord();
  }, [id]);

  return { record, loading, error };
};
