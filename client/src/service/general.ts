export const serverUrl = () => {
  return process.env.NODE_ENV === "development"
    ? "http://localhost:8080/api"
    : process.env.SERVER_URL + "/api";
};
