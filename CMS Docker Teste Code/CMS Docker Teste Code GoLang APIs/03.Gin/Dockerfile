# syntax=docker/dockerfile:1

## Build
FROM golang:1.18-buster AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /main

## Deploy
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /main /main
#EXPOSE 8080
USER nonroot:nonroot
CMD ["./main"] 



# ## Build
# #FROM golang:1.16 AS builder
# FROM golang:latest AS builder
# WORKDIR /app
# RUN go env -w GOPROXY=direct GOFLAGS="-insecure"
# RUN apt-get update && apt-get install -y ca-certificates openssl
# ARG cert_location=/usr/local/share/ca-certificates
# RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
# RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
# RUN update-ca-certificates
# #RUN go version && go get -v golang.org/x/vgo
# RUN go clean --modcache
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# #COPY *.go ./
# RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ${MAIN_PATH}

# ## Deploy
# #FROM scratch
# FROM gcr.io/distroless/base-debian10
# ENV GIN_MODE=release
# COPY --from=builder /go/src/github.com/alextanhongpin/go-gin-starter/app .
# COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=build /docker-gs-ping /docker-gs-ping
# WORKDIR /app
# COPY --from=builder /app/main .
# EXPOSE 3001
# CMD ["./main"] 
# #ENTRYPOINT /app/main --port 3001
