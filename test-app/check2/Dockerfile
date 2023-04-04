# Use Node.js 14.x as the base image
FROM node:14

# Set the working directory to /app
WORKDIR /app

# Copy the package.json and package-lock.json files to the container
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code to the container
COPY . .

# Expose port 3000 for the server
EXPOSE 3000

# Set the command to start the server
CMD [ "npm", "start" ]
