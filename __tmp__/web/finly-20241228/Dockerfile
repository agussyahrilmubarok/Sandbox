# Use Node.js LTS as the base image
FROM node:22-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json files into the container
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy all files into the container
COPY . .

# Generate CSS
RUN npm run css

# Expose the application's port (adjust to match the port used by your Express.js app)
EXPOSE 8080

# Run the application
CMD ["npm", "start"]
