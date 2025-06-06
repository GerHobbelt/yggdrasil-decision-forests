load("//yggdrasil_decision_forests/utils:compile.bzl", "all_proto_library", "cc_library_ydf")

package(
    default_visibility = ["//visibility:public"],
    licenses = ["notice"],
)

cc_library_ydf(
    name = "all_dataset_formats",
    deps = [
        ":avro_example",
        ":csv_example_reader",
        ":csv_example_writer",
        "//yggdrasil_decision_forests/dataset/tensorflow_no_dep:tf_record_tf_example",
        "//yggdrasil_decision_forests/learner/distributed_decision_tree/dataset_cache:dataset_cache_reader",
    ] + select({
        "//yggdrasil_decision_forests:use_tensorflow_io": ["//yggdrasil_decision_forests/dataset/tensorflow:tf_example_io_tfrecord"],
        "//yggdrasil_decision_forests:use_ydf_tensorflow_proto": [],
        "//conditions:default": [
        ],
    }) + select({
        "//conditions:default": [
        ],
    }),
    alwayslink = 1,
)

# Libraries
# ========

cc_library_ydf(
    name = "formats",
    srcs = [
        "formats.cc",
    ],
    hdrs = [
        "formats.h",
    ],
    deps = [
        ":formats_cc_proto",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:status_macros",
        "@com_google_absl//absl/log",
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
    ],
)

cc_library_ydf(
    name = "data_spec",
    srcs = [
        "data_spec.cc",
    ],
    hdrs = [
        "data_spec.h",
    ],
    deps = [
        ":data_spec_cc_proto",
        ":example_cc_proto",
        "//yggdrasil_decision_forests/utils:hash",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:regex",
        "//yggdrasil_decision_forests/utils:status_macros",
        "@com_google_absl//absl/base:core_headers",
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
        "@com_google_absl//absl/strings:str_format",
        "@com_google_absl//absl/types:span",
    ],
)

cc_library_ydf(
    name = "example_reader",
    srcs = [
        "example_reader.cc",
    ],
    hdrs = [
        "example_reader.h",
    ],
    deps = [
        ":data_spec_cc_proto",
        ":example_reader_interface",
        ":formats",
        ":formats_cc_proto",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:status_macros",
        "@com_google_absl//absl/log",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
    ],
)

cc_library_ydf(
    name = "example_reader_interface",
    hdrs = ["example_reader_interface.h"],
    deps = [
        ":data_spec_cc_proto",
        ":example_cc_proto",
        "//yggdrasil_decision_forests/utils:registration",
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
    ],
)

cc_library_ydf(
    name = "csv_example_reader",
    srcs = [
        "csv_example_reader.cc",
    ],
    hdrs = [
        "csv_example_reader.h",
    ],
    deps = [
        ":data_spec",
        ":data_spec_cc_proto",
        ":data_spec_inference",
        ":example_cc_proto",
        ":example_reader_interface",
        "//yggdrasil_decision_forests/utils:csv",
        "//yggdrasil_decision_forests/utils:filesystem",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:sharded_io",
        "//yggdrasil_decision_forests/utils:status_macros",
        "@com_google_absl//absl/log",
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
    ],
    alwayslink = 1,
)

cc_library_ydf(
    name = "vertical_dataset",
    srcs = [
        "vertical_dataset.cc",
        "vertical_dataset_column.cc",
    ],
    hdrs = [
        "vertical_dataset.h",
    ],
    deps = [
        ":data_spec",
        ":data_spec_cc_proto",
        ":example_cc_proto",
        ":types",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:status_macros",
        "@com_google_absl//absl/base:core_headers",
        "@com_google_absl//absl/log",
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
        "@com_google_absl//absl/strings:str_format",
        "@com_google_absl//absl/types:span",
    ],
)

cc_library_ydf(
    name = "vertical_dataset_html",
    srcs = [
        "vertical_dataset_html.cc",
    ],
    hdrs = [
        "vertical_dataset_html.h",
    ],
    deps = [
        ":data_spec_cc_proto",
        ":vertical_dataset",
        "//yggdrasil_decision_forests/utils:html",
        "@com_google_absl//absl/log",
        "@com_google_absl//absl/strings",
    ],
)

cc_library_ydf(
    name = "vertical_dataset_io",
    srcs = ["vertical_dataset_io.cc"],
    hdrs = ["vertical_dataset_io.h"],
    deps = [
        ":data_spec_cc_proto",
        ":example_cc_proto",
        ":example_reader",
        ":example_reader_interface",
        ":example_writer",
        ":example_writer_interface",
        ":formats",
        ":vertical_dataset",
        "//yggdrasil_decision_forests/utils:concurrency",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:sharded_io",
        "//yggdrasil_decision_forests/utils:status_macros",
        "//yggdrasil_decision_forests/utils:usage",
        "@com_google_absl//absl/log",
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
    ],
)

cc_library_ydf(
    name = "example_writer",
    srcs = [
        "example_writer.cc",
    ],
    hdrs = [
        "example_writer.h",
    ],
    deps = [
        ":data_spec_cc_proto",
        ":example_writer_interface",
        ":formats",
        ":formats_cc_proto",
        "//yggdrasil_decision_forests/utils:status_macros",
        "//yggdrasil_decision_forests/utils:usage",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
    ],
)

cc_library_ydf(
    name = "example_writer_interface",
    hdrs = ["example_writer_interface.h"],
    deps = [
        ":data_spec_cc_proto",
        ":example_cc_proto",
        "//yggdrasil_decision_forests/utils:registration",
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/strings",
    ],
)

cc_library_ydf(
    name = "csv_example_writer",
    srcs = [
        "csv_example_writer.cc",
    ],
    hdrs = [
        "csv_example_writer.h",
    ],
    deps = [
        ":data_spec",
        ":data_spec_cc_proto",
        ":example_cc_proto",
        ":example_writer_interface",
        "//yggdrasil_decision_forests/utils:csv",
        "//yggdrasil_decision_forests/utils:filesystem",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:sharded_io",
        "//yggdrasil_decision_forests/utils:status_macros",
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/strings",
    ],
    alwayslink = 1,
)

cc_library_ydf(
    name = "data_spec_inference",
    srcs = [
        "data_spec_inference.cc",
    ],
    hdrs = [
        "data_spec_inference.h",
    ],
    deps = [
        ":data_spec",
        ":data_spec_cc_proto",
        ":formats",
        ":formats_cc_proto",
        "//yggdrasil_decision_forests/utils:accurate_sum",
        "//yggdrasil_decision_forests/utils:concurrency",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:registration",
        "//yggdrasil_decision_forests/utils:sharded_io",
        "//yggdrasil_decision_forests/utils:status_macros",
        "@com_google_absl//absl/container:flat_hash_set",
        "@com_google_absl//absl/log",
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
    ],
)

cc_library_ydf(
    name = "weight",
    srcs = [
        "weight.cc",
    ],
    hdrs = [
        "weight.h",
    ],
    deps = [
        ":data_spec",
        ":data_spec_cc_proto",
        ":example_cc_proto",
        ":vertical_dataset",
        ":weight_cc_proto",
        "//yggdrasil_decision_forests/learner:abstract_learner_cc_proto",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:status_macros",
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
    ],
)

cc_library_ydf(
    name = "synthetic_dataset",
    srcs = [
        "synthetic_dataset.cc",
    ],
    hdrs = [
        "synthetic_dataset.h",
    ],
    deps = [
        ":data_spec",
        ":formats",
        ":formats_cc_proto",
        ":synthetic_dataset_cc_proto",
        ":vertical_dataset",
        "//yggdrasil_decision_forests/dataset/tensorflow:tf_example_io_interface",
        "//yggdrasil_decision_forests/dataset/tensorflow_no_dep:tf_example",
        "//yggdrasil_decision_forests/utils:csv",
        "//yggdrasil_decision_forests/utils:filesystem",
        "//yggdrasil_decision_forests/utils:hash",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:random",
        "@com_google_absl//absl/log",
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
    ],
)

cc_library_ydf(
    name = "types",
    srcs = ["types.cc"],
    hdrs = ["types.h"],
    defines = select({
        "//yggdrasil_decision_forests:example_idx_32bits": ["YGGDRASIL_EXAMPLE_IDX_32_BITS"],
        "//yggdrasil_decision_forests:example_idx_64bits": ["YGGDRASIL_EXAMPLE_IDX_64_BITS"],
        "//conditions:default": ["YGGDRASIL_EXAMPLE_IDX_32_BITS"],
    }),
    deps = [
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/strings",
    ],
)

cc_library_ydf(
    name = "example_builder",
    srcs = ["example_builder.cc"],
    hdrs = ["example_builder.h"],
    deps = [
        ":data_spec",
        ":data_spec_cc_proto",
        ":example_cc_proto",
        "//yggdrasil_decision_forests/utils:logging",
        "@com_google_absl//absl/strings",
        "@com_google_absl//absl/types:span",
    ],
)

cc_library_ydf(
    name = "avro",
    srcs = ["avro.cc"],
    hdrs = ["avro.h"],
    deps = [
        "//yggdrasil_decision_forests/utils:bytestream",
        "//yggdrasil_decision_forests/utils:filesystem",
        "//yggdrasil_decision_forests/utils:status_macros",
        "//yggdrasil_decision_forests/utils:zlib",
        "@com_google_absl//absl/log:check",
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
        "@nlohmann_json//:json",
    ],
)

cc_library_ydf(
    name = "avro_example",
    srcs = ["avro_example.cc"],
    hdrs = ["avro_example.h"],
    deps = [
        ":avro",
        ":data_spec",
        ":data_spec_cc_proto",
        ":data_spec_inference",
        ":example_reader_interface",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:sharded_io",
        "//yggdrasil_decision_forests/utils:status_macros",
        "@com_google_absl//absl/log",
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
    ],
    alwayslink = 1,
)

# Proto
# ========

all_proto_library(
    name = "formats_proto",
    srcs = ["formats.proto"],
)

all_proto_library(
    name = "data_spec_proto",
    srcs = ["data_spec.proto"],
)

all_proto_library(
    name = "example_proto",
    srcs = ["example.proto"],
)

all_proto_library(
    name = "weight_proto",
    srcs = ["weight.proto"],
)

all_proto_library(
    name = "synthetic_dataset_proto",
    srcs = ["synthetic_dataset.proto"],
)

# Tests
# =====

cc_test(
    name = "formats_test",
    srcs = ["formats_test.cc"],
    deps = [
        ":formats",
        ":formats_cc_proto",
        "@com_google_googletest//:gtest_main",
    ],
)

cc_test(
    name = "data_spec_test",
    size = "small",
    srcs = ["data_spec_test.cc"],
    data = [
        "//yggdrasil_decision_forests/test_data",
    ],
    deps = [
        ":all_dataset_formats",
        ":data_spec",
        ":data_spec_cc_proto",
        ":data_spec_inference",
        ":example_cc_proto",
        #"@com_google_googletest//:gtest_main", # When fixed
        "@com_google_googletest//:gtest_main",
        "@com_google_absl//absl/log",
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/strings",
        "//yggdrasil_decision_forests/utils:filesystem",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:test",
    ],
)

cc_test(
    name = "example_reader_test",
    srcs = ["example_reader_test.cc"],
    data = [
        "//yggdrasil_decision_forests/test_data",
    ],
    deps = [
        ":all_dataset_formats",
        ":data_spec",
        ":data_spec_cc_proto",
        ":data_spec_inference",
        ":example_cc_proto",
        ":example_reader",
        ":example_reader_interface",
        #"@com_google_googletest//:gtest_main", # When fixed
        "@com_google_googletest//:gtest_main",
        "@com_google_absl//absl/log",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
        "//yggdrasil_decision_forests/utils:filesystem",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:test",
    ],
)

cc_test(
    name = "vertical_dataset_test",
    srcs = ["vertical_dataset_test.cc"],
    data = ["//yggdrasil_decision_forests/test_data"],
    deps = [
        ":all_dataset_formats",
        ":data_spec",
        ":data_spec_cc_proto",
        ":data_spec_inference",
        ":example_cc_proto",
        ":types",
        ":vertical_dataset",
        ":vertical_dataset_io",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
        #"@com_google_googletest//:gtest_main", # When fixed
        "@com_google_googletest//:gtest_main",
        "//yggdrasil_decision_forests/utils:filesystem",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:test",
        "//yggdrasil_decision_forests/utils:testing_macros",
    ],
)

cc_test(
    name = "vertical_dataset_io_test",
    srcs = ["vertical_dataset_io_test.cc"],
    data = ["//yggdrasil_decision_forests/test_data"],
    deps = [
        ":all_dataset_formats",
        ":data_spec_cc_proto",
        ":data_spec_inference",
        ":example_cc_proto",
        ":vertical_dataset",
        ":vertical_dataset_io",
        #"@com_google_googletest//:gtest_main", # When fixed
        "@com_google_googletest//:gtest_main",
        "@com_google_absl//absl/log",
        "@com_google_absl//absl/strings",
        "//yggdrasil_decision_forests/utils:filesystem",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:test",
    ],
)

cc_test(
    name = "example_writer_test",
    srcs = ["example_writer_test.cc"],
    deps = [
        ":all_dataset_formats",
        ":data_spec_cc_proto",
        ":example_cc_proto",
        ":example_writer",
        ":example_writer_interface",
        "//yggdrasil_decision_forests/utils:filesystem",
        "//yggdrasil_decision_forests/utils:test",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
        "@com_google_googletest//:gtest_main",
    ],
)

cc_test(
    name = "data_spec_inference_test",
    srcs = ["data_spec_inference_test.cc"],
    data = [
        "//yggdrasil_decision_forests/test_data",
    ],
    deps = [
        ":all_dataset_formats",
        ":data_spec",
        ":data_spec_cc_proto",
        ":data_spec_inference",
        ":formats",
        ":formats_cc_proto",
        ":synthetic_dataset",
        ":synthetic_dataset_cc_proto",
        #"@com_google_googletest//:gtest_main", # When fixed
        "@com_google_googletest//:gtest_main",
        "@com_google_absl//absl/log",
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/status:statusor",
        "@com_google_absl//absl/strings",
        "//yggdrasil_decision_forests/utils:filesystem",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:sharded_io",
        "//yggdrasil_decision_forests/utils:test",
    ],
)

cc_test(
    name = "vertical_dataset_html_test",
    srcs = ["vertical_dataset_html_test.cc"],
    data = ["//yggdrasil_decision_forests/test_data"],
    deps = [
        ":data_spec_cc_proto",
        ":data_spec_inference",
        ":vertical_dataset",
        ":vertical_dataset_io",
        #"@com_google_googletest//:gtest_main", # When fixed
        "@com_google_googletest//:gtest_main",
        ":vertical_dataset_html",
        "@com_google_absl//absl/strings",
        "//yggdrasil_decision_forests/utils:filesystem",
        ":all_dataset_formats",
        "//yggdrasil_decision_forests/utils:test",
    ],
)

cc_test(
    name = "weight_test",
    srcs = ["weight_test.cc"],
    deps = [
        ":data_spec_cc_proto",
        ":example_cc_proto",
        ":vertical_dataset",
        ":weight",
        ":weight_cc_proto",
        "//yggdrasil_decision_forests/utils:logging",
        "//yggdrasil_decision_forests/utils:test",
        "@com_google_absl//absl/status",
        "@com_google_googletest//:gtest_main",
    ],
)

cc_test(
    name = "synthetic_dataset_test",
    srcs = ["synthetic_dataset_test.cc"],
    deps = [
        ":all_dataset_formats",
        ":data_spec",
        ":data_spec_cc_proto",
        ":data_spec_inference",
        ":synthetic_dataset",
        ":synthetic_dataset_cc_proto",
        "//yggdrasil_decision_forests/utils:filesystem",
        "//yggdrasil_decision_forests/utils:test",
        "//yggdrasil_decision_forests/utils:testing_macros",
        "@com_google_absl//absl/log",
        "@com_google_absl//absl/log:check",
        "@com_google_absl//absl/status",
        "@com_google_absl//absl/strings",
        "@com_google_googletest//:gtest_main",
    ],
)

cc_test(
    name = "example_builder_test",
    srcs = ["example_builder_test.cc"],
    deps = [
        ":data_spec",
        ":example_builder",
        "//yggdrasil_decision_forests/utils:test",
        "@com_google_absl//absl/types:span",
        "@com_google_googletest//:gtest_main",
    ],
)

cc_test(
    name = "avro_test",
    srcs = ["avro_test.cc"],
    data = [
        "//yggdrasil_decision_forests/test_data",
    ],
    deps = [
        ":avro",
        "//yggdrasil_decision_forests/utils:bytestream",
        "//yggdrasil_decision_forests/utils:filesystem",
        "//yggdrasil_decision_forests/utils:test",
        "//yggdrasil_decision_forests/utils:testing_macros",
        "@com_google_absl//absl/log",
        "@com_google_absl//absl/status",
        "@com_google_googletest//:gtest_main",
    ],
)

cc_test(
    name = "avro_example_test",
    srcs = ["avro_example_test.cc"],
    data = [
        "//yggdrasil_decision_forests/test_data",
    ],
    deps = [
        ":avro_example",
        ":data_spec",
        ":data_spec_cc_proto",
        ":data_spec_inference",
        ":example_reader",
        "//yggdrasil_decision_forests/utils:filesystem",
        "//yggdrasil_decision_forests/utils:test",
        "//yggdrasil_decision_forests/utils:testing_macros",
        "@com_google_absl//absl/log",
        "@com_google_absl//absl/strings",
        "@com_google_googletest//:gtest_main",
    ],
)
