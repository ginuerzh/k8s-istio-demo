FROM golang:1.13-alpine as build

WORKDIR /src

ADD . /src/

ARG GOPROXY=https://goproxy.io
ARG GOOS=linux
ARG GOARCH=amd64

RUN GOPROXY=$GOPROXY CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH GO111MODULE=on \
    go build --ldflags="-s -w" -o /app/account-grpc github.com/ginuerzh/k8s-istio-demo/account-svc/grpc && \
    GOPROXY=$GOPROXY CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH GO111MODULE=on \
    go build --ldflags="-s -w" -o /app/account-gw github.com/ginuerzh/k8s-istio-demo/account-svc/gw && \
    echo 'nobody:x:65534:' > /src/group.nobody && \
    echo 'nobody:x:65534:65534::/:' > /src/passwd.nobody && \
    cp /src/bin/grpc_health_probe /bin/grpc_health_probe && \
    chmod +x /bin/grpc_health_probe

FROM alpine

# GRPC port
EXPOSE 8000
# Gateway port
EXPOSE 8080

COPY --from=build /src/group.nobody /etc/group
COPY --from=build /src/passwd.nobody /etc/passwd
USER nobody:nobody

COPY --from=build /app/account-grpc /bin/account-grpc
COPY --from=build /app/account-gw /bin/account-gw
# Add grpc-health-probe to use with readiness and liveness probes
COPY --from=build /bin/grpc_health_probe /bin/grpc_health_probe

ENTRYPOINT ["/bin/account-grpc"]
