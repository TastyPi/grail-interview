load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "server",
    srcs = ["server.go"],
    importpath = "github.com/TastyPi/grail-interview/internal/participant/server",
    visibility = ["//:__subpackages__"],
    deps = [
        "//api/participant",
        "@io_bazel_rules_go//proto/wkt:empty_go_proto",
    ],
)
