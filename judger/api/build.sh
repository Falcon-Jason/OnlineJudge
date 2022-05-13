#!/usr/bin/bash
protoc judge.proto -I . --cpp_out=../client/pb --go_out=../server