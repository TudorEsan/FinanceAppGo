import { Card, CardActionArea, CardContent } from "@mui/material";
import React from "react";

interface IMyCardProps {
  children: React.ReactNode;
}

export const MyCard = ({ children }: IMyCardProps) => {
  return (
    <Card elevation={7}>
      <CardContent>{children}</CardContent>
    </Card>
  );
};
