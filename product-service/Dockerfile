# Use the official Node.js image
FROM node:20-alpine


ENV ENVIRONMENT=stage
ENV PORT_API=3007
ENV PORT_GRCP=50057


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
EXPOSE 50057
EXPOSE 3007

# Command to run the application
CMD ["npm", "start"]
