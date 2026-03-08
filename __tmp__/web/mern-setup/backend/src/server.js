import express from "express";
import mongoose from "mongoose";
import bodyParser from "body-parser";
import cookieParser from "cookie-parser";
import compress from "compression";
import helmet from "helmet";
import cors from "cors";

import config from "./config.js";
import userRoutes from "./routes/user.routes.js";
import authRoutes from "./routes/auth.routes.js";

const app = express();

// 🔹 Middleware
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(cookieParser());
app.use(compress());
app.use(helmet());
app.use(cors({
  origin: config.clientUrl, 
  credentials: true,  
}));
// 🔹 Routes
app.use("/", userRoutes);
app.use("/", authRoutes);

// 🔹 Connect to MongoDB & Start Server
(async () => {
  try {
    await mongoose.connect(config.mongoUri);
    console.log("✅ MongoDB Connected...");

    app.listen(config.port, () => {
      console.log(`🚀 Server running on port ${config.port}`);
    });
  } catch (err) {
    console.error("❌ MongoDB Connection Error:", err.message);
    process.exit(1);
  }
})();

export default app;
