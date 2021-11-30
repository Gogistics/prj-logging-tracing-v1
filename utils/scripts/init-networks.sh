#!/bin/bash

trap "finish" INT TERM

finish() {
    local existcode=$?
    echo "Abort..."
    # Note: add actions here to clean up environment
    echo "Cleaning up environment..."
    exit $existcode
}

LOGGING_TRACING_NETWORK="atai_logging_tracing"
LOGGING_TRACING_NETWORK_INSPECTION=$(docker network inspect $LOGGING_TRACING_NETWORK)
LOGGING_TRACING_NETWORK_INSPECTION=$?
if [ $LOGGING_TRACING_NETWORK_INSPECTION -ne 0 ]
then
    echo "Creating $LOGGING_TRACING_NETWORK network..."
    docker network create \
        --driver="bridge" \
        --subnet="190.10.0.0/24" \
        --gateway="190.10.0.1" \
        $LOGGING_TRACING_NETWORK
else
    echo "$LOGGING_TRACING_NETWORK already exists"
fi