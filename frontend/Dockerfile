# Use an official Node runtime as a parent image
FROM node:latest

# Set the working directory in the container
WORKDIR /usr/src/app/

# Copy the rest of the application code
COPY . .

# Install project dependencies
RUN npm ci

# Build the application
RUN npm run build

# Expose the port that the app will run on
EXPOSE 3000

# Define the command to run your application
CMD ["npm", "run", "start"]