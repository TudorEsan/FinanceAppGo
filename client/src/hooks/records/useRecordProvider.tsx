import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { getErrorMessage } from "../../helpers/errors";
import { round } from "../../helpers/generalHelpers";
import { addRecordReq } from "../../service/RecordService";
import { IRecordForm } from "../../types/record";

export const useRecordProvider = () => {
  // provides record functions
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<null | string>(null);

  const addRecord = async (data: IRecordForm) => {
    try {
      setLoading(true);
      // add valuedAt to each stock and crypto
      data.stocks.forEach((stock) => {
        stock.valuedAt = round(stock.shares * stock.currentPrice, 2);
      });
      data.cryptos.forEach((crypto) => {
        crypto.valuedAt = round(crypto.coins * crypto.currentPrice, 2);
      });
      const response = await addRecordReq(data);
      setError(null);
    } catch (error) {
      setError(getErrorMessage(error));
      setLoading(false);
      throw error;
    }
    setLoading(false);
  };
  return { loading, error, addRecord };
};
