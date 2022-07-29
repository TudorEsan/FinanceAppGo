import React from "react";
import { useAuth } from "../hooks/useAuth";
import { Navigate } from "react-router-dom";

interface ProtectedRouteProps {
  children: JSX.Element;
}

export const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ children}) => {
  const { isAuthenticated } = useAuth();
  console.log("wtf");
  if (!isAuthenticated) {
    return <Navigate to="/login" />;
  }
  return children;
};
