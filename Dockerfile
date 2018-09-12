# build stage
FROM golang:1.9.4-stretch as builder

ENV SERVICE_NAME url_minifier
ENV PKG github.com/lillilli/url_minificator

RUN mkdir -p /go/src/${PKG}
WORKDIR /go/src/${PKG}

COPY . .

RUN cd src/ && CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o ${SERVICE_NAME}

FROM alpine:3.7
RUN apk --no-cache add ca-certificates

WORKDIR /root/

ENV SERVICE_NAME url_minifier
ENV PKG github.com/lillilli/url_minificator

COPY --from=builder /go/src/${PKG}/src/${SERVICE_NAME} .

ENTRYPOINT ["./url_minifier"]
