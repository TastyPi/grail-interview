load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "participant_proto",
    srcs = [
        "participant.proto",
        "service.proto",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "@com_google_protobuf//:empty_proto",
        "@com_google_protobuf//:field_mask_proto",
        "@go_googleapis//google/api:annotations_proto",
        "@go_googleapis//google/type:date_proto",
    ],
)

go_proto_library(
    name = "participant_go_proto",
    compilers = ["@io_bazel_rules_go//proto:go_grpc"],
    importpath = "github.com/TastyPi/grail-interview/api/participant",
    proto = ":participant_proto",
    visibility = ["//visibility:public"],
    deps = [
        "@go_googleapis//google/api:annotations_go_proto",
        "@go_googleapis//google/type:date_go_proto",
    ],
)

go_library(
    name = "participant",
    embed = [":participant_go_proto"],
    importpath = "github.com/TastyPi/grail-interview/api/participant",
    visibility = ["//visibility:public"],
)
