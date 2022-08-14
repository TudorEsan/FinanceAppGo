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
import { useRecord } from "../../hooks/useRecord";
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { recordFormSchema } from "../../helpers/validation/recordValidation";
import { getUtcIso } from "../../helpers/date";
import { useCloneRecord } from "../../hooks/records/useCloneRecord";

export const CloneRecord = () => {
  const { cloneRecord, loading, error, addRecord } = useCloneRecord();

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
    if (cloneRecord.data !== null) {
      console.log(cloneRecord.data);
      cloneRecord.data.date = new Date(getUtcIso(cloneRecord.data.date));
      reset(cloneRecord.data);
    }
  }, [cloneRecord]);

  const onSubmit = async (data: IRecordForm) => {
    data.date = getUtcIso(data.date);
    console.log(data);
    try {
      await addRecord(data);
      navigate("records");
    } catch (e) {
      console.error(e);
    }
  };
  if (cloneRecord.loading) {
    return <CircularProgress />;
  }
  if (cloneRecord.error) {
    return <Typography color="error">{cloneRecord.error}</Typography>;
  }
  return (
    <Card>
      <CardContent>
        <form onSubmit={handleSubmit(onSubmit)}>
          <Typography variant="h4" mb={4}>
            Clone Record
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
              Add
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
