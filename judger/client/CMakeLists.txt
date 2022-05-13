cmake_minimum_required(VERSION 3.22)
project(Judge)

set(CMAKE_CXX_STANDARD 14)
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/build)

include_directories(".")

file(GLOB SOURCE "src/*.c")
add_executable(judge src/main.cpp pb/judge.pb.cc ${SOURCE})
target_link_libraries(judge protobuf pthread seccomp)



