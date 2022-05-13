package service

import "path"

type compileConfig struct {
	srcName     string
	exeName     string
	command     []string
	maxCpuTime  int32
	maxRealTime int32
	maxMemory   int64
}

type runConfig struct {
	command     []string
	maxCpuTime  int32
	maxRealTime int32
	maxMemory   int64
}

type languageConfig struct {
	cc compileConfig
	rc runConfig
}

var languageConfigs = map[string]func(ws string) languageConfig{
	"c": func(ws string) languageConfig {
		return languageConfig{
			cc: compileConfig{
				srcName: path.Join(ws, "main.c"),
				exeName: path.Join(ws, "main"),
				command: []string{
					"/usr/bin/gcc",
					path.Join(ws, "main.c"),
					"-o",
					path.Join(ws, "main"),
					"-ansi",
				},
				maxCpuTime:  3000,
				maxRealTime: 5000,
				maxMemory:   128 * 1024 * 1024,
			}, rc: runConfig{
				command:     []string{path.Join(ws, "main")},
				maxCpuTime:  3000,
				maxRealTime: 5000,
				maxMemory:   128 * 1024 * 1024,
			},
		}
	},
	"cpp": func(ws string) languageConfig {
		return languageConfig{
			cc: compileConfig{
				srcName: path.Join(ws, "main.cpp"),
				exeName: path.Join(ws, "main"),
				command: []string{
					"/usr/bin/g++",
					path.Join(ws, "main.cpp"),
					"-o",
					path.Join(ws, "main"),
					"-std=c++03",
				},
				maxCpuTime:  3000,
				maxRealTime: 5000,
				maxMemory:   1024 * 1024 * 1024,
			}, rc: runConfig{
				command:     []string{path.Join(ws, "main")},
				maxCpuTime:  3000,
				maxRealTime: 5000,
				maxMemory:   128 * 1024 * 1024,
			},
		}
	},
	"cpp14": func(ws string) languageConfig {
		return languageConfig{
			cc: compileConfig{
				srcName: path.Join(ws, "main.cpp"),
				exeName: path.Join(ws, "main"),
				command: []string{
					"/usr/bin/g++",
					path.Join(ws, "main.cpp"),
					"-o",
					path.Join(ws, "main"),
					"-std=c++14",
				},
				maxCpuTime:  3000,
				maxRealTime: 5000,
				maxMemory:   1024 * 1024 * 1024,
			}, rc: runConfig{
				command:     []string{path.Join(ws, "main")},
				maxCpuTime:  3000,
				maxRealTime: 5000,
				maxMemory:   128 * 1024 * 1024,
			},
		}
	},
}
