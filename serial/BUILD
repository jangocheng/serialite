load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "serial.go",
        "serial_linux.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["@org_golang_x_sys//unix:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = ["serial_test.go"],
    library = ":go_default_library",
)
