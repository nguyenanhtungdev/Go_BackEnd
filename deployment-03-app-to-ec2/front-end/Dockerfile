# Base image
FROM node:18

# Set working directory
WORKDIR /app

# Copy package.json and install dependencies
COPY package.json package-lock.json ./
RUN npm install

# Copy source code
COPY . /app

# Expose React dev server port
EXPOSE 3000

# Start React development server
CMD ["npm", "start"]