load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.5/rules_go-0.18.5.tar.gz"],
    sha256 = "a82a352bffae6bee4e95f68a8d80a70e87f42c4741e6a448bec11998fcc82329",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()