#!/bin/bash
protoc companyServices.proto  --go_out=plugins=grpc:.