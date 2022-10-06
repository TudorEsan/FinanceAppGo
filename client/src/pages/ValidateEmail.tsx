import { Box, Button, Card, Link, Typography } from "@mui/material";
import React from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";

export const ValidateEmail = () => {
  const navigate = useNavigate();
  const { emailValidated } = useAuth();
  
  console.log(emailValidated);
  if (emailValidated) {
    navigate("/");
  }

  return (
    <Box
      p={2}
      sx={{
        maxWidth: "600px",
        position: "absolute",
        top: "50%",
        left: "50%",
        transform: "translate(-50%, -50%)",
        width: "100%",
      }}
    >
      <Card
        sx={{
          p: 2,
          display: "flex",
          justifyContent: "center",
          flexDirection: "column",
        }}
      >
        <Typography variant="h4" textAlign="center">
          Account not verified.
        </Typography>
        <Typography textAlign="center">
          Please check your email for a verification link. If you did not
          receive an email, please check your spam folder.
        </Typography>
        {/* <Button sx={{ margin: "auto" }}>Resend verification email</Button> */}
        <Link
          component="button"
          textAlign="center"
          underline="hover"
          onClick={() => navigate("/login")}
        >
          Login
        </Link>
        <Link
          component="button"
          textAlign="center"
          underline="hover"
          onClick={() => navigate("/register")}
        >
          Register
        </Link>
      </Card>
    </Box>
  );
};
