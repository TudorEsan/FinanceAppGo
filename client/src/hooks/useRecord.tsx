import React from "react";
import { useParams } from "react-router-dom";
import { getErrorMessage } from "../helpers/errors";
import { handleError, handleSuccess } from "../helpers/state";
import { deleteRecordReq, getRecordReq } from "../service/RecordService";
import { IRequestState } from "../types/general";
import { IRecord } from "../types/record";

export const useRecord = () => {
  const [record, setRecord] = React.useState<IRequestState<IRecord>>({
    data: null,
    error: null,
    loading: true,
  });
  const [loading, setLoading] = React.useState(true);
  const [error, setError] = React.useState<string | null>(null);
  const { recordId: id } = useParams();

  const getRecord = async () => {
    try {
      const record = await getRecordReq(id!);
      handleSuccess(record, setRecord);
    } catch (e) {
      console.error(e);
      handleError(setError, getErrorMessage(e));
    }
  };

  const deleteRecord = async () => {
    setLoading(true);
    setError(null);
    try {
      await deleteRecordReq(id!);
    } catch (e) {
      console.error(e);
      setError(getErrorMessage(e));
    }
    setLoading(false);
  };

  React.useEffect(() => {
    getRecord();
  }, [id]);

  return { record, loading, error, deleteRecord };
};
