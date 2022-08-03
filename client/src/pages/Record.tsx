import {
  Button,
  Card,
  CardContent,
  CircularProgress,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  Divider,
  Typography,
} from "@mui/material";
import { Box } from "@mui/system";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import React from "react";
import { Navigate, useNavigate } from "react-router-dom";
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

const ConfirmationDialog = ({
  open,
  handleClose,
  handleConfirm,
}: {
  open: boolean;
  handleClose: () => void;
  handleConfirm: () => void;
}) => {
  return (
    <div>
      <Dialog
        open={open}
        onClose={handleClose}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
        <DialogTitle id="alert-dialog-title">Warning</DialogTitle>
        <DialogContent>
          <DialogContentText id="alert-dialog-description">
            Are you sure you want to delete this record?
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>Disagree</Button>
          <Button onClick={handleConfirm} autoFocus>
            Agree
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  );
};

export const Record = () => {
  const { record, loading, error, deleteRecord } = useRecord();
  const navigate = useNavigate();
  const [confirmationOpen, setConfirmationOpen] = React.useState(false);

  const handleClose = () => {
    setConfirmationOpen(false);
  };

  const handleConfirm = async () => {
    await deleteRecord();
    setConfirmationOpen(false);
    navigate(-1);
  };

  const openConfirmation = () => {
    setConfirmationOpen(true);
  };

  if (record.loading) {
    return <CircularProgress />;
  }

  if (record.error !== null) {
    return (
      <Typography variant="h6" color="error">
        {record.error}
      </Typography>
    );
  }

  return (
    <>
      <ConfirmationDialog
        open={confirmationOpen}
        handleClose={handleClose}
        handleConfirm={handleConfirm}
      />
      <Card elevation={10}>
        <CardContent>
          <Box
            display="flex"
            justifyContent="space-between"
            alignItems="center"
          >
            <Typography variant="h4">Record</Typography>
            <Box>
              <Button sx={{ mr: 2 }} variant="contained" color="primary">
                Edit
              </Button>
              <Button
                variant="outlined"
                onClick={() => openConfirmation()}
                color="primary"
              >
                Delete
              </Button>
            </Box>
          </Box>
          <Divider sx={{ mt: 1, mb: 1 }} />
          <Typography variant="h6">
            From: {formatDate(record!.data!.date)} $
          </Typography>
          <Typography variant="h6" mb={2} gutterBottom>
            Total: {record!.data!.investedAmount + record!.data!.liquidity} $
          </Typography>
          <Typography gutterBottom>
            Liquidity: {record!.data!.liquidity} $
          </Typography>
          <Typography gutterBottom>
            Invested Amount: {record!.data!.investedAmount} $
          </Typography>
          <Typography gutterBottom>
            Crypto Value: {record!.data!.cryptosValue} $
          </Typography>
          <Typography gutterBottom>
            Stocks Value: {record!.data!.stocksValue} $
          </Typography>
        </CardContent>
      </Card>
      <CryptoGrid cryptos={record!.data!.cryptos} />
      <StocksGrid stocks={record!.data!.stocks} />
    </>
  );
};
