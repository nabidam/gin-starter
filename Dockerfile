FROM golang:1.21 AS build

WORKDIR /app
COPY go.mod go.sum ./


ENV http_proxy http://172.17.0.1:2080
ENV https_proxy http://172.17.0.1:2080
# ENV GOPROXY=https://host.docker.internal:2080,direct
# ENV GOPROXY=https://172.17.0.1:2080
RUN go mod download
COPY . .

WORKDIR /app/cmd/server
RUN go build -o ./main .

EXPOSE 8000

CMD ["./main"]