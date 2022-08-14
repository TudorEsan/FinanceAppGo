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
import { yupResolver } from "@hookform/resolvers/yup";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { recordFormSchema } from "../../helpers/validation/recordValidation";
import { getUtcIso } from "../../helpers/date";
import { useRecord } from "../../hooks/useRecord";

export const EditRecord = () => {
  const { record, loading, error, updateRecord } = useRecord();

  const { control, handleSubmit, reset } = useForm<IRecordForm>({
    resolver: yupResolver(recordFormSchema),
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
      record.data.date = new Date(getUtcIso(record.data.date));
      reset(record.data);
    }
  }, [record]);

  const onSubmit = async (data: IRecordForm) => {
    data.date = getUtcIso(data.date);
    console.log(data);
    try {
      await updateRecord(data);
      navigate(-1);
    } catch (e) {
      console.error(e);
    }
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
