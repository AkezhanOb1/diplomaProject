#!/bin/bash
protoc testPB/test.proto  --go_out=plugins=grpc:.