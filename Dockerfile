FROM ubuntu:latest as terraform

# Install dependencies for Terraform
RUN apt-get update \
    && apt-get install -y wget unzip

# Download and install Terraform
RUN wget https://releases.hashicorp.com/terraform/0.15.4/terraform_0.15.4_linux_amd64.zip \
    && unzip terraform_0.15.4_linux_amd64.zip \
    && mv terraform /usr/bin/terraform \
    && rm terraform_0.15.4_linux_amd64.zip

FROM golang:1.22.1-alpine

# RUN apk update && apk add wget

# RUN apk --no-cache add zip

# RUN apk --no-cache add git

# # Download and install Terraform
# RUN wget https://releases.hashicorp.com/terraform/1.8.4/terraform_1.8.4_linux_amd64.zip \
#     && unzip terraform_1.8.4_linux_amd64.zip \
#     && rm terraform_1.8.4_linux_amd64.zip \
#     && mv terraform /usr/bin/terraform \
#     && chmod +x /usr/bin/terraform


COPY --from=terraform /usr/bin/terraform /usr/bin/terraform

WORKDIR /usr/src/app


# RUN go install git hub.com/cosmtrek/air@latest
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go install github.com/air-verse/air@latest

# WORKDIR /usr/src/app/cmd/server

ENTRYPOINT ["air", "-c", ".air.toml"]

# RUN go mod download

# CMD ["air", "-c", ".air.toml"]
