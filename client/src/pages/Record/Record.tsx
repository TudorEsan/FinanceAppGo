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
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
} from "@mui/material";
import { Box } from "@mui/system";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import React from "react";
import { Navigate, useNavigate } from "react-router-dom";
import { MyCard } from "../../components/Cards/MyCard";
import { MyPie } from "../../components/Charts/PieChart";
import { formatDate } from "../../helpers/date";
import { useMobile } from "../../hooks/useMobile";
import { useRecord } from "../../hooks/useRecord";
import { ICrypto, IRecord, IStock } from "../../types/record";

interface IStockGridProps {
  stocks: IStock[];
}

const stocksCol: GridColDef[] = [
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

const DesktopTable = ({ record }: { record: IRecord }) => (
  <>
    <TableHead>
      <TableRow>
        <TableCell>Total $</TableCell>
        <TableCell>Liquidity</TableCell>
        <TableCell>Invested Amount</TableCell>
        <TableCell>Crypto Value</TableCell>
        <TableCell>Stocks Value</TableCell>
      </TableRow>
    </TableHead>
    <TableBody>
      <TableRow sx={{ "&:last-child td, &:last-child th": { border: 0 } }}>
        <TableCell>{record!.investedAmount + record!.liquidity}</TableCell>
        <TableCell>{record!.liquidity} $</TableCell>
        <TableCell>{record!.investedAmount} $</TableCell>
        <TableCell>{record!.cryptosValue} $</TableCell>
        <TableCell>{record!.stocksValue} $</TableCell>
      </TableRow>
    </TableBody>
  </>
);

const MobileTable = ({ record }: { record: IRecord }) => (
  <>
    <TableBody>
      <TableRow sx={{ "&:last-child td, &:last-child th": { border: 0 } }}>
        <TableCell>Total $</TableCell>
        <TableCell>{record!.investedAmount + record!.liquidity}</TableCell>
      </TableRow>
      <TableRow sx={{ "&:last-child td, &:last-child th": { border: 0 } }}>
        <TableCell>Invested Amount</TableCell>
        <TableCell>{record!.liquidity} $</TableCell>
      </TableRow>
      <TableRow sx={{ "&:last-child td, &:last-child th": { border: 0 } }}>
        <TableCell>Crypto Value</TableCell>
        <TableCell>{record!.cryptosValue} $</TableCell>
      </TableRow>
      <TableRow sx={{ "&:last-child td, &:last-child th": { border: 0 } }}>
        <TableCell>Stocks Value</TableCell>
        <TableCell>{record!.stocksValue} $</TableCell>
      </TableRow>
      <TableRow sx={{ "&:last-child td, &:last-child th": { border: 0 } }}>
        <TableCell>Invested Amount</TableCell>
        <TableCell>{record!.investedAmount} $</TableCell>
      </TableRow>
    </TableBody>
  </>
);

const OverviewTable = ({ record }: { record: IRecord }) => {
  const isMobile = useMobile();
  return (
    <TableContainer component={Paper} sx={{ mt: 4 }}>
      <Table aria-label="simple table">
        {!isMobile && <DesktopTable record={record} />}
        {isMobile && <MobileTable record={record} />}
      </Table>
    </TableContainer>
  );
};

export const Record = () => {
  const { record, loading, error, deleteRecord, diversification } = useRecord();
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
      <MyCard>
        <Box
          display="flex"
          flexWrap="wrap"
          justifyContent="space-between"
          alignItems="center"
        >
          <Typography variant="h4">Record</Typography>
          <Box>
            <Button
              sx={{ mr: 2 }}
              variant="contained"
              color="primary"
              onClick={() => navigate("edit")}
            >
              Edit
            </Button>
            <Button
              sx={{ mr: 2 }}
              variant="contained"
              color="primary"
              onClick={() => navigate("clone")}
            >
              Clone
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
        <Box
          display="flex"
          flexWrap="wrap"
          justifyContent="center"
          alignItems="center"
        >
          <MyPie data={record!.data!.cryptoDiversification} />
          <MyPie data={record!.data!.stockDiversification} />
          <MyPie data={diversification} />
        </Box>
        <OverviewTable record={record!.data!} />
      </MyCard>
      <CryptoGrid cryptos={record!.data!.cryptos} />
      <StocksGrid stocks={record!.data!.stocks} />
    </>
  );
};
