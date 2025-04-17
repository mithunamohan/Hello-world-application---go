# Use official Node.js image as base
FROM node:lts

# Set working directory inside container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json
COPY ["package.json", "package-lock.json", "./"]

# Install dependencies for React app
RUN npm install --silent

# Copy the rest of the React app
COPY . .

# Expose the port React runs on
EXPOSE 5173

# Run the React app
CMD ["npm", "run", "dev", "--", "--host"]
