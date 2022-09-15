import { Box, Typography } from "@mui/material";
import { Datum, ResponsiveLine, Serie } from "@nivo/line";
import { formatDate } from "../../helpers/date";
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
          xFormat={(d) => String(d).replaceAll("/", "-")}
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
            format: (d) => (isMobile ? "" : null),
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
          pointSize={12}
          pointColor="rgba(255,255,255,0.8)"
          pointBorderWidth={3}
          // point={0}
          lineWidth={4}
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
            grid: {
              line: {
                stroke: "rgba(255, 255, 255, 0.2)",
                strokeWidth: 1,
              },
            },
          }}
        />
      </Box>
    </>
  );
};
