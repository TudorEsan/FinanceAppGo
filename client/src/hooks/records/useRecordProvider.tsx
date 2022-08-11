import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { getErrorMessage } from "../../helpers/errors";
import { addRecordReq } from "../../service/RecordService";
import { IRecordForm } from "../../types/record";

export const useRecordProvider = () => {
  // provides record functions 
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<null | string>(null);
  const addRecord = async (data: IRecordForm) => {
    let response: any;
    try {
      setLoading(true);
      response = await addRecordReq(data);
      setError(null);
    } catch (error) {
      setError(getErrorMessage(error));
      throw error;
    }
    setLoading(false);
  };
  return { loading, error, addRecord };
};
