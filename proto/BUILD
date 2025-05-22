load("@rules_go//proto:def.bzl", "go_proto_library")
load("@rules_proto//proto:defs.bzl", "proto_library")

proto_library(
    name = "wal_proto",
    srcs = ["wal.proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "wal_go_proto",
    importpath = "simplewal/proto",
    protos = [":wal_proto"],
)
