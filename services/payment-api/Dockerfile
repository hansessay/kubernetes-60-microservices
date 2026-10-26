FROM golang:1.27 AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /platform-api .

FROM registry.access.redhat.com/ubi9/ubi-minimal

COPY --from=builder /platform-api /platform-api

USER 1001

EXPOSE 8080

ENTRYPOINT ["/platform-api"]