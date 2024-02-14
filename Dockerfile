# Start with a Node.js image for building frontend
FROM node:21-alpine as build-stage

WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the frontend code
COPY . .

# Build the frontend
RUN npm run tailwind && npm run esbuild

# Now start a new stage with golang base image
FROM golang:latest

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go mod download

# Copy the Go source code
COPY . .

# Copy the built frontend files from the build-stage
COPY --from=build-stage /app/public/out ./public/out

# Build the Go app
RUN templ generate
RUN go build -o ./bin ./app

EXPOSE 3000

CMD ["./bin/app"]