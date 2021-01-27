load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "storage",
    srcs = [
        "in-memory.go",
        "storage.go",
    ],
    importpath = "github.com/TastyPi/grail-interview/internal/participant/storage",
    visibility = ["//:__subpackages__"],
    deps = ["//api/participant"],
)
