import { Box, Typography } from "@mui/material";
import { Datum, ResponsiveLine, Serie } from "@nivo/line";
import { useMobile } from "../../hooks/useMobile";
import { ILinear } from "../../types/overview";

interface IProps {
  data: Datum[];
  id: string;
}

export const LineChart = ({ id, data }: IProps) => {
  const isMobile = useMobile();
  const serie = { id, data } as Serie;

  return (
    <>
      <Box height={isMobile ? 300 : 350} margin="auto">
        <ResponsiveLine
          colors="#17C6B1"
          // colors={{ scheme: "accent" }}
          yFormat={(d) => d + "$"}
          data={[serie]}
          margin={{ top: 10, right: 70, bottom: 50, left: 70 }}
          xScale={{ type: "point" }}
          yScale={{
            type: "linear",
            min: "auto",
            max: "auto",
            stacked: true,
            reverse: false,
          }}
          axisTop={null}
          axisRight={null}
          axisBottom={{
            // orient: "bottom",
            tickSize: 5,
            tickPadding: 5,
            tickRotation: 18,
            legend: "date",

            legendOffset: 36,
            legendPosition: "middle",
          }}
          axisLeft={{
            // orient: "left",
            tickSize: 5,
            tickPadding: 5,
            tickRotation: 0,
            legend: "total $",
            legendOffset: -50,
            legendPosition: "middle",
          }}
          pointSize={10}
          pointColor={{ theme: "background" }}
          pointBorderWidth={2}
          pointBorderColor={{ from: "serieColor" }}
          pointLabelYOffset={-12}
          useMesh={true}
          theme={{
            crosshair: {
              line: {
                stroke: "rgba(255, 255, 255, 0.75)",
                strokeWidth: 1,
              },
            },
            tooltip: {
              container: {
                color: "black",
              },
            },
            axis: {
              legend: {
                text: {
                  fontSize: "16px",
                  fill: "white",
                },
              },
              ticks: {
                text: {
                  fontSize: "14px",
                  fill: "white",
                },
              },
            },
          }}
          // legends={[
          //   {
          //     anchor: "bottom-right",
          //     itemTextColor: "white",
          //     direction: "column",
          //     justify: false,
          //     translateX: 100,
          //     translateY: 0,
          //     itemsSpacing: 0,
          //     itemDirection: "left-to-right",
          //     itemWidth: 80,
          //     itemHeight: 20,
          //     itemOpacity: 0.75,
          //     symbolSize: 12,
          //     symbolShape: "circle",
          //     symbolBorderColor: "rgba(0, 0, 0, .5)",
          //     effects: [
          //       {
          //         on: "hover",
          //         style: {
          //           itemBackground: "rgba(0, 0, 0, .03)",
          //           itemOpacity: 1,
          //         },
          //       },
          //     ],
          //   },
          // ]}
        />
      </Box>
    </>
  );
};
