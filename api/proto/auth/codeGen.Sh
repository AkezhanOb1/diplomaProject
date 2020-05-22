#!/bin/bash
protoc auth.proto  --go_out=plugins=grpc:.  --proto_path=$HOME/Desktop/SELF/auth/api/business