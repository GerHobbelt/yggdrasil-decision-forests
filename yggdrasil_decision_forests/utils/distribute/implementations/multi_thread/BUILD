load("//yggdrasil_decision_forests/utils:compile.bzl", "all_proto_library", "cc_library_ydf")

package(
    default_visibility = ["//visibility:public"],
    licenses = ["notice"],
)

# Proto
# ========

all_proto_library(
    name = "multi_thread_proto",
    srcs = ["multi_thread.proto"],
    deps = ["//yggdrasil_decision_forests/utils/distribute:distribute_proto"],
)

# Library
# =======

cc_library_ydf(
    name = "multi_thread",
    srcs = ["multi_thread.cc"],
    hdrs = ["multi_thread.h"],
    deps = [
        ":multi_thread_cc_proto",
        "//yggdrasil_decision_forests/utils:concurrency",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:status_macros",
        "//yggdrasil_decision_forests/utils/distribute:core",
        "//yggdrasil_decision_forests/utils/distribute:utils",
        "@com_google_absl//absl/log",
        "@com_google_absl//absl/log:check",
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
    ],
    alwayslink = 1,
)
