#!/bin/bash

function cleanup {
    kill "$SERVICE_A_PID"
    # kill "$PRODUCTS_PID"
    # kill "$REVIEWS_PID"
}
trap cleanup EXIT

cd ./federated_service_a && go build -o /tmp/federated_service_a ./federated_service_a && cd ..
# go build -o /tmp/srv-products ./products
# go build -o /tmp/srv-reviews ./reviews

/tmp/federated_service_a &
SERVICE_A_PID=$!

# /tmp/srv-products &
# PRODUCTS_PID=$!

# /tmp/srv-reviews &
# REVIEWS_PID=$!

sleep 1

yarn start
