load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "server",
    srcs = ["server.go"],
    importpath = "github.com/TastyPi/grail-interview/internal/participant/server",
    visibility = ["//:__subpackages__"],
    deps = [
        "//api/participant",
        "//internal/participant/storage",
        "@com_github_mennanov_fmutils//:fmutils",
        "@io_bazel_rules_go//proto/wkt:empty_go_proto",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
        "@org_golang_google_protobuf//proto:go_default_library",
    ],
)
