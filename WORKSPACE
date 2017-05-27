git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.4.4",
)

load("@io_bazel_rules_go//go:def.bzl",
	"go_repositories", "new_go_repository")

go_repositories()

new_go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    commit = "a2e06a1",
)
