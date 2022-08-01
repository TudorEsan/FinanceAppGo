import MomentUtils from "@date-io/moment";
import { Controller } from "react-hook-form";
import { TextField } from "@mui/material";
import { DatePicker, LocalizationProvider } from "@mui/x-date-pickers";


interface ControlledDatePickerProps {
    name: string;
    label: string;
    defaultValue?: any;
    control: any;
    rest?: any;
}


export function ControlledDatePicker({
  name,
  label,
  defaultValue,
  control,
  rest,
}: ControlledDatePickerProps) {
  return (
    <Controller
      control={control}
      name={name}
      defaultValue={defaultValue || null}
      render={({ field: { onChange, value } }) => (
        <LocalizationProvider dateAdapter={MomentUtils}>
          <DatePicker
            label={label}
            value={value}
            onChange={onChange}
            renderInput={(params: any) => (
              <TextField fullWidth sx={{ minWidth: 150 }} {...params} />
            )}
          />
        </LocalizationProvider>
      )}
    />
  );
}
