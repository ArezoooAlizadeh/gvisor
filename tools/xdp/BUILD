load("//tools:defs.bzl", "go_binary")

package(licenses = ["notice"])

go_binary(
    name = "xdp_loader",
    srcs = [
        "main.go",
    ],
    embedsrcs = [
        "//tools/xdp/bpf:drop_ebpf.o",  # keep
        "//tools/xdp/bpf:pass_ebpf.o",  # keep
    ],
    visibility = ["//:sandbox"],
    deps = [
        "@com_github_cilium_ebpf//:go_default_library",
        "@com_github_cilium_ebpf//link:go_default_library",
        "@org_golang_x_sys//unix:go_default_library",
    ],
)
