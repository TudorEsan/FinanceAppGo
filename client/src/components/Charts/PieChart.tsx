import { Box } from "@mui/system";
import React from "react";
import { ResponsiveContainer, PieChart, Pie, Cell } from "recharts";
import { IDiversification } from "../../types/record";

const data = [
  { name: "network 1", value: 0.01 },
  { name: "network 3", value: 4 },
];

const COLORS = [
  "#fd7f6f",
  "#7eb0d5",
  "#b2e061",
  "#bd7ebe",
  "#ffb55a",
  "#ffee65",
  "#beb9db",
  "#fdcce5",
  "#8bd3c7",
];

interface IProps {
  data: IDiversification[];
}

export const MyPie = ({ data }: IProps) => {
  if (data.length === 0) {
    return null;
  }
  return (
    <ResponsiveContainer minWidth="400px" width={500} height={350}>
      <PieChart width={500} height={200}>
        <Pie
          data={data}
          cx="50%"
          cy="50%"
          outerRadius={100}
          fill="#8884d8"
          dataKey="percent"
          label={({
            cx,
            cy,
            midAngle,
            innerRadius,
            outerRadius,
            value,
            index,
          }) => {
            console.log("handling label?");
            const RADIAN = Math.PI / 180;
            // eslint-disable-next-line
            const radius = 25 + innerRadius + (outerRadius - innerRadius);
            // eslint-disable-next-line
            const x = cx + radius * Math.cos(-midAngle * RADIAN);
            // eslint-disable-next-line
            const y = cy + radius * Math.sin(-midAngle * RADIAN);

            return (
              <text
                x={x}
                y={y}
                fill={COLORS[index % COLORS.length]}
                textAnchor={x > cx ? "start" : "end"}
                dominantBaseline="central"
              >
                {data[index].symbol} {data[index].percent}%
              </text>
            );
          }}
        >
          {data.map((entry, index) => (
            <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
          ))}
        </Pie>
      </PieChart>
    </ResponsiveContainer>

    // </Box>
  );
};
