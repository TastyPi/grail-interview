load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "main",
    srcs = ["main.go"],
    deps = [
        "//api/proto:participant_go_proto",
        "//api/proto:participant_service_go_proto",
    ],
)