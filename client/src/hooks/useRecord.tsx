import React from "react";
import { useParams } from "react-router-dom";
import { getErrorMessage } from "../helpers/errors";
import { getDiversification } from "../helpers/recordHelper";
import { handleError, handleSuccess } from "../helpers/state";
import {
  deleteRecordReq,
  getRecordReq,
  updateRecordReq,
} from "../service/RecordService";
import { IRequestState } from "../types/general";
import { IDiversification, IRecord, IRecordForm } from "../types/record";

export const useRecord = () => {
  const [record, setRecord] = React.useState<IRequestState<IRecord>>({
    data: null,
    error: null,
    loading: true,
  });
  const [diversification, setDiversification] = React.useState<
    IDiversification[]
  >([]);
  const [loading, setLoading] = React.useState(false);
  const [error, setError] = React.useState<string | null>(null);
  const { recordId: id } = useParams();

  const getRecord = async () => {
    try {
      const record = await getRecordReq(id!);
      setDiversification(getDiversification(record));
      handleSuccess(record, setRecord);
    } catch (e) {
      console.error(e);
      handleError(setRecord, getErrorMessage(e));
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

  const updateRecord = async (record: IRecordForm) => {
    setLoading(true);
    setError(null);
    try {
      const updatedRecord = await updateRecordReq(id!, record);
      
    } catch (e) {
      setError(getErrorMessage(e));
      setLoading(false);
      throw e;
    }
    setLoading(false);
  };

  React.useEffect(() => {
    getRecord();
  }, [id]);

  return {
    record,
    loading,
    error,
    deleteRecord,
    diversification,
    updateRecord,
  };
};
