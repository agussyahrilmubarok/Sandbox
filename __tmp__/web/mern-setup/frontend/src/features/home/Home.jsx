import React from "react";
import { Container, Typography, Box, Button } from "@mui/material";
import { Link } from "react-router";

const Home = () => {
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
          Welcome to Home Page
        </Typography>
        <Typography variant="body1" color="textSecondary" sx={{ mb: 3 }}>
          This is a simple homepage built with Material UI.
        </Typography>
        <Button
          variant="contained"
          color="primary"
          component={Link}
          to="/login"
        >
          Go to Sign In
        </Button>
      </Box>
    </Container>
  );
};

export default Home;
