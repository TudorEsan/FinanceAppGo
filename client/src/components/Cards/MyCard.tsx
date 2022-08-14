import { Card, CardActionArea, CardContent } from "@mui/material";
import React from "react";

interface IMyCardProps {
  children: React.ReactNode;
  [x: string]: any;
}

export const MyCard = ({ children, rest }: IMyCardProps) => {
  return (
    <Card elevation={7} sx={{ mt: 2, mb: 2 }}>
      <CardContent>{children}</CardContent>
    </Card>
  );
};
