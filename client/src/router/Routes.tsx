import React from "react";
import { Route, Routes } from "react-router-dom";
import { AddRecord, Dashboard, EditRecord, Login, Record, Records, Register } from "../pages";
import { CloneRecord } from "../pages/Record/CloneRecord";
import { ProtectedRoute } from "./ProtectedRoute";

export const AppRoutes = () => {
  return (
    <Routes>
      <Route
        path="/"
        element={
          <ProtectedRoute>
            <Dashboard />
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
        path="/records/:recordId/clone"
        element={
          <ProtectedRoute>
            <CloneRecord />
          </ProtectedRoute>
        }
      />
      <Route
        path="/records/:recordId/edit"
        element={
          <ProtectedRoute>
            <EditRecord />
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
