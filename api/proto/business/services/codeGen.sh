#!/bin/bash
protoc services.proto  --go_out=plugins=grpc:.