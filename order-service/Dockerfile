# Use the official Node.js image
FROM node:20-alpine

RUN apk add curl
# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application
COPY . .

# Expose the port
EXPOSE 50052
EXPOSE 3008

# Command to run the application
CMD ["npm", "start" ]
