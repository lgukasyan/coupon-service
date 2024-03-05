FROM golang:latest

# Workdir
WORKDIR /app

# Copy module
COPY go.mod .
COPY go.sum .

# Install modules
RUN go mod download

# Copy project
COPY . .

# Make build (coupon_service)
RUN go build .

# Expose
EXPOSE 8080

# Execute
CMD ["./coupon_service", "serve"]