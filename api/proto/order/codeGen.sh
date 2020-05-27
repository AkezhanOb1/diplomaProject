#!/bin/bash
protoc order.proto  --go_out=plugins=grpc:.  --proto_path=$HOME/Desktop/SELF/orders/api/order/