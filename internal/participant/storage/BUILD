load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "storage",
    srcs = [
        "in_memory.go",
        "storage.go",
    ],
    importpath = "github.com/TastyPi/grail-interview/internal/participant/storage",
    visibility = ["//:__subpackages__"],
    deps = ["//api/participant"],
)

go_test(
    name = "storage_test",
    srcs = ["in_memory_test.go"],
    embed = [":storage"],
    deps = [
        "//api/participant",
        "@com_github_google_go_cmp//cmp:go_default_library",
        "@com_github_stretchr_testify//assert",
        "@go_googleapis//google/type:date_go_proto",
        "@org_golang_google_protobuf//proto:go_default_library",
        "@org_golang_google_protobuf//testing/protocmp:go_default_library",
    ],
)
