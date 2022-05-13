#!/usr/bin/bash

protoc submission.proto -I . --go_out=../backend/ --go-grpc_out=../backend/

protoc submission.proto -I . --go_out=../judger/server/ --go-grpc_out=../judger/server/

protoc judge_result.proto -I . --go_out=../judger/server/ --go-grpc_out=../judger/server/

protoc judge_result.proto -I . --go_out=../write_back --go-grpc_out=../write_back