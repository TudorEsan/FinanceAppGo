// @ts-nocheck
import { createTheme, ThemeProvider } from "@mui/material";
import React from "react";
import { AuthContext, AuthProvider } from "./context/AuthProvider";
import { AppRoutes } from "./router/Routes";
import "./index.css";

function App() {
  const theme = createTheme({
    palette: {
      primary: {
        main: "#3661EB",
      },
    },
    components: {
      MuiButton: {
        styleOverrides: {
          root: {
            borderRadius: "1000px",
          },
        },
      },
    },
  });
  return (
    <div>
      <AuthProvider>
        <ThemeProvider theme={theme}>
          <AppRoutes />
        </ThemeProvider>
      </AuthProvider>
    </div>
  );
}

export default App;
