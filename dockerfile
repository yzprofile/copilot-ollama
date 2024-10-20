FROM golang:1.22.2 AS builder

WORKDIR /app

COPY . .

RUN GOOS=linux go build -o copilot-ollama ./*.go 

EXPOSE 11435

ENTRYPOINT ["/app/copilot-ollama"]
