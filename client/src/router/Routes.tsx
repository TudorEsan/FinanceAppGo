import React from "react";
import { Route, Routes } from "react-router-dom";
import App from "../App";
import { Login } from "../pages/Login";
import { Protected } from "../pages/Home";
import { Register } from "../pages/Register";
import { ProtectedRoute } from "./ProtectedRoute";
import { AddRecord } from "../pages/AddRecord";
import { Records } from "../pages/Records";
import { Record } from "../pages/Record";

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

      <Route
        path="/records"
        element={
          <ProtectedRoute>
            <Records />
          </ProtectedRoute>
        }
      />
      <Route
        path="/records/:recordId"
        element={
          <ProtectedRoute>
            <Record />
          </ProtectedRoute>
        }
      />
      <Route
        path="/records/add"
        element={
          <ProtectedRoute>
            <AddRecord />
          </ProtectedRoute>
        }
      />
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Register />} />
    </Routes>
  );
};
