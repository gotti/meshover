FROM golang:1.19 AS builder

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY / ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/server/main.go

FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=builder /workspace/app .

USER 65532:65532
ENTRYPOINT ["/app"]
