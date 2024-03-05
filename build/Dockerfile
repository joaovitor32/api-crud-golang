FROM golang:1.21 AS build

WORKDIR /app

COPY . .

RUN apt update && apt install ca-certificates libgnutls30 -y
RUN go mod download


EXPOSE 3333

# Run
CMD ["go", "run", "cmd/server/main.go"]