import { Add } from "@mui/icons-material";
import { Button, Icon, Typography } from "@mui/material";
import { Box } from "@mui/system";
import { DataGrid, GridColDef, GridRenderCellParams } from "@mui/x-data-grid";
import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { formatDate } from "../helpers/date";
import { useRecords } from "../hooks/useRecords";

const columns: GridColDef[] = [
  // { field: 'id', headerName: 'ID', width: 100 },
  {
    field: "date",
    headerName: "Date",
    renderCell: (rowData: GridRenderCellParams<Date>) => {
      if (rowData.id) {
        const date = new Date(rowData.value || Date.now());
        return (
          <Link
            style={{
              color: "#17C6B1",
              fontSize: "1rem",
              textDecoration: "none",
            }}
            to={rowData.id as string}
          >
            {formatDate(date)}
          </Link>
        );
      }
      return "error";
    },
    flex: 1,
    editable: false,
  },
  {
    field: "investedAmount",
    headerName: "Invested Amount",
    flex: 1,
    editable: false,
  },
  {
    field: "cryptosValue",
    headerName: "Crypto Value",
    flex: 1,
    editable: false,
  },
  {
    field: "stocksValue",
    headerName: "Stocks Value",
    flex: 1,
    editable: false,
  },
];

const mobileColums = [
  {
    field: "date",
    headerName: "Date",
    flex: 1,
    editable: false,
  },
  {
    field: "investedAmount",
    headerName: "Invested Amount",
    flex: 1,
    editable: false,
  },
];

const RecordGrid = () => {
  const { records, loading, error } = useRecords();
  const [isMobile, setIsMobile] = useState(false);
  const navigate = useNavigate();
  //choose the screen size
  const handleResize = () => {
    if (window.innerWidth < 720) {
      setIsMobile(true);
    } else {
      setIsMobile(false);
    }
  };

  useState(() => {
    window.addEventListener("resize", handleResize);
    return () => {
      window.removeEventListener("resize", handleResize);
    };
  });

  return (
    <>
      <DataGrid
        columns={isMobile ? mobileColums : columns}
        rows={records}
        loading={loading}
        error={error !== null ? undefined : error}
        disableSelectionOnClick
        hideFooter
        autoHeight
      />
      {error && (
        <Typography mt={2} variant="body1" color="error">
          {error}
        </Typography>
      )}
    </>
  );
};

export const Records = () => {
  const navigate = useNavigate();
  return (
    <>
      <Box mb={2} justifyContent="space-between" alignItems="center" display="flex">
        <Typography variant="h4">Records</Typography>
        <Button
          size="small"
          variant="contained"
          color="primary"
          startIcon={<Add />}
          onClick={() => navigate("add")}
        >
          Add Record
        </Button>
      </Box>
      <RecordGrid />
    </>
  );
};
