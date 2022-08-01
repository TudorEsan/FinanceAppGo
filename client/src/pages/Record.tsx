import { CircularProgress, Typography } from "@mui/material";
import { Box } from "@mui/system";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import React from "react";
import { formatDate } from "../helpers/date";
import { useRecord } from "../hooks/useRecord";
import { ICrypto, IStock } from "../types/record";

interface IStockGridProps {
  stocks: IStock[];
}

const stocksCol: GridColDef[] = [
  {
    field: "name",
    headerName: "Name",
    flex: 1,
    editable: false,
  },
  {
    field: "symbol",
    headerName: "Symbol",
    flex: 1,
    editable: false,
  },
  {
    field: "shares",
    headerName: "Shares",
    flex: 1,
    editable: false,
  },
  {
    field: "valuedAt",
    headerName: "Value",
    flex: 1,
    editable: false,
  },
];
// test

const StocksGrid = ({ stocks }: IStockGridProps) => {
  return (
    <Box mt={4}>
      <Typography gutterBottom variant="h6">
        Stocks
      </Typography>
      <DataGrid
        columns={stocksCol}
        rows={stocks}
        rowHeight={50}
        getRowId={(row) => row.symbol}
        autoHeight
        hideFooter
      />
    </Box>
  );
};
interface ICryptoGridProps {
  cryptos: ICrypto[];
}

const cryptoColumns: GridColDef[] = [
  {
    field: "name",
    headerName: "Name",
    flex: 1,
    editable: false,
  },
  {
    field: "symbol",
    headerName: "Symbol",
    flex: 1,
    editable: false,
  },
  {
    field: "coins",
    headerName: "Coins",
    flex: 1,
    editable: false,
  },
  {
    field: "valuedAt",
    headerName: "Value",
    flex: 1,
    editable: false,
  },
];

const CryptoGrid = ({ cryptos }: ICryptoGridProps) => {
  return (
    <Box mt={4}>
      <Typography gutterBottom variant="h6">
        Crypto
      </Typography>
      <DataGrid
        getRowId={(row) => row.symbol}
        columns={cryptoColumns}
        rows={cryptos}
        rowHeight={50}
        autoHeight
        hideFooter
      />
    </Box>
  );
};

export const Record = () => {
  const { record, loading, error } = useRecord();
  if (loading) {
    return <CircularProgress />;
  }

  if (error !== null) {
    return <Typography variant="h6">{error}</Typography>;
  }

  console.log(record, error, loading);

  return (
    <>
      <Typography variant="h4">Record</Typography>
      <Typography variant="h6">From: {formatDate(record!.date)} $</Typography>
      <Typography variant="h6" mb={2} gutterBottom>
        Total: {record!.investedAmount + record!.liquidity} $
      </Typography>
      <Typography gutterBottom>Liquidity: {record!.liquidity} $</Typography>
      <Typography gutterBottom>
        Invested Amount: {record!.investedAmount} $
      </Typography>
      <Typography gutterBottom>
        Crypto Value: {record!.cryptosValue} $
      </Typography>
      <Typography gutterBottom>
        Stocks Value: {record!.stocksValue} $
      </Typography>
      <CryptoGrid cryptos={record!.cryptos} />
      <StocksGrid stocks={record!.stocks} />
    </>
  );
};
