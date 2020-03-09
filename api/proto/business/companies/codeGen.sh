#!/bin/bash
protoc companies.proto  --go_out=plugins=grpc:.