#!/bin/bash

function cleanup {
    kill "$SERVICE_A_PID"
    # kill "$PRODUCTS_PID"
    # kill "$REVIEWS_PID"
}
trap cleanup EXIT

go build -o /tmp/service_a ./service_a
# go build -o /tmp/srv-products ./products
# go build -o /tmp/srv-reviews ./reviews

/tmp/srv-accounts &
SERVICE_A_PID=$!

# /tmp/srv-products &
# PRODUCTS_PID=$!

# /tmp/srv-reviews &
# REVIEWS_PID=$!

sleep 1
