import React, { useContext } from "react";
import { useNavigate } from "react-router";
import { Container, Typography, Box, Button } from "@mui/material";
import { AuthContext } from "../../context/AuthContext";

const Dashboard = () => {
  const navigate = useNavigate();
  const { logout } = useContext(AuthContext);

  const handleLogout = () => {
    logout();
    navigate("/login");
  };

  return (
    <Container maxWidth="md">
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
          alignItems: "center",
          height: "100vh",
          textAlign: "center",
        }}
      >
        <Typography variant="h3" color="primary" gutterBottom>
          Welcome to Dashboard Page
        </Typography>
        <Typography variant="body1" color="textSecondary" sx={{ mb: 3 }}>
          This is a simple dashboard built with Material UI.
        </Typography>
        <Button onClick={handleLogout} variant="contained">
          Logout
        </Button>
      </Box>
    </Container>
  );
};

export default Dashboard;
