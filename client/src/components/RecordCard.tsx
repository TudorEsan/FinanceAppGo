import { Add, Remove } from "@mui/icons-material";
import {
  Box,
  Button,
  Card,
  CardContent,
  CircularProgress,
  Divider,
  Grid,
  Icon,
  IconButton,
  Typography,
} from "@mui/material";
import {
  Control,
  FieldArrayWithId,
  useFieldArray,
  UseFieldArrayRemove,
  useForm,
} from "react-hook-form";
import { IRecordForm } from "../types/record";
import { ControlledDatePicker } from "./ControlledInputs/ControlledDatePicker";
import { ControlledTextField } from "./ControlledInputs/CotrolledTextField";
import * as Yup from "yup";
import { yupResolver } from "@hookform/resolvers/yup";
import { useAddRecord } from "../hooks/useAddRecord";
import { useBlock } from "../hooks/useBlock";

const formSchema = Yup.object({
  date: Yup.string().required("Date is required"),
  liquidity: Yup.number()
    .required("Liquidity is required")
    .min(0, "Liquidity must be greater or equal to 0"),
  stocks: Yup.array().of(
    Yup.object().shape({
      name: Yup.string().required("Name is required"),
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
      name: Yup.string().required("Name is required"),
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

interface IRecordStocksProps {
  fields: FieldArrayWithId<IRecordForm, "stocks", "id">[];
  control: Control<IRecordForm, object>;
  remove: UseFieldArrayRemove;
}

interface IRecordCryptosProps {
  fields: FieldArrayWithId<IRecordForm, "cryptos", "id">[];
  control: Control<IRecordForm, object>;
  remove: UseFieldArrayRemove;
}

export const Stocks = ({ control, fields, remove }: IRecordStocksProps) => {
  return (
    <>
      {fields.map((field, index) => (
        <div key={field.id}>
          <Grid
            container
            spacing={2}
            display="flex"
            justifyContent="center"
            alignItems="center"
          >
            <Grid item container md={11} spacing={2} xs={12}>
              <Grid item md={6} lg={3} xs={12}>
                <ControlledTextField
                  key={field.id}
                  control={control}
                  name={`stocks.${index}.name`}
                  label="Name"
                />
              </Grid>
              <Grid item md={6} lg={3} xs={12}>
                <ControlledTextField
                  key={field.id}
                  control={control}
                  name={`stocks.${index}.symbol`}
                  label="Symbol"
                />
              </Grid>
              <Grid item md={6} lg={3} xs={12}>
                <ControlledTextField
                  key={field.id}
                  control={control}
                  type="number"
                  name={`stocks.${index}.shares`}
                  label="Shares"
                />
              </Grid>
              <Grid item md={6} lg={3} xs={12}>
                <ControlledTextField
                  key={field.id}
                  type="number"
                  control={control}
                  name={`stocks.${index}.valuedAt`}
                  label="Valued At"
                />
              </Grid>
            </Grid>
            <Grid item md={1} xs={12} textAlign="center" key={field.id + "1"}>
              <IconButton onClick={() => remove(index)}>
                <Remove sx={{ margin: "0 auto" }} />
              </IconButton>
            </Grid>
          </Grid>
          <Divider sx={{ marginBottom: 2, mt: 2 }} />
        </div>
      ))}
    </>
  );
};

const Cryptos = ({ control, fields, remove }: IRecordCryptosProps) => {
  return (
    <>
      {fields.map((field, index) => (
        <div key={field.id}>
          <Grid
            container
            spacing={2}
            display="flex"
            justifyContent="center"
            alignItems="center"
          >
            <Grid item container md={11} spacing={2} xs={12}>
              <Grid item md={6} lg={3} xs={12}>
                <ControlledTextField
                  key={field.id}
                  control={control}
                  name={`cryptos.${index}.name`}
                  label="Name"
                />
              </Grid>
              <Grid item md={6} lg={3} xs={12}>
                <ControlledTextField
                  key={field.id}
                  control={control}
                  name={`cryptos.${index}.symbol`}
                  label="Symbol"
                />
              </Grid>
              <Grid item md={6} lg={3} xs={12}>
                <ControlledTextField
                  key={field.id}
                  control={control}
                  type="number"
                  name={`cryptos.${index}.coins`}
                  label="Coins"
                />
              </Grid>
              <Grid item md={6} lg={3} xs={12}>
                <ControlledTextField
                  key={field.id}
                  type="number"
                  control={control}
                  name={`cryptos.${index}.valuedAt`}
                  label="Valued At"
                />
              </Grid>
            </Grid>
            <Grid item md={1} xs={12} textAlign="center">
              <IconButton onClick={() => remove(index)}>
                <Remove sx={{ margin: "0 auto" }} />
              </IconButton>
            </Grid>
          </Grid>
          <Divider sx={{ marginBottom: 2, mt: 2 }} />
        </div>
      ))}
    </>
  );
};

export const RecordCard = () => {
  const { error, loading, addRecord } = useAddRecord();

  const {
    control,
    handleSubmit,
    formState: { errors, isDirty },
  } = useForm<IRecordForm>({
    defaultValues: {
      date: new Date(),
      stocks: [{ name: "", shares: 0, valuedAt: 0, symbol: "" }],
      cryptos: [{ name: "", coins: 0, valuedAt: 0, symbol: "" }],
    },
    resolver: yupResolver(formSchema),
  });
  useBlock(isDirty);

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
    stockAppend({ name: "", shares: 0, valuedAt: 0, symbol: "" });
  };
  const appendCrypto = () => {
    cryptoAppend({ name: "", coins: 0, valuedAt: 0, symbol: "" });
  };

  const onSubmit = (data: IRecordForm) => {
    data.date = new Date(data.date).toISOString();
    addRecord(data);
  };

  return (
    <Card>
      <CardContent>
        <form onSubmit={handleSubmit(onSubmit)}>
          <Typography variant="h4" mb={4}>
            Add Record
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
                  defaultValue={0}
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
