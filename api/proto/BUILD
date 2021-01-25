load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@rules_proto//proto:defs.bzl", "proto_library")

proto_library(
    name = "participant_proto",
    srcs = ["participant.proto"],
    deps = [
        "@com_google_googleapis//google/api:resource_proto",
        "@com_google_googleapis//google/type:date_proto",
    ],
)

go_proto_library(
    name = "participant_go_proto",
    importpath = "github.com/TastyPi/grail-interview/api",
    proto = ":participant_proto",
    deps = [
        "@com_google_googleapis//google/api:annotations_go_proto",
        "@com_google_googleapis//google/api:resource_go_proto",
        "@com_google_googleapis//google/type:date_go_proto",
    ],
)

proto_library(
    name = "participant_service_proto",
    srcs = ["participant-service.proto"],
    deps = [
        ":participant_proto",
        "@com_google_googleapis//google/api:annotations_proto",
        "@com_google_googleapis//google/api:client_proto",
        "@com_google_googleapis//google/api:field_behavior_proto",
        "@com_google_googleapis//google/api:resource_proto",
        "@com_google_protobuf//:empty_proto",
        "@com_google_protobuf//:field_mask_proto",
    ],
)

go_proto_library(
    name = "participant_service_go_proto",
    importpath = "github.com/TastyPi/grail-interview/api/service",
    proto = ":participant_service_proto",
    deps = [
        ":participant_go_proto",
        "@com_google_googleapis//google/api:annotations_go_proto",
        "@com_google_googleapis//google/api:client_go_proto",
        "@com_google_googleapis//google/api:field_behavior_go_proto",
        "@com_google_googleapis//google/api:resource_go_proto",
    ],
)
