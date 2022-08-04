import React from "react";
import { IDiversification } from "../../types/record";
import { MyPie } from "../Charts/PieChart";

export const GeneralDiversification = () => {
  return (
    <div>
      <MyPie
        data={
          [
            { symbol: "aapl", percent: 12 },
            { symbol: "aapl", percent: 12 },
            { symbol: "aapl", percent: 12 },
            { symbol: "aapl", percent: 12 },
            { symbol: "aapl", percent: 12 },
            { symbol: "aapl", percent: 12 },
            { symbol: "aapl", percent: 12 },
            { symbol: "aapl", percent: 12 },
            { symbol: "aapl", percent: 12 },
            { symbol: "aapl", percent: 12 },
            { symbol: "aapl", percent: 12 },
          ] as IDiversification[]
        }
      />
    </div>
  );
};
