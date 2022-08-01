// @ts-nocheck
import { createTheme, CssBaseline, ThemeProvider } from "@mui/material";
import React from "react";
import { AuthContext, AuthProvider } from "./context/AuthProvider";
import { AppRoutes } from "./router/Routes";
import "./index.css";
import axios from "axios";
import { deleteAllCookies } from "./helpers/authHelper";

axios.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    if (error.response.status === 401) {
      deleteAllCookies();
      window.location.reload();
    }
    return Promise.reject(error);
  }
);

function App() {
  const { palette } = createTheme();
  const theme = createTheme({
    palette: {
      background: {
        default: "#inherit",
      },
      mode: "dark",
      primary: {
        main: "#17C6B1",
      },
      secondary: {
        main: "#72E8C9",
      },
      mycolor: { main: "red" },
    },
    components: {
      MuiButton: {
        styleOverrides: {
          root: {
            // borderRadius: "1000px",
          },
        },
      },
      MuiCard: {
        styleOverrides: {
          root: {
            backgroundColor: "#161616",
            borderRadius: "20px",
          },
        },
      },
    },
  });
  return (
    <div>
      <AuthProvider>
        <ThemeProvider theme={theme}>
          <CssBaseline />
          <AppRoutes />
        </ThemeProvider>
      </AuthProvider>
    </div>
  );
}

export default App;