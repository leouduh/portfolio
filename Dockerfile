FROM  golang:1.25-bookworm AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN apt-get update && apt-get install -y curl && \
    curl -sLo /usr/local/bin/tailwindcss \
      https://github.com/tailwindlabs/tailwindcss/releases/download/v4.3.3/tailwindcss-linux-arm64 && \
    chmod +x /usr/local/bin/tailwindcss

RUN tailwindcss -i input.css -o static/css/output.css --minify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /app/server .

# another image
FROM  gcr.io/distroless/static-debian12
WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/templates ./templates/
COPY --from=builder /app/static ./static


EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/app/server"]
