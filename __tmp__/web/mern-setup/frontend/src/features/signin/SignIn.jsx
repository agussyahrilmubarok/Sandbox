import React, { useContext, useState } from "react";
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
  Alert,
} from "@mui/material";
import { AuthContext } from "../../context/AuthContext";
import signInApi from "./signin.api.js";

const SignIn = () => {
  const navigate = useNavigate();
  const { login } = useContext(AuthContext);
  const [values, setValues] = useState({
    email: "",
    password: "",
    error: "",
  });

  const handleChange = (event) => {
    setValues({ ...values, [event.target.name]: event.target.value });
  };

  const handleSubmit = async () => {
    const result = await signInApi({
      email: values.email,
      password: values.password,
    });

    if (result.error) {
      setValues({ ...values, error: result.error });
    } else {
      login(result.token);
      navigate("/dashboard");
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
            Sign In
          </Typography>

          {values.error && (
            <Alert severity="error" sx={{ mb: 2 }}>
              {values.error}
            </Alert>
          )}

          <TextField
            label="Email"
            name="email"
            type="email"
            variant="outlined"
            fullWidth
            sx={{ my: 1 }}
            value={values.email}
            onChange={handleChange}
          />
          <TextField
            label="Password"
            name="password"
            type="password"
            variant="outlined"
            fullWidth
            sx={{ my: 1 }}
            value={values.password}
            onChange={handleChange}
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
            Sign In
          </Button>
          <Typography variant="body2" sx={{ mt: 2 }}>
            Do not have an account?{" "}
            <Link href="/register" color="primary" underline="hover">
              Sign Up
            </Link>
          </Typography>
        </CardActions>
      </Card>
    </Box>
  );
};

export default SignIn;
