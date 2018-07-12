FROM arm32v6/golang:1.10-alpine as builder

RUN apk add --no-cache git ca-certificates

WORKDIR /go/src/github.com/Eiryyy/agrec
RUN go get -u github.com/golang/dep/...

COPY . .

RUN dep ensure
RUN go install

FROM arm32v6/alpine

RUN apk add --no-cache ffmpeg rtmpdump
ENV GOROOT /go
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /go/lib/time/zoneinfo.zip
WORKDIR /root
COPY --from=builder /go/bin/agrec .
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/github.com/Eiryyy/agrec/programs.toml .
CMD ["./agrec"]
