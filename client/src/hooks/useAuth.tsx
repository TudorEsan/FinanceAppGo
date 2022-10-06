import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { AuthContext, isEmailValidated } from "../context/AuthProvider";
import { deleteAllCookies, getCookie, isLoggedIn } from "../helpers/authHelper";
import { getErrorMessage } from "../helpers/errors";
import { signIn, signUp } from "../service/AuthService";
import jwt_decode from "jwt-decode";
import { IClaims, RegisterInput } from "../types/auth";

export const useAuth = () => {
  const {
    isAuthenticated,
    setIsAuthenticated,
    setEmailValidated,
    emailValidated,
  } = React.useContext(AuthContext);
  const [isLoading, setIsLoading] = React.useState(false);
  const [error, setError] = React.useState<null | string>(null);
  const navigate = useNavigate();

  const login = async (username: string, password: string) => {
    setIsLoading(true);
    setError(null);
    try {
      const user = await signIn(username, password);
      console.log(user);
      setEmailValidated(user.emailValidated);
      setIsAuthenticated(true);
      navigate("/");
    } catch (error: any) {
      setError(getErrorMessage(error));
    }
    setIsLoading(false);
  };

  const register = async (data: RegisterInput) => {
    setIsLoading(true);
    try {
      const user = await signUp(data);
      console.log(user);
      setIsAuthenticated(true);
      navigate("/email-verification");
    } catch (error: any) {
      setError(getErrorMessage(error));
    }
    setIsLoading(false);
  };

  const logout = () => {
    setIsAuthenticated(false);
    navigate("/login");
    window.localStorage.removeItem("emailValidated");
    deleteAllCookies();
  };

  useEffect(() => {
    console.log(getCookie("token"));
  }, []);

  return {
    emailValidated,
    logout,
    isAuthenticated,
    login,
    isLoading,
    error,
    register,
  };
};
