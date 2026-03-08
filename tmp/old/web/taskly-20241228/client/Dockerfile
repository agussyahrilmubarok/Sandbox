# Use Node.js version 22 as the base image
FROM node:22.11.0-alpine

# Set the working directory
WORKDIR /app

# Copy package.json and pnpm-lock.yaml files
COPY package.json pnpm-lock.yaml ./

# Install pnpm (the package manager you're using)
RUN npm install -g pnpm

# Install dependencies using pnpm
RUN pnpm install

# Copy all project files into the container
COPY . .

# Build the project for production
RUN pnpm build

# Expose the port for the Vite application (default: 5173)
EXPOSE 3000

# Command to run the Vite server
CMD ["pnpm", "preview", "--host", "0.0.0.0"]
