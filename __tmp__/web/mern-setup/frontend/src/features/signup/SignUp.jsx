import React, { useState } from "react";
import { useNavigate } from "react-router";
import {
  Button,
  Card,
  CardActions,
  CardContent,
  TextField,
  Typography,
  Box,
  Link,
} from "@mui/material";
import signUpApi from "./signup.api.js";

const SignUp = () => {
  const navigate = useNavigate();

  const [values, setValues] = useState({
    name: "",
    email: "",
    password: "",
    error: "",
  });

  const handleChange = (event) => {
    setValues({ ...values, [event.target.name]: event.target.value });
  };

  const handleSubmit = async () => {
    const result = await signUpApi({
      name: values.name,
      email: values.email,
      password: values.password,
    });

    if (result.error) {
      setValues({ ...values, error: result.error });
    } else {
      console.log("Register successful:", result);
      navigate("/login");
    }
  };

  return (
    <Box
      sx={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "100vh",
      }}
    >
      <Card sx={{ maxWidth: 400, p: 3, textAlign: "center" }}>
        <CardContent>
          <Typography variant="h5" color="primary" gutterBottom>
            Sign Up
          </Typography>

          {values.error && (
            <Alert severity="error" sx={{ mb: 2 }}>
              {values.error}
            </Alert>
          )}

          <TextField
            label="Name"
            type="text"
            variant="outlined"
            fullWidth
            name="name"
            value={values.name}
            onChange={handleChange}
            sx={{ my: 1 }}
          />
          <TextField
            label="Email"
            type="email"
            variant="outlined"
            fullWidth
            name="email"
            value={values.email}
            onChange={handleChange}
            sx={{ my: 1 }}
          />
          <TextField
            label="Password"
            type="password"
            variant="outlined"
            fullWidth
            name="password"
            value={values.password}
            onChange={handleChange}
            sx={{ my: 1 }}
          />
        </CardContent>
        <CardActions sx={{ flexDirection: "column", alignItems: "center" }}>
          <Button
            variant="contained"
            color="primary"
            fullWidth
            sx={{ minHeight: 48 }}
            onClick={handleSubmit}
          >
            Sign Up
          </Button>
          <Typography variant="body2" sx={{ mt: 2 }}>
            Already have an account?{" "}
            <Link href="/login" color="primary" underline="hover">
              Sign In
            </Link>
          </Typography>
        </CardActions>
      </Card>
    </Box>
  );
};

export default SignUp;
