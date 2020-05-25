FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o .

FROM amazon/aws-cli
COPY --from=builder /build/golang-cli-tool /app/
COPY --from=builder /build/cmd/bash /app/cmd/bash
WORKDIR /app
ENTRYPOINT ["./golang-cli-tool"]