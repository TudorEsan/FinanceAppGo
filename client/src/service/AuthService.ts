import axios from "../axiosConfig";

export const signIn = (username: string, password: string) => {
  return axios.post("/auth/login", { username, password });
};

export const signUp = (data: RegisterInput) => {
  return axios.post("/auth/signup", data);
};
