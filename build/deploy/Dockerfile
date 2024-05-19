# syntax=docker/dockerfile:1.4

FROM golang:1.22-alpine as builder

# Set working directory.
WORKDIR /compiler

# Copy dependency locks so we can cache.
COPY go.mod go.sum ./

# Get all of our dependencies.
RUN go mod download

# Copy all of our remaining application.
COPY . .

# Build our application.
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/main.go

FROM alpine:latest as certs
RUN apk --update add ca-certificates

# Use 'scratch' image for super-mini build.
FROM scratch AS prod

# Set environment variables.
ARG PORT=8080
ARG APP_ENV=development

ENV PORT=$PORT
ENV APP_ENV=$APP_ENV

ENV PATH=/bin
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Set working directory for this stage.
WORKDIR /production

# Copy our compiled executable from the last stage.
COPY --from=builder --link /compiler/main .
COPY --from=builder --link /compiler/cmd/app/config/config.$APP_ENV.yaml ./cmd/app/config/config.$APP_ENV.yaml

# Run application and expose port.
EXPOSE $PORT

CMD ["./main"]