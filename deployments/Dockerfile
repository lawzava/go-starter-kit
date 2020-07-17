# Compile app binary
FROM golang:latest as build-env

WORKDIR goapp
ENV GO111MODULE=on

COPY go.* .

COPY cmd cmd
COPY internal internal
COPY pkg pkg

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -mod vendor -installsuffix cgo ./cmd/...

# Run app in scratch
FROM scratch

COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=build-env /go/bin/* /app/

ENV PATH="/app:${PATH}"
