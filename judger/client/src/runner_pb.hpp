#pragma once

#include "pb/judge.pb.h"
#include "runner.h"

void GetConfig(const judge::Config &pb, struct Config &c) {
    c.max_cpu_time = pb.max_cpu_time() > 0 ? pb.max_cpu_time() : UNLIMITED;
    c.max_real_time = pb.max_real_time() > 0 ? pb.max_cpu_time() : UNLIMITED;
    c.max_memory = pb.max_memory() > 0 ? pb.max_memory() : UNLIMITED;
    c.memory_limit_check_only = pb.memory_limit_check_only() != 0 ? 1 : 0;
    c.max_stack = pb.max_stack() > 0 ? pb.max_stack() : 16 * 1024 * 1024;
    c.max_process_number = pb.max_process_number() > 0 ? pb.max_process_number() : UNLIMITED;
    c.max_output_size = pb.max_output_size() > 0 ? pb.max_output_size() : UNLIMITED;
    c.exe_path = pb.exe_path().c_str();
    c.input_path = !pb.input_path().empty() ? pb.input_path().c_str() : "/dev/stdin";
    c.output_path = !pb.output_path().empty() ? pb.output_path().c_str() : "/dev/stdout";
    c.error_path = !pb.error_path().empty() ? pb.error_path().c_str() : "/dev/stderr";
    c.args[0] = const_cast<char *>(c.exe_path);

    [&]() {
        int i = 1;
        if (!pb.args().empty()) {
            for (i = 1; i < pb.args().size() + 1; i++) {
                c.args[i] = const_cast<char *>(pb.args().at(i - 1).c_str());
            }
        }
        c.args[i] = nullptr;
    }();

    [&]() {
        int i = 0;
        if (!pb.env().empty()) {
            for (; i < pb.env().size(); i++) {
                c.env[i] = const_cast<char *>(pb.env().at(i).c_str());
            }
        }
        c.env[i] = nullptr;
    }();

    c.log_path = !pb.log_path().empty() ? pb.log_path().c_str() : "judger.log";
    c.uid = pb.uid();
    c.gid = pb.gid();
    c.use_seccomp = pb.use_seccomp();
}

void GetResult(judge::Result &pb, const struct Result &r) {
    pb.set_cpu_time(r.cpu_time);
    pb.set_real_time(r.real_time);
    pb.set_memory(r.memory);
    pb.set_signal(r.signal);
    pb.set_exit_code(r.exit_code);
    pb.set_error(judge::ErrorCode(r.error));
    pb.set_result(judge::ResultCode(r.result));
}