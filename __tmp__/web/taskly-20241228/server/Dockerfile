# Use Node.js version 22 as the base image
FROM node:22.11.0-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json to the container
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy all the application files to the container
COPY . .

# Expose the application port (default: 8000 from .env)
EXPOSE 8000

# Define the environment (e.g., "development" or "production")
ENV NODE_ENV=production

# Command to run the application
CMD ["npm", "run", "prod"]
