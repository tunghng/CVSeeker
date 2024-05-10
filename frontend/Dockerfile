# Use an Alpine-based Node.js image for a smaller final image
FROM node:16-alpine

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json (or yarn.lock)
COPY package.json package-lock.json ./

# Install all dependencies, including 'devDependencies' since 'vite' is needed
RUN npm install

# Copy the rest of your application code into the container
COPY . .

# Expose the port that Vite uses
EXPOSE 5173

# Command to run the development server
CMD ["npm", "run", "dev"]
