FROM golang:1.17.2 as development

# Setup workdir
WORKDIR /app

# Copy source code
COPY . .

# Install dependencies
RUN go mod download

# Install Reflex for development -> Enables live reload
RUN go install github.com/cespare/reflex@latest

# Start app
CMD reflex -g '*.go' go run main.go --start-service