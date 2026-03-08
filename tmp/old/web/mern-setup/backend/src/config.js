const config = {
    name: process.env.APP_NAME || "server",
    port: process.env.APP_PORT || 3000,
    env: process.env.APP_ENV || "development",
    jwtSecret: process.env.JWT_SECRET || "the_secret_key",
    mongoUri: process.env.MONGODB_URI || "mongodb://dev:my-secret-pw@localhost:27017/mydatabase?authSource=mydatabase",
    clientUrl: process.env.CLIENT_URL || "http://localhost:3000"
};

export default config;
