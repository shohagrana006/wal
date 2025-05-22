load("@rules_go//go:def.bzl", "go_binary", "go_test")

go_binary(
    name = "client",
    srcs = ["main.go"],
    importpath = "simplewal/examples/basic",
)

go_test(
    name = "main_test",
    srcs = ["main_test.go"],
    deps = ["//pkg/api/proto:proto"],
)