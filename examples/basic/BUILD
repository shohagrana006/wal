load("@rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_binary(
    name = "client",
    embed = [":basic_lib"],
    importpath = "simplewal/examples/basic",
    visibility = ["//visibility:public"],
)

go_test(
    name = "main_test",
    srcs = ["main_test.go"],
    deps = ["//pkg/api/proto"],
)

go_library(
    name = "basic_lib",
    srcs = ["main.go"],
    importpath = "github.com/shohagrana006/simplewal/examples/basic",
    visibility = ["//visibility:private"],
    deps = ["@com_github_shohagrana006_wal//pkg/api/wal"],
)

go_test(
    name = "basic_test",
    srcs = ["main_test.go"],
    embed = [":basic_lib"],
    deps = ["@com_github_shohagrana006_wal//pkg/api/wal"],
)
