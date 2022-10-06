import React from "react";
import { useAuth } from "../hooks/useAuth";
import { Navigate } from "react-router-dom";
import { collapseTextChangeRangesAcrossMultipleVersions } from "typescript";
import ResponsiveDrawer from "../components/ResponsiveDrawer";

interface ProtectedRouteProps {
  children: JSX.Element;
}

export const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ children }) => {
  const { isAuthenticated, emailValidated } = useAuth();
  if (!isAuthenticated) {
    return <Navigate to="/login" />;
  }

  if (isAuthenticated && !emailValidated) {
    return <Navigate to="/email-verification" />;
  }
  return (
    <>
      <ResponsiveDrawer>{children}</ResponsiveDrawer>
    </>
  );
};
