FROM --platform=${BUILDPLATFORM} golang:1.22-alpine3.19 AS builder

RUN apk add git

WORKDIR /go/src/app
COPY . .

ARG TARGETOS TARGETARCH TARGETVARIANT

ENV CGO_ENABLED=0
RUN go get \
    && go mod download \
    && GOOS=${TARGETOS} GOARCH=${TARGETARCH} GOARM=${TARGETVARIANT#"v"} go build -a -o rtsp-streamer

FROM golang:1.22

WORKDIR /app

COPY --from=builder /go/src/app/rtsp-streamer /app/
COPY --from=builder /go/src/app/web /app/web
COPY --from=builder /go/src/app/config.json /app/config.json
COPY --from=builder /go/src/app/.env /app/.env

ENV GO111MODULE="on"
ENV GIN_MODE="release"

CMD ["./rtsp-streamer"]