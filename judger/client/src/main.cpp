#include <iostream>
#include "pb/judge.pb.h"
#include "runner_pb.hpp"
#include "runner.h"

int main() {
    judge::Config config{};
    Config cConfig{};

    config.ParseFromIstream(&std::cin);
    GetConfig(config, cConfig);

    Result cResult{};
    run(&cConfig, &cResult);

    judge::Result result{};
    GetResult(result, cResult);

    result.SerializeToOstream(&std::cout);
    return 0;
}
