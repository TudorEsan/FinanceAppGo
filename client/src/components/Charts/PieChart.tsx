import { Box } from "@mui/system";
import { ResponsivePie } from "@nivo/pie";
import { BasicTooltip } from "@nivo/tooltip";
import { useMobile } from "../../hooks/useMobile";
import { IDiversification } from "../../types/record";

interface IProps {
  data: IDiversification[];
}

export const MyPie = ({ data }: IProps) => {
  const isMobile = useMobile();
  return (
    <Box height={isMobile ? 250 : 300} width="100%" maxWidth={400}>
      {/* <p>tf is thiz</p> */}
      <ResponsivePie
        colors={{ scheme: "dark2" }}
        value="percent"
        id="symbol"
        data={data}
        margin={{ right: 90, left: 90 }}
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
