# Use golang container to compile static binary
FROM golang as builder
WORKDIR /go/gone
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
COPY . .
RUN go build --ldflags="-w -s" -o gone

# Build minimalist container with only the binary
FROM scratch
COPY --from=builder /go/gone/gone .
ENTRYPOINT ["/gone"]
EXPOSE 80
