import React from "react";
import { Route, Routes } from "react-router-dom";
import App from "../App";
import { Login } from "../pages/Login";
import { Protected } from "../pages/Protected";
import { Register } from "../pages/Register";
import { ProtectedRoute } from "./ProtectedRoute";

export const AppRoutes = () => {
  return (
    <Routes>
      <Route
        path="/"
        element={
          <ProtectedRoute>
            <Protected />
          </ProtectedRoute>
        }
      />
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Register />} />
    </Routes>
  );
};
