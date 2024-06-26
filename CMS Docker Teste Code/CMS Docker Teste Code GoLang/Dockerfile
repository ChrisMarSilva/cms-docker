FROM golang:1.18 AS builder
WORKDIR /app
# ENV GO111MODULE=off CGO_ENABLED=1 GOOS=linux OARCH=amd64
RUN go env -w GOPROXY=direct GOFLAGS="-insecure"
RUN apt-get update && apt-get install -y ca-certificates openssl
ARG cert_location=/usr/local/share/ca-certificates
RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
RUN update-ca-certificates
#RUN go install github.com/cespare/reflex@latest
# RUN apk --update add --no-cache ca-certificates openssl git tzdata
# COPY your/cert/path /usr/local/share/ca-certificates/your-cert-name
# RUN update-ca-certificates
# RUN apk --update add --no-cache ca-certificates openssl git tzdata && update-ca-certificates
# RUN apk add --no-cache bash
RUN go clean --modcache
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ${MAIN_PATH}
#CMD reflex -g '*.go' go run main.go --start-service

FROM scratch
# RUN apk --no-cache add ca-certificates 
# RUN apk add --no-cache git make musl-dev go
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
COPY --from=builder /app/main .
# ENV GOROOT /usr/lib/go
# ENV GOPATH /go
# ENV PATH /go/bin:$PATH
# RUN mkdir -p
CMD ["./main"] 





# ## Build
# FROM golang:1.16-buster AS build
# WORKDIR /app
# COPY go.mod .
# COPY go.sum .
# RUN go mod download
# COPY *.go ./
# RUN go build -o /docker-gs-ping

# ## Deploy
# FROM gcr.io/distroless/base-debian10
# WORKDIR /
# COPY --from=build /docker-gs-ping /docker-gs-ping
# EXPOSE 8080
# USER nonroot:nonroot
# ENTRYPOINT ["/docker-gs-ping"]



# FROM golang:1.13-alpine as builder
# ENV GOPATH="$HOME/go"
# RUN apk --no-cache add git
# WORKDIR $GOPATH/src
# COPY . $GOPATH/src
# RUN go get -d -v golang.org/x/net/html
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app consumer/*.go

# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# WORKDIR /root/
# COPY --from=builder $HOME/go/src/app .
# CMD ["./app"]  


# FROM golang:1.16.5 AS development
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN go install github.com/cespare/reflex@latest
# EXPOSE 4000
# CMD reflex -g '*.go' go run main.go --start-service

# FROM golang:1.16.5 AS builder
# ENV GOOS linux
# ENV CGO_ENABLED 0
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN go build -o app

# FROM alpine:latest AS production
# RUN apk add --no-cache ca-certificates
# COPY --from=builder app .
# EXPOSE 4000
# CMD ./app







# RUN go mod download && go mod verify


# FROM golang:latest as builder
# RUN mkdir /app
# WORKDIR /app
# COPY . ./
# RUN make test
# ARG version=dev
# RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -ldflags "-X main.version=$version" -o kn-be-api -v ./cmd/api/main.go

# FROM alpine
# COPY --from=builder /app/kn-be-api /



# FROM golang:latest as build
# WORKDIR /build
# ADD . .
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o app
# FROM scratch
# COPY --from=build /build/app /app
# ENTRYPOINT ["/app"]


# # Build stage
# FROM golang:1.16-alpine3.13 AS builder
# WORKDIR /app
# COPY . .
# RUN go build -o main main.go
# RUN apk --no-cache add curl
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# # Run stage
# FROM alpine:3.13
# WORKDIR /app
# COPY --from=builder /app/main .
# COPY --from=builder /app/migrate.linux-amd64 ./migrate
# COPY app.env .
# COPY start.sh .
# COPY wait-for.sh .
# COPY db/migration ./migration

# EXPOSE 8080
# CMD [ "/app/main" ]
# ENTRYPOINT [ "/app/start.sh" ]



# # Build stage
# FROM golang:1.16-alpine3.13 AS builder
# WORKDIR /app
# COPY . .
# RUN go build -o main main.go
# RUN apk --no-cache add curl
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# # Run stage
# FROM alpine:3.13
# WORKDIR /app
# COPY --from=builder /app/main .
# COPY --from=builder /app/migrate.linux-amd64 ./migrate
# COPY app.env .
# COPY start.sh .
# COPY wait-for.sh .
# COPY db/migration ./migration

# EXPOSE 8080
# CMD [ "/app/main" ]
# ENTRYPOINT [ "/app/start.sh" ]