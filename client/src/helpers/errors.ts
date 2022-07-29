export const getErrorMessage = (error: any): string => {
  return (
    (error?.response?.data?.message as string) ||
    (error?.response?.data as string) ||
    "Something went wrong"
  );
};
