FROM golang:1.18.1-buster AS build

ENV GO111MODULE=auto

ARG KEY

WORKDIR /app

COPY ./ ./
RUN go mod download
RUN go mod tidy
RUN go build -ldflags "-X main.KEY=$KEY" -o /signedKeyServer

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /signedKeyServer /signedKeyServer

USER nonroot:nonroot

ENTRYPOINT ["/signedKeyServer"]
