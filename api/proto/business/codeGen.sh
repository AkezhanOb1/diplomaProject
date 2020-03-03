#!/bin/bash
protoc registration/registration.proto  --go_out=plugins=grpc:.