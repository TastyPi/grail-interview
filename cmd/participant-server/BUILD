load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "participant-server_lib",
    srcs = ["main.go"],
    importpath = "github.com/TastyPi/grail-interview/cmd/participant-server",
    visibility = ["//visibility:private"],
    deps = [
        "//api/participant",
        "//internal/participant/server",
        "@org_golang_google_grpc//:go_default_library",
        "@org_golang_google_grpc//reflection",
    ],
)

go_binary(
    name = "participant-server",
    embed = [":participant-server_lib"],
    visibility = ["//visibility:public"],
)
