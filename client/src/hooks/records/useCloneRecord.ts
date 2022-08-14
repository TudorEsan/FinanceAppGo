import React from "react";
import { handleError, handleSuccess } from "../../helpers/state";
import { IRequestState } from "../../types/general";
import { IRecord } from "../../types/record";
import { useRecord } from "../useRecord";
import { useRecordProvider } from "./useRecordProvider";

export const useCloneRecord = () => {
  const { record } = useRecord();
  const [cloneRecord, setCloneRecord] = React.useState<IRequestState<IRecord>>({
    data: null,
    loading: true,
    error: null,
  });
  const { addRecord, loading, error } = useRecordProvider();

  React.useEffect(() => {
    if (record.data !== null) {
      handleSuccess(record.data, setCloneRecord);
    }
    if (record.error !== null) {
      handleError(setCloneRecord, record.error);
    }
  }, [record]);

  return { cloneRecord, addRecord, loading, error };
};
