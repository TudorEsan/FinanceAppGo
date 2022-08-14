import { GridCallbackDetails } from "@mui/x-data-grid";
import React, { useEffect } from "react";
import { getErrorMessage } from "../helpers/errors";
import { getRecordCountReq, getRecordsReq } from "../service/RecordService";
import { IRecord } from "../types/record";

export const useRecords = () => {
  const [records, setRecords] = React.useState<IRecord[]>([]);
  const [loading, setLoading] = React.useState(true);
  const [error, setError] = React.useState<null | string>(null);
  const [recordCount, setRecordCount] = React.useState(0);
  const [pageSize, setPageSize] = React.useState(10);
  const [page, setPage] = React.useState(0);

  const getRecords = async (page = 0, perPage = 0) => {
    setLoading(true);
    setError(null);
    try {
      const rec = await getRecordsReq(page, perPage);
      if (recordCount === 0) {
        setRecordCount(await getRecordCountReq());
      }
      setRecords(rec);
    } catch (e) {
      console.error(e);
      setError(getErrorMessage(e));
    }
    setLoading(false);
  };

  const handlePageSizeChange = (pageSize: number) => {
    setPageSize(pageSize);
  };

  const handlePageChange = (page: number) => {
    setPage(page);
  };

  useEffect(() => {
    getRecords(page, pageSize);
  }, [pageSize, page]);

  return {
    records,
    loading,
    error,
    handlePageChange,
    pageSize,
    recordCount,
    handlePageSizeChange,
  };
};
