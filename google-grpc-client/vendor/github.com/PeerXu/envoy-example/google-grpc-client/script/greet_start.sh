#! /usr/bin/env sh

greet serve --bind 127.0.0.1:5556 &
/usr/local/bin/envoy -c /etc/greet-grpc-envoy.yaml --service-cluster greet
