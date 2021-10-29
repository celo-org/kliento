#  docker build -t jcortejoso/check-geth-sync:latest -f examples/check_node_health/Dockerfile .
FROM golang:1.16 AS builder
WORKDIR /app
COPY . .
RUN cd examples/check_node_health/ && go build

FROM ubuntu:latest
RUN apt update && apt install -y ca-certificates && apt clean
COPY --from=builder /app/examples/check_node_health/check_node_health /
CMD ["/check_node_health"]
