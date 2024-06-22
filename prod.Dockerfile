# Stage 1: Install Terraform on an Ubuntu base image
FROM ubuntu:latest as terraform

# Install dependencies for Terraform
RUN apt-get update \
    && apt-get install -y wget unzip

# Download and install Terraform
RUN wget https://releases.hashicorp.com/terraform/0.15.4/terraform_0.15.4_linux_amd64.zip \
    && unzip terraform_0.15.4_linux_amd64.zip \
    && mv terraform /usr/local/bin/ \
    && rm terraform_0.15.4_linux_amd64.zip


# Build Go app
FROM golang:1.22.1-alpine as builder

WORKDIR /usr/src/app

# RUN go install git hub.com/cosmtrek/air@latest
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 go build -o /tenant-onboarding-api

FROM alpine:latest

WORKDIR /root/

# Copy the Terraform binary from the first stage
COPY --from=terraform /usr/local/bin/terraform /usr/local/bin/

# Copy the built Go application from the builder stage
COPY --from=builder /tenant-management-api /tenant-management-api

# Expose the necessary port
EXPOSE 8085

# Set environment variables
ENV APP_PORT 8085
ENV TF_EXEC_PATH "/usr/bin/terraform"
ENV TF_WORKDIR "/root/terraform/"
ENV MODULE_NAME "tenant_management"
ENV GOOGLE_APPLICATION_CREDENTIALS "/creds.json"
ENV GOOGLE_PROJECT_ID = "myits-saas"
ENV INTEGRATED_MODE = "false"
ENV JWT_SECRET = "secret"
ENV DEPLOYMENT_QUEUE = "tenant_onboarding_deployment"
ENV DEPLOYMENT_QUEUE_SUBSCRIPTION = "tenant_onboarding_tenant_onboarding_deployment"

# Command to run the application
CMD ["/tenant-onboarding-api"]