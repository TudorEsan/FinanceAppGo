import { Add, Remove } from "@mui/icons-material";
import {
  Box,
  Button,
  Card,
  CardContent,
  CircularProgress,
  Divider,
  Grid,
  IconButton,
  Typography,
} from "@mui/material";
import { useFieldArray, useForm } from "react-hook-form";
import { IRecordForm } from "../../types/record";
import {
  ControlledTextField,
  ControlledDatePicker,
  Stocks,
  Cryptos,
} from "../../components";
import * as Yup from "yup";
import { yupResolver } from "@hookform/resolvers/yup";
import { useAddRecord } from "../../hooks/useAddRecord";
import { useRecord } from "../../hooks/useRecord";
import { useEffect } from "react";
import { Navigate, useNavigate } from "react-router-dom";

const formSchema = Yup.object({
  date: Yup.string().required("Date is required"),
  liquidity: Yup.number()
    .required("Liquidity is required")
    .min(0, "Liquidity must be greater or equal to 0"),
  stocks: Yup.array().of(
    Yup.object().shape({
      symbol: Yup.string().required("Symbol is required"),
      valuedAt: Yup.number()
        .required("Field is required")
        .min(0, "Price must be greater or equal to 0"),
      shares: Yup.number()
        .required("Field is required")
        .min(0, "Shares must be greater or equal to 0"),
    })
  ),
  cryptos: Yup.array().of(
    Yup.object().shape({
      symbol: Yup.string().required("Symbol is required"),
      valuedAt: Yup.number()
        .required("Field is required")
        .min(0, "Price must be greater or equal to 0"),
      coins: Yup.number()
        .required("Field is required")
        .min(0, "Coins must be greater or equal to 0"),
    })
  ),
});

export const EditRecord = () => {
  const { record, loading, error, updateRecord } = useRecord();

  const {
    control,
    handleSubmit,
    reset,
  } = useForm<IRecordForm>({
    resolver: yupResolver(formSchema),
  });
  const navigate = useNavigate();

  const {
    fields: stockFields,
    append: stockAppend,
    remove: stockRemove,
  } = useFieldArray({
    control,
    name: "stocks",
  });
  const {
    fields: cryptoFields,
    append: cryptoAppend,
    remove: cryptoRemove,
  } = useFieldArray({
    control,
    name: "cryptos",
  });

  const appendStock = () => {
    stockAppend({ shares: 0, valuedAt: 0, symbol: "" });
  };
  const appendCrypto = () => {
    cryptoAppend({ coins: 0, valuedAt: 0, symbol: "" });
  };

  useEffect(() => {
    if (record.data !== null) {
      console.log(record.data);
      reset(record.data);
    }
  }, [record]);

  const onSubmit = (data: IRecordForm) => {
    data.date = new Date(data.date).toISOString();
    updateRecord(data);
    navigate(-1);
  };
  if (record.loading) {
    return <CircularProgress />;
  }
  if (record.error) {
    return <Typography color="error">{record.error}</Typography>;
  }
  return (
    <Card>
      <CardContent>
        <form onSubmit={handleSubmit(onSubmit)}>
          <Typography variant="h4" mb={4}>
            Update Record
          </Typography>
          <Grid container>
            <Grid item container md={11} spacing={2}>
              <Grid item md={6} xs={12}>
                <ControlledDatePicker
                  name="date"
                  label="Date"
                  control={control}
                />
              </Grid>
              <Grid item md={6} xs={12}>
                <ControlledTextField
                  name="liquidity"
                  label="Liquidity"
                  control={control}
                  type="number"
                />
              </Grid>
            </Grid>
          </Grid>
          <Divider sx={{ mb: 2, mt: 2 }} />
          <Grid container mt={2} spacing={2}>
            <Grid item md={11} sm={10} justifyContent="space-between">
              <Typography gutterBottom variant="h5">
                Stocks
              </Typography>
            </Grid>
            <Grid item md={1} sm={2} sx={{ textAlign: "center" }}>
              <IconButton onClick={() => appendStock()} color="primary">
                <Add color="primary" />
              </IconButton>
            </Grid>
          </Grid>
          <Stocks remove={stockRemove} fields={stockFields} control={control} />
          <Grid container mt={2} spacing={2} alignItems="center">
            <Grid item md={11} sm={10} justifyContent="space-between">
              <Typography gutterBottom variant="h5">
                Cryptos
              </Typography>
            </Grid>
            <Grid item md={1} sm={2} sx={{ textAlign: "center" }}>
              <IconButton onClick={() => appendCrypto()} color="primary">
                <Add color="primary" />
              </IconButton>
            </Grid>
          </Grid>
          <Cryptos
            remove={cryptoRemove}
            fields={cryptoFields}
            control={control}
          />
          <Box mt={3}>
            <Button
              variant="contained"
              color="primary"
              type="submit"
              sx={{ mr: 1 }}
            >
              Update
            </Button>
            <Button variant="outlined" color="primary">
              Discard
            </Button>
          </Box>
        </form>
        {error && (
          <Typography variant="body2" color="error">
            {error}
          </Typography>
        )}
        {loading && <CircularProgress />}
      </CardContent>
    </Card>
  );
};
