#!/bin/bash
protoc email.proto  --go_out=plugins=grpc:.