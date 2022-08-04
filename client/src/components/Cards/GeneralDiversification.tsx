import React from "react";
import { IDiversification } from "../../types/record";
import { MyPie } from "../Charts/PieChart";

export const GeneralDiversification = () => {
  return (
    <MyPie
      data={
        [
          { symbol: "1", percent: 40 },
          { symbol: "2", percent: 20 },
          { symbol: "3", percent: 20 },
          { symbol: "4", percent: 5 },
          { symbol: "5", percent: 10 },
          { symbol: "6", percent: 1 },
        ] as IDiversification[]
      }
    />
  );
};
