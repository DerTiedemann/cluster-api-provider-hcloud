load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8e968b5fcea1d2d64071872b12737bbb5514524ee5f0a4f54f5920266c261acb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
    ],
)

## Load gazelle and dependencies
http_archive(
    name = "bazel_gazelle",
    sha256 = "62ca106be173579c0a167deb23358fdfe71ffa1e4cfdddf5582af26520f1c66f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_rules_dependencies()

go_register_toolchains(version = "1.16.5")

gazelle_dependencies()

## Load kubernetes repo-infra for tools like kazel
http_archive(
    name = "io_k8s_repo_infra",
    sha256 = "5baf1b698bc20080ba714401c34ddc6ffa447757896c26399c00ab8bb39cc980",
    strip_prefix = "repo-infra-0.1.8",
    urls = [
        "https://github.com/kubernetes/repo-infra/archive/v0.1.8.tar.gz",
    ],
)

# Load repositories from external files
# gazelle:repository_macro hack/build/repos.bzl%go_repositories
load("//hack/build:repos.bzl", "go_repositories")

go_repositories()

## Load rules_docker and depdencies, for working with docker images
# Download the rules_docker repository at release v0.14.4
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "4521794f0fba2e20f3bf15846ab5e01d5332e587e9ce81629c7f96c793bb7036",
    strip_prefix = "rules_docker-0.14.4",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.14.4/rules_docker-v0.14.4.tar.gz"],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

load("@io_bazel_rules_docker//repositories:pip_repositories.bzl", "pip_deps")

pip_deps()

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
)

## Use 'static' distroless image for all builds
container_pull(
    name = "static_base",
    #tag = "3.11",
    digest = "sha256:ddba4d27a7ffc3f86dd6c2f92041af252a1f23a8e742c90e6e1297bfa1bc0c45",
    registry = "index.docker.io",
    repository = "library/alpine",
)

## Setup jsonnet
http_archive(
    name = "io_bazel_rules_jsonnet",
    sha256 = "7f51f859035cd98bcf4f70dedaeaca47fe9fbae6b199882c516d67df416505da",
    strip_prefix = "rules_jsonnet-0.3.0",
    urls = ["https://github.com/bazelbuild/rules_jsonnet/archive/0.3.0.tar.gz"],
)

load("@io_bazel_rules_jsonnet//jsonnet:jsonnet.bzl", "jsonnet_repositories")

jsonnet_repositories()

load("@jsonnet_go//bazel:repositories.bzl", "jsonnet_go_repositories")

jsonnet_go_repositories()

load("@jsonnet_go//bazel:deps.bzl", "jsonnet_go_dependencies")

jsonnet_go_dependencies()

git_repository(
    name = "com_github_jemdiggity_rules_os_dependent_http_archive",
    commit = "b1e3ed2fd829dfd1602bc31df4804ff34149f659",
    remote = "https://github.com/jemdiggity/rules_os_dependent_http_archive.git",
)

load("@com_github_jemdiggity_rules_os_dependent_http_archive//:os_dependent_http_archive.bzl", "os_dependent_http_archive")

# Packer binary dependencies
PACKER_VERSION = "1.7.3"

http_archive(
    name = "packer_linux_amd64_bin",
    build_file_content = '''filegroup(
    name="bin",
    srcs=["packer"],
    visibility = ["//visibility:public"],
)''',
    sha256 = "1a8719f0797e9e45abd98d2eb38099b09e5566ec212453052d2f21facc990c73",
    urls = ["https://releases.hashicorp.com/packer/%s/packer_%s_linux_amd64.zip" % (PACKER_VERSION, PACKER_VERSION)],
)

http_archive(
    name = "packer_darwin_amd64_bin",
    build_file_content = '''filegroup(
    name="bin",
    srcs=["packer"],
    visibility = ["//visibility:public"],
)''',
    sha256 = "aa64cc001134ffc7e1f40186dd01822242a146304f646c2dd31e441193e6a688",
    urls = ["https://releases.hashicorp.com/packer/%s/packer_%s_darwin_amd64.zip" % (PACKER_VERSION, PACKER_VERSION)],
)

KUBEBUILDER_VERSION = "2.2.0"

# kubebuilder for testing our controllers
http_archive(
    name = "kubebuilder_linux_amd64_bin",
    build_file_content = '''filegroup(
    name="bin",
    srcs=["etcd","kubectl","kube-apiserver", "kubebuilder"],
    visibility = ["//visibility:public"],
)''',
    sha256 = "9ef35a4a4e92408f7606f1dd1e68fe986fa222a88d34e40ecc07b6ffffcc8c12",
    strip_prefix = "kubebuilder_2.2.0_linux_amd64/bin/",
    urls = ["https://github.com/kubernetes-sigs/kubebuilder/releases/download/v2.2.0/kubebuilder_2.2.0_linux_amd64.tar.gz"],
)

http_archive(
    name = "kubebuilder_darwin_amd64_bin",
    build_file_content = '''filegroup(
    name="bin",
    srcs=["etcd","kubectl","kube-apiserver", "kubebuilder"],
    visibility = ["//visibility:public"],
)''',
    sha256 = "5ccb9803d391e819b606b0c702610093619ad08e429ae34401b3e4d448dd2553",
    strip_prefix = "kubebuilder_2.2.0_darwin_amd64/bin/",
    urls = ["https://github.com/kubernetes-sigs/kubebuilder/releases/download/v2.2.0/kubebuilder_2.2.0_darwin_amd64.tar.gz"],
)

# kubectl binary dependencies
KUBECTL_VERSION = "1.21.2"

http_file(
    name = "kubectl_linux_amd64_bin",
    downloaded_file_path = "kubectl",
    executable = True,
    sha256 = "55b982527d76934c2f119e70bf0d69831d3af4985f72bb87cd4924b1c7d528da",
    urls = ["https://storage.googleapis.com/kubernetes-release/release/v%s/bin/linux/amd64/kubectl" % KUBECTL_VERSION],
)

http_file(
    name = "kubectl_darwin_amd64_bin",
    downloaded_file_path = "kubectl",
    executable = True,
    sha256 = "4a6c072223d5944b98601fc9f4cfdc5652ff0919ca91210b7eed5c83f2422fa1",
    urls = ["https://storage.googleapis.com/kubernetes-release/release/v%s/bin/darwin/amd64/kubectl" % KUBECTL_VERSION],
)

go_repository(
    name = "com_github_gosuri_uitable",
    build_file_generation = "on",
    build_file_proto_mode = "disable_global",
    importpath = "github.com/gosuri/uitable",
    sum = "h1:IG2xLKRvErL3uhY6e1BylFzG+aJiwQviDDTfOKeKTpY=",
    version = "v0.0.4",
)
