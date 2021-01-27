load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

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

go_test(
    name = "server_test",
    srcs = ["server_test.go"],
    embed = [":server"],
    deps = [
        "//api/participant",
        "//internal/participant/storage",
        "@com_github_google_go_cmp//cmp:go_default_library",
        "@com_github_stretchr_testify//assert",
        "@go_googleapis//google/type:date_go_proto",
        "@io_bazel_rules_go//proto/wkt:field_mask_go_proto",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
        "@org_golang_google_protobuf//testing/protocmp:go_default_library",
    ],
)
