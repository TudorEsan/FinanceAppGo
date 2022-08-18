import { fabClasses } from "@mui/material";
import { Box } from "@mui/system";
import { ResponsivePie } from "@nivo/pie";
import { BasicTooltip } from "@nivo/tooltip";
import { createFalse } from "typescript";
import { useMobile } from "../../hooks/useMobile";
import { IDiversification } from "../../types/record";

interface IProps {
  data: IDiversification[];
  enableArcLinkLabels?: boolean;
  showLegend?: boolean;
}

export const MyPie = ({
  data,
  showLegend = false,
  enableArcLinkLabels = true,
}: IProps) => {
  const isMobile = useMobile();
  return (
    <Box
      height={isMobile ? 250 : 300}
      // width="100%"
      maxWidth={400}
      margin="auto"
    >
      {/* <p>tf is thiz</p> */}
      <ResponsivePie
        colors={{ scheme: "dark2" }}
        value="percent"
        legends={
          showLegend
            ? [
                {
                  anchor: "bottom",
                  direction: "row",
                  translateX: 0,
                  translateY: 50,
                  itemWidth: 75,
                  itemHeight: 20,
                  itemsSpacing: 5,
                  symbolSize: 20,
                  itemDirection: "left-to-right",
                  itemTextColor: "white",
                },
              ]
            : []
        }
        id="symbol"
        data={data}
        margin={
          enableArcLinkLabels
            ? { right: 90, left: 90, bottom: 90 }
            : { right: 0, left: 0, bottom: 90 }
        }
        innerRadius={0.5}
        padAngle={2}
        cornerRadius={1}
        sortByValue
        enableArcLinkLabels={enableArcLinkLabels}
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
          labels: {
            text: {
              // color: "white",
              fontWeight: "bold",
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
