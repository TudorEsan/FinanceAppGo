import axios from "../axiosConfig";
import {  IUser, RegisterInput } from "../types/auth";
import { serverUrl } from "./general";

export const signIn = async (username: string, password: string) => {
  return (await axios.post(serverUrl() + "/auth/login", { username, password })).data as IUser;
} ;

export const signUp = async (data: RegisterInput) => {
  return (await axios.post(serverUrl() + "/auth/signup", data)).data as IUser;
};
