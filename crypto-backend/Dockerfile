FROM golang:1.16.4-buster AS builder

ARG VERSION=dev

WORKDIR /go/src/app
# was COPY main.go .
COPY . .
RUN go build -o main -ldflags=-X=main.version=${VERSION} .
#RUN go build -ldflags=-X=main.version=${VERSION} .
#go build -o . -ldflags=-X=main.version=${VERSION} .

FROM debian:buster-slim
COPY --from=builder /go/src/app/main /go/bin/main
ENV PATH="/go/bin:${PATH}"
CMD ["main"]