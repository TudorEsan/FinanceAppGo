import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { getErrorMessage } from "../helpers/errors";
import { addRecordReq } from "../service/RecordService";
import { IRecordForm } from "../types/record";

export const useAddRecord = () => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<null | string>(null);
  const navigate = useNavigate();
  const addRecord = async (data: IRecordForm) => {
    try {
      setLoading(true);
      await addRecordReq(data);
      setError(null);
      navigate("/records");
    } catch (error) {
      setError(getErrorMessage(error));
    }
    setLoading(false);
  };
  return { loading, error, addRecord };
};
