import { Box } from "@mui/system";
import { ResponsivePie } from "@nivo/pie";
import { BasicTooltip } from "@nivo/tooltip";
import { IDiversification } from "../../types/record";

interface IProps {
  data: IDiversification[];
}

export const MyPie = ({ data }: IProps) => {
  console.log(data);
  return (
    <Box height={300} minWidth={400}>
      {/* <p>tf is thiz</p> */}
      <ResponsivePie
        colors={{ scheme: "dark2" }}
        value="percent"
        id="symbol"
        data={data}
        margin={{ right: 80, left: 80 }}
        innerRadius={0.5}
        padAngle={2}
        cornerRadius={1}
        sortByValue
        arcLabel={(d) => `${d.value}%`}
        arcLabelsSkipAngle={10}
        arcLinkLabelsTextColor={{
          from: "color",
          modifiers: [["brighter", 0.8]],
        }}
        arcLinkLabelsSkipAngle={10}
        theme={{
          tooltip: {
            container: {
              color: "black",
            },
          },
        }}
        // ={(d) => `${d}%`}
        tooltip={(d) => {
          console.log(d);
          return (
            <BasicTooltip
              id={d.datum.id}
              color={d.datum.color}
              value={d.datum.value + "%"}
              enableChip
            />
          );
        }}

        // activeOuterRadiusOffset={8}
      />
    </Box>
  );
};
