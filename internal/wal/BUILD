load("@rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "wal_lib",
    srcs = ["wal.go"],
    importpath = "internal/wal",
    visibility = ["//visibility:public"],
)

go_test(
    name = "wal_test",
    srcs = ["wal_test.go"],
    embed = [":wal"],
)

go_library(
    name = "wal",
    srcs = [
        "readwrite.go",
        "wal.go",
    ],
    importpath = "github.com/shohagrana006/simplewal/internal/wal",
    visibility = ["//:__subpackages__"],
    deps = [
        "@com_github_shohagrana006_wal//pkg/api/user",
        "@com_github_shohagrana006_wal//pkg/api/wal",
        "@org_golang_google_protobuf//proto",
    ],
)
