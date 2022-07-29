load("//tools:defs.bzl", "go_library", "go_test", "proto_library")

package(licenses = ["notice"])

go_library(
    name = "metric",
    srcs = [
        "condmetric_kvm.go",
        "condmetric_kvm_fake.go",
        "fake_metric.go",
        "metric.go",
        "metric_unsafe.go",
        "test_utils.go",
    ],
    visibility = ["//:sandbox"],
    deps = [
        ":metric_go_proto",
        "//pkg/atomicbitops",
        "//pkg/eventchannel",
        "//pkg/gohacks",
        "//pkg/log",
        "//pkg/sync",
        "@org_golang_google_protobuf//proto:go_default_library",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)

proto_library(
    name = "metric",
    srcs = ["metric.proto"],
    visibility = ["//:sandbox"],
    deps = [
        "@com_google_protobuf//:timestamp_proto",
    ],
)

go_test(
    name = "metric_test",
    srcs = ["metric_test.go"],
    library = ":metric",
    deps = [
        ":metric_go_proto",
        "//pkg/sync",
        "@org_golang_google_protobuf//proto:go_default_library",
    ],
)

go_test(
    name = "condmetric_test",
    srcs = [
        "condmetric_kvm_disabled_test.go",
        "condmetric_kvm_enabled_test.go",
    ],
    library = ":metric",
    deps = [":metric_go_proto"],
)
