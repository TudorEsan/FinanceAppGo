import plainAxios from "axios";
import { deleteAllCookies } from "./helpers/authHelper";

const inDevelopment = process.env.NODE_ENV === "development";

const axios = plainAxios.create({
  withCredentials: true,
  timeout: 5000,
});

axios.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    console.log(error);
    if (error.response.status === 401) {
      deleteAllCookies();
      window.location.reload();
    }
    return Promise.reject(error);
  }
);

export default axios;
