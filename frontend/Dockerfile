# Use the official Bun image as the base image
FROM docker.io/oven/bun:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and bun.lockb to the working directory
COPY package.json bun.lock ./

# Install the dependencies
RUN bun install --frozen-lockfile

# Copy the source code to the working directory
COPY . .

# Build the application
RUN bun run build

EXPOSE 3000
CMD ["bun", "run", "start"]

#