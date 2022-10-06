import { createContext, useState } from "react";
import { getCookie, isLoggedIn } from "../helpers/authHelper";
import { IClaims, IUser } from "../types/auth";
import decodeJwt from "jwt-decode";
interface IAuthContext {
  isAuthenticated: boolean;
  setIsAuthenticated: (isAuthenticated: boolean) => void;
  user: IUser | null;
  emailValidated: boolean;
  setEmailValidated: (emailValidated: boolean) => void;
}

export const AuthContext = createContext<IAuthContext>({
  isAuthenticated: isLoggedIn(),
  setIsAuthenticated: () => {},
  user: null,
  emailValidated: false,
  setEmailValidated: () => {},
});

interface AuthProviderProps {
  children: JSX.Element;
}

export const isEmailValidated = () => {
  try {
    const token = getCookie("token");
    if (!token) {
      return false;
    }
    const claims: IClaims = decodeJwt(token);
    return claims.EmailValidated;
  } catch (e) {
    console.error(e);
    return false;
  }
};

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(isLoggedIn());
  const [user, setUser] = useState<IUser | null>(null);
  const [emailValidated, setEmailValidatedState] = useState(isEmailValidated());
  const setEmailValidated = (val: boolean) => {
    setEmailValidatedState(val);
    // more logic if needed
  };

  return (
    <AuthContext.Provider
      value={{
        user,
        isAuthenticated,
        setIsAuthenticated,
        emailValidated,
        setEmailValidated,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};
