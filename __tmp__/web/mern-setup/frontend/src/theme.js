import { createTheme } from "@mui/material/styles";

const theme = createTheme({
  typography: {
    fontFamily: "'Inter', 'Arial', sans-serif", // Font mirip dengan Notion
  },
  palette: {
    mode: "light", // Notion menggunakan tampilan terang secara default
    primary: {
      main: "#000000", // Hitam sebagai warna utama untuk teks utama
      contrastText: "#ffffff", // Kontras putih
    },
    secondary: {
      main: "#f5f5f5", // Abu-abu terang untuk elemen sekunder
      contrastText: "#000000",
    },
    background: {
      default: "#ffffff", // Latar belakang putih seperti Notion
      paper: "#f7f7f7", // Warna abu-abu lembut untuk kartu
    },
    text: {
      primary: "#37352f", // Warna teks utama mirip dengan Notion
      secondary: "#706f6c", // Warna teks sekunder agak redup
    },
    divider: "#e3e3e3", // Garis pembatas tipis seperti Notion
  },
  components: {
    MuiCard: {
      styleOverrides: {
        root: {
          borderRadius: "8px",
          boxShadow: "none",
          border: "1px solid #e3e3e3",
        },
      },
    },
    MuiButton: {
      styleOverrides: {
        root: {
          textTransform: "none",
          fontWeight: 500,
          borderRadius: "6px",
        },
      },
    },
  },
});

export default theme;
