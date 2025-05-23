load("//yggdrasil_decision_forests/utils:compile.bzl", "all_proto_library", "cc_binary_ydf", "cc_library_ydf")

package(
    default_visibility = ["//visibility:public"],
    licenses = ["notice"],
)

# Proto
# ========

all_proto_library(
    name = "grpc_proto",
    srcs = ["grpc.proto"],
    has_services = True,
    deps = ["//yggdrasil_decision_forests/utils/distribute:distribute_proto"],
)

# Worker binaries
# ===============

cc_binary_ydf(
    name = "grpc_worker_main",
    deps = [":grpc_worker_lib_with_main"],
)

cc_library_ydf(
    name = "grpc_worker_lib_with_main",
    srcs = ["grpc_worker_main.cc"],
    deps = [
        ":grpc_worker",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils/distribute:all_workers",
        "@com_google_absl//absl/flags:flag",
        "@com_google_absl//absl/log",
    ],
)

# Library
# =======

alias(
    name = "grpc",
    actual = ":grpc_manager",
)

cc_library_ydf(
    name = "grpc_manager",
    srcs = ["grpc_manager.cc"],
    hdrs = ["grpc_manager.h"],
    deps = [
        ":grpc_common",
        ":grpc_grpc_proto",
        "//yggdrasil_decision_forests/utils:concurrency",
        "//yggdrasil_decision_forests/utils:filesystem",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:status_macros",
        "//yggdrasil_decision_forests/utils:synchronization_primitives",
        "//yggdrasil_decision_forests/utils/distribute:core",
        "//yggdrasil_decision_forests/utils/distribute:utils",
        "@com_github_grpc_grpc//:grpc++",
        "@com_google_absl//absl/container:flat_hash_map",
        "@com_google_absl//absl/log",
        "@com_google_absl//absl/log:check",
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
        "@com_google_absl//absl/time",
    ],
    alwayslink = 1,
)

cc_library_ydf(
    name = "grpc_worker",
    srcs = ["grpc_worker.cc"],
    hdrs = ["grpc_worker.h"],
    deps = [
        ":grpc_common",
        ":grpc_grpc_proto",
        "//yggdrasil_decision_forests/utils:concurrency",
        "//yggdrasil_decision_forests/utils:filesystem",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:status_macros",
        "//yggdrasil_decision_forests/utils:synchronization_primitives",
        "//yggdrasil_decision_forests/utils/distribute:core",
        "//yggdrasil_decision_forests/utils/distribute:utils",
        "@com_github_grpc_grpc//:grpc++",
        "@com_google_absl//absl/log",
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
        "@com_google_absl//absl/time",
    ],
    alwayslink = 1,
)

cc_library_ydf(
    name = "grpc_common",
    srcs = ["grpc_common.cc"],
    hdrs = ["grpc_common.h"],
    deps = [
        "@com_github_grpc_grpc//:grpc++",
        "@com_google_absl//absl/strings",
    ],
)

# Tests
# =====

cc_test(
    name = "grpc_test",
    srcs = ["grpc_test.cc"],
    deps = [
        ":grpc_cc_proto",
        ":grpc_manager",
        ":grpc_worker",
        "//yggdrasil_decision_forests/utils:concurrency",
        "//yggdrasil_decision_forests/utils:filesystem",
        "//yggdrasil_decision_forests/utils:test",
        "//yggdrasil_decision_forests/utils/distribute",
        "//yggdrasil_decision_forests/utils/distribute:distribute_without_implementations",
        "//yggdrasil_decision_forests/utils/distribute:test_utils",
        "//yggdrasil_decision_forests/utils/distribute:toy_worker",
        "@com_google_absl//absl/log:check",
        "@com_google_absl//absl/strings",
        "@com_google_absl//absl/time",
        "@com_google_googletest//:gtest_main",
    ],
)
