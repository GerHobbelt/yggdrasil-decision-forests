/*
 * Copyright 2022 Google LLC.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include <limits>
#include <memory>
#include <string>
#include <unordered_map>
#include <utility>
#include <vector>

#include "gmock/gmock.h"
#include "gtest/gtest.h"
#include "absl/log/log.h"
#include "yggdrasil_decision_forests/dataset/data_spec.pb.h"
#include "yggdrasil_decision_forests/dataset/vertical_dataset.h"
#include "yggdrasil_decision_forests/model/decision_tree/decision_tree.h"
#include "yggdrasil_decision_forests/model/gradient_boosted_trees/gradient_boosted_trees.h"
#include "yggdrasil_decision_forests/serving/example_set.h"
#include "yggdrasil_decision_forests/utils/test.h"

#include "yggdrasil_decision_forests/serving/decision_forest/quick_scorer_extended.h"

namespace yggdrasil_decision_forests {
namespace serving {
namespace decision_forest {
namespace {

using model::decision_tree::DecisionTree;
using model::decision_tree::NodeWithChildren;
using model::decision_tree::proto::Condition;
using model::gradient_boosted_trees::GradientBoostedTreesModel;
using model::gradient_boosted_trees::proto::Loss;
using testing::ElementsAre;

void BuildToyModelAndToyDataset(
    const model::proto::Task task, const bool use_num_feature,
    const bool use_cat_feature, const bool use_discnum_feature,
    const bool use_catset_feature, GradientBoostedTreesModel* model,
    dataset::VerticalDataset* dataset, const int duplicate_factor = 1) {
  dataset::proto::DataSpecification dataspec;
  if (use_num_feature) {
    {
      auto* d = dataspec.add_columns();
      d->set_name("num_1");
      d->set_type(dataset::proto::ColumnType::NUMERICAL);
    }
    {
      auto* d = dataspec.add_columns();
      d->set_name("num_2");
      d->set_type(dataset::proto::ColumnType::NUMERICAL);
    }
  }
  if (use_cat_feature) {
    auto* d = dataspec.add_columns();
    d->set_name("cat");
    d->set_type(dataset::proto::ColumnType::CATEGORICAL);
    d->mutable_categorical()->set_is_already_integerized(true);
    d->mutable_categorical()->set_number_of_unique_values(4);
  }
  if (use_discnum_feature) {
    auto* d = dataspec.add_columns();
    d->set_name("disc_num");
    d->set_type(dataset::proto::ColumnType::DISCRETIZED_NUMERICAL);
    d->mutable_numerical()->set_mean(-1);
    d->mutable_discretized_numerical()->add_boundaries(0.0);
    d->mutable_discretized_numerical()->add_boundaries(0.1);
    d->mutable_discretized_numerical()->add_boundaries(0.2);
    d->mutable_discretized_numerical()->add_boundaries(0.3);
  }
  if (use_catset_feature) {
    auto* d = dataspec.add_columns();
    d->set_name("cat_set");
    d->set_type(dataset::proto::ColumnType::CATEGORICAL_SET);
    d->mutable_categorical()->set_is_already_integerized(false);
    d->mutable_categorical()->set_number_of_unique_values(4);
    auto& items = *d->mutable_categorical()->mutable_items();
    items["v0"].set_index(0);
    items["v1"].set_index(1);
    items["v2"].set_index(2);
    items["v3"].set_index(3);
  }

  dataset->set_data_spec(dataspec);
  ASSERT_OK(dataset->CreateColumnsFromDataspec());

  std::unordered_map<std::string, std::string> example;
  if (use_num_feature) {
    example["num_1"] = "0.5";
    example["num_2"] = "0.5";
  }
  if (use_cat_feature) {
    example["cat"] = "1";
  }
  if (use_discnum_feature) {
    example["disc_num"] = "1.5";
  }
  if (use_catset_feature) {
    example["cat_set"] = "2 3";
  }
  ASSERT_OK(dataset->AppendExampleWithStatus(example));

  struct NodeHelper {
    Condition* condition;
    NodeWithChildren* pos;
    NodeWithChildren* neg;
    NodeWithChildren* node;
  };

  const auto split_node = [](NodeWithChildren* node,
                             const int attribute) -> NodeHelper {
    node->CreateChildren();
    node->mutable_node()->mutable_condition()->set_attribute(attribute);
    return {/*.condition =*/
            node->mutable_node()->mutable_condition()->mutable_condition(),
            /*.pos =*/node->mutable_pos_child(),
            /*.neg =*/node->mutable_neg_child(),
            /*.node =*/node};
  };

  model->set_task(task);
  model->set_label_col_idx(0);
  model->set_data_spec(dataspec);
  model::gradient_boosted_trees::proto::LossConfiguration loss_config;
  model->set_loss(Loss::SQUARED_ERROR, loss_config);
  model->mutable_initial_predictions()->push_back(duplicate_factor);

  for (int duplication_idx = 0; duplication_idx < duplicate_factor;
       duplication_idx++) {
    if (use_num_feature) {
      auto tree = std::make_unique<DecisionTree>();
      tree->CreateRoot();
      auto n1 = split_node(tree->mutable_root(), 1);
      n1.condition->mutable_higher_condition()->set_threshold(2.0f);

      auto n2 = split_node(n1.pos, 1);
      n2.condition->mutable_higher_condition()->set_threshold(3.0f);
      n2.pos->mutable_node()->mutable_regressor()->set_top_value(1.f);
      n2.neg->mutable_node()->mutable_regressor()->set_top_value(2.f);

      auto n3 = split_node(n1.neg, 1);
      n3.condition->mutable_higher_condition()->set_threshold(1.0f);
      n3.pos->mutable_node()->mutable_regressor()->set_top_value(3.f);
      n3.neg->mutable_node()->mutable_regressor()->set_top_value(4.f);

      model->mutable_decision_trees()->push_back(std::move(tree));
    }
    if (use_cat_feature && use_num_feature) {
      auto tree = std::make_unique<DecisionTree>();
      tree->CreateRoot();
      auto n1 = split_node(tree->mutable_root(), 2);
      n1.condition->mutable_contains_condition()->add_elements(1);
      n1.condition->mutable_contains_condition()->add_elements(2);

      auto n2 = split_node(n1.pos, 1);
      n2.condition->mutable_higher_condition()->set_threshold(2.5f);
      n2.pos->mutable_node()->mutable_regressor()->set_top_value(10.f);
      n2.neg->mutable_node()->mutable_regressor()->set_top_value(20.f);

      auto n3 = split_node(n1.neg, 1);
      n3.condition->mutable_higher_condition()->set_threshold(1.5f);
      n3.pos->mutable_node()->mutable_regressor()->set_top_value(30.f);
      n3.neg->mutable_node()->mutable_regressor()->set_top_value(40.f);

      model->mutable_decision_trees()->push_back(std::move(tree));
    }

    if (use_cat_feature && use_num_feature) {
      auto tree = std::make_unique<DecisionTree>();
      tree->CreateRoot();
      auto n1 = split_node(tree->mutable_root(), 1);
      n1.condition->mutable_higher_condition()->set_threshold(10.0f);

      auto n2 = split_node(n1.pos, 2);
      n2.condition->mutable_contains_bitmap_condition()->set_elements_bitmap(
          "\x02");  // [1]
      n2.pos->mutable_node()->mutable_regressor()->set_top_value(100.f);
      n2.neg->mutable_node()->mutable_regressor()->set_top_value(200.f);

      auto n3 = split_node(n1.neg, 2);
      n3.condition->mutable_contains_bitmap_condition()->set_elements_bitmap(
          "\x0A");  // [1,3]
      n3.pos->mutable_node()->mutable_regressor()->set_top_value(300.f);
      n3.neg->mutable_node()->mutable_regressor()->set_top_value(400.f);

      model->mutable_decision_trees()->push_back(std::move(tree));
    }

    if (use_discnum_feature) {
      auto tree = std::make_unique<DecisionTree>();
      tree->CreateRoot();
      auto n1 = split_node(tree->mutable_root(), 3);
      n1.condition->mutable_discretized_higher_condition()->set_threshold(
          2);  // value>=0.1f
      n1.neg->mutable_node()->mutable_regressor()->set_top_value(10000.f);
      n1.pos->mutable_node()->mutable_regressor()->set_top_value(20000.f);
      model->mutable_decision_trees()->push_back(std::move(tree));
    }

    if (use_catset_feature) {
      auto tree = std::make_unique<DecisionTree>();
      tree->CreateRoot();
      auto n1 = split_node(tree->mutable_root(), 4);
      n1.condition->mutable_contains_condition()->add_elements(2);
      n1.condition->mutable_contains_condition()->add_elements(3);
      n1.node->mutable_node()->mutable_condition()->set_na_value(true);

      n1.neg->mutable_node()->mutable_regressor()->set_top_value(1000.f);

      auto n2 = split_node(n1.pos, 4);
      n2.condition->mutable_contains_bitmap_condition()->set_elements_bitmap(
          "\x08");  // [3]
      n2.neg->mutable_node()->mutable_regressor()->set_top_value(2000.f);
      n2.pos->mutable_node()->mutable_regressor()->set_top_value(3000.f);
      n2.node->mutable_node()->mutable_condition()->set_na_value(false);
      model->mutable_decision_trees()->push_back(std::move(tree));
    }
  }
}

TEST(QuickScorer, Compilation_Num_Cat_DiscNum) {
  GradientBoostedTreesModel model;
  dataset::VerticalDataset dataset;
  BuildToyModelAndToyDataset(model::proto::Task::REGRESSION, true, true, true,
                             /*use_catset_feature=*/false, &model, &dataset);
  GradientBoostedTreesRegressionQuickScorerExtended quick_scorer_model;
  ASSERT_OK(GenericToSpecializedModel(model, &quick_scorer_model));

  const auto model_description = DescribeQuickScorer(quick_scorer_model);
  LOG(INFO) << "Model:\n" << model_description;

  EXPECT_EQ(quick_scorer_model.features().input_features().size(), 3);

  // Examples in FORMAT_FEATURE_MAJOR, see decision_forest.h.
  using V = NumericalOrCategoricalValue;
  std::vector<V> examples = {
      // Feature 1
      V::Numerical(0.5f),
      V::Numerical(1.0f),
      V::Numerical(1.5f),
      V::Numerical(2.5f),
      V::Numerical(3.5f),
      // Feature 2
      V::Categorical(0),
      V::Categorical(1),
      V::Categorical(2),
      V::Categorical(0),
      V::Categorical(1),
      // Feature 3
      V::Numerical(0.00f),
      V::Numerical(0.05f),
      V::Numerical(0.10f),
      V::Numerical(0.20f),
      V::Numerical(0.30f),
  };
  std::vector<float> predictions;
  PredictQuickScorer(quick_scorer_model, examples, 5, &predictions);

  EXPECT_THAT(predictions,
              ElementsAre(1 + 4 + 40 + 400 + 10000, 1 + 3 + 20 + 300 + 10000,
                          1 + 3 + 20 + 400 + 20000, 1 + 2 + 30 + 400 + 20000,
                          1 + 1 + 10 + 300 + 20000));
}

TEST(QuickScorer, Compilation_Num) {
  GradientBoostedTreesModel model;
  dataset::VerticalDataset dataset;
  BuildToyModelAndToyDataset(
      model::proto::Task::REGRESSION, /*use_num_feature=*/true,
      /*use_cat_feature=*/false, /*use_discnum_feature=*/false,
      /*use_catset_feature=*/false, &model, &dataset);
  GradientBoostedTreesRegressionQuickScorerExtended quick_scorer_model;
  ASSERT_OK(GenericToSpecializedModel(model, &quick_scorer_model));

  const auto model_description = DescribeQuickScorer(quick_scorer_model);
  LOG(INFO) << "Model:\n" << model_description;

  EXPECT_EQ(quick_scorer_model.features().input_features().size(), 1);

  // Examples in FORMAT_FEATURE_MAJOR, see decision_forest.h.
  using V = NumericalOrCategoricalValue;
  std::vector<V> examples = {
      // Feature 1
      V::Numerical(0.5f), V::Numerical(1.0f), V::Numerical(1.5f),
      V::Numerical(2.5f), V::Numerical(3.5f),
  };
  std::vector<float> predictions;
  PredictQuickScorer(quick_scorer_model, examples, 5, &predictions);

  EXPECT_THAT(predictions, ElementsAre(1 + 4, 1 + 3, 1 + 3, 1 + 2, 1 + 1));
}

TEST(QuickScorer, Compilation_Num_Nan) {
  struct NodeHelper {
    Condition* condition;
    NodeWithChildren* pos;
    NodeWithChildren* neg;
    NodeWithChildren* node;
  };

  const auto split_node = [](NodeWithChildren* node,
                             const int attribute) -> NodeHelper {
    node->CreateChildren();
    node->mutable_node()->mutable_condition()->set_attribute(attribute);
    return {/*.condition =*/
            node->mutable_node()->mutable_condition()->mutable_condition(),
            /*.pos =*/node->mutable_pos_child(),
            /*.neg =*/node->mutable_neg_child(),
            /*.node =*/node};
  };

  GradientBoostedTreesModel model;
  dataset::VerticalDataset dataset;
  dataset::proto::DataSpecification dataspec;
  auto* d_label = dataspec.add_columns();
  d_label->set_name("label");
  d_label->set_type(dataset::proto::ColumnType::NUMERICAL);
  auto* d_num = dataspec.add_columns();
  d_num->set_name("num");
  d_num->set_type(dataset::proto::ColumnType::NUMERICAL);
  dataset.set_data_spec(dataspec);
  ASSERT_OK(dataset.CreateColumnsFromDataspec());

  model.set_task(model::proto::Task::REGRESSION);
  model.set_label_col_idx(0);
  model.set_data_spec(dataspec);
  model::gradient_boosted_trees::proto::LossConfiguration loss_config;
  model.set_loss(Loss::SQUARED_ERROR, loss_config);
  model.mutable_initial_predictions()->push_back(1.f);

  auto tree = std::make_unique<DecisionTree>();
  tree->CreateRoot();
  auto n1 = split_node(tree->mutable_root(), 1);
  n1.condition->mutable_higher_condition()->set_threshold(2.0f);
  n1.node->mutable_node()->mutable_condition()->set_na_value(true);
  n1.neg->mutable_node()->mutable_regressor()->set_top_value(1.f);
  n1.pos->mutable_node()->mutable_regressor()->set_top_value(2.f);
  model.mutable_decision_trees()->push_back(std::move(tree));

  GradientBoostedTreesRegressionQuickScorerExtended quick_scorer_model;
  ASSERT_OK(GenericToSpecializedModel(model, &quick_scorer_model));

  EXPECT_FALSE(quick_scorer_model.global_imputation_optimization);

  const auto model_description = DescribeQuickScorer(quick_scorer_model);
  LOG(INFO) << "Model:\n" << model_description;

  EXPECT_EQ(quick_scorer_model.features().input_features().size(), 1);

  // Examples in FORMAT_FEATURE_MAJOR, see decision_forest.h.
  using V = NumericalOrCategoricalValue;
  std::vector<V> examples = {
      // Feature 1
      V::Numerical(0.5f),
      V::Numerical(std::numeric_limits<float>::quiet_NaN()),
      V::Numerical(3.5f),
      V::Numerical(std::numeric_limits<float>::quiet_NaN()),
      V::Numerical(std::numeric_limits<float>::quiet_NaN()),
  };
  std::vector<float> predictions;
  PredictQuickScorer(quick_scorer_model, examples, 5, &predictions);

  EXPECT_THAT(predictions, ElementsAre(1 + 1, 1 + 2, 1 + 2, 1 + 2, 1 + 2));
}

TEST(QuickScorer, Compilation_Num_Cat) {
  GradientBoostedTreesModel model;
  dataset::VerticalDataset dataset;
  BuildToyModelAndToyDataset(model::proto::Task::REGRESSION,
                             /*use_num_feature=*/true, /*use_cat_feature=*/true,
                             /*use_discnum_feature=*/false,
                             /*use_catset_feature=*/false, &model, &dataset);
  GradientBoostedTreesRegressionQuickScorerExtended quick_scorer_model;
  ASSERT_OK(GenericToSpecializedModel(model, &quick_scorer_model));

  const auto model_description = DescribeQuickScorer(quick_scorer_model);
  LOG(INFO) << "Model:\n" << model_description;

  ASSERT_EQ(quick_scorer_model.features().input_features().size(), 2);

  // Examples in FORMAT_FEATURE_MAJOR, see decision_forest.h.
  using V = NumericalOrCategoricalValue;
  std::vector<V> examples = {
      // Feature 1
      V::Numerical(0.5f),
      V::Numerical(1.0f),
      V::Numerical(1.5f),
      V::Numerical(2.5f),
      V::Numerical(3.5f),
      // Feature 2
      V::Categorical(0),
      V::Categorical(1),
      V::Categorical(2),
      V::Categorical(0),
      V::Categorical(1),
  };
  std::vector<float> predictions;
  PredictQuickScorer(quick_scorer_model, examples, 5, &predictions);

  EXPECT_THAT(predictions,
              ElementsAre(1 + 4 + 40 + 400, 1 + 3 + 20 + 300, 1 + 3 + 20 + 400,
                          1 + 2 + 30 + 400, 1 + 1 + 10 + 300));
}

TEST(QuickScorer, ExampleSet) {
  GradientBoostedTreesModel model;
  dataset::VerticalDataset dataset;
  BuildToyModelAndToyDataset(model::proto::Task::REGRESSION,
                             /*use_num_feature=*/true, /*use_cat_feature=*/true,
                             /*use_discnum_feature=*/true,
                             /*use_catset_feature=*/true, &model, &dataset);
  GradientBoostedTreesRegressionQuickScorerExtended quick_scorer_model;
  ASSERT_OK(GenericToSpecializedModel(model, &quick_scorer_model));

  const auto model_description = DescribeQuickScorer(quick_scorer_model);
  LOG(INFO) << "Model:\n" << model_description;

  GradientBoostedTreesRegressionQuickScorerExtended::ExampleSet examples(
      5, quick_scorer_model);
  examples.FillMissing(quick_scorer_model);

  EXPECT_EQ(quick_scorer_model.features().input_features().size(), 4);

  const auto feature_1 =
      GradientBoostedTreesRegressionQuickScorerExtended::ExampleSet::
          GetNumericalFeatureId("num_2", quick_scorer_model)
              .value();
  const auto feature_2 =
      GradientBoostedTreesRegressionQuickScorerExtended::ExampleSet::
          GetCategoricalFeatureId("cat", quick_scorer_model)
              .value();
  const auto feature_3 =
      GradientBoostedTreesRegressionQuickScorerExtended::ExampleSet::
          GetCategoricalSetFeatureId("cat_set", quick_scorer_model)
              .value();
  const auto feature_4 =
      GradientBoostedTreesRegressionQuickScorerExtended::ExampleSet::
          GetNumericalFeatureId("disc_num", quick_scorer_model)
              .value();

  examples.SetNumerical(0, feature_1, 0.5f, quick_scorer_model);
  examples.SetNumerical(1, feature_1, 1.0f, quick_scorer_model);
  examples.SetNumerical(2, feature_1, 1.5f, quick_scorer_model);
  examples.SetNumerical(3, feature_1, 2.5f, quick_scorer_model);
  examples.SetNumerical(4, feature_1, 3.5f, quick_scorer_model);

  examples.SetCategorical(0, feature_2, 0, quick_scorer_model);
  examples.SetCategorical(1, feature_2, 1, quick_scorer_model);
  examples.SetCategorical(2, feature_2, 2, quick_scorer_model);
  examples.SetCategorical(3, feature_2, 0, quick_scorer_model);
  examples.SetCategorical(4, feature_2, 1, quick_scorer_model);

  examples.SetCategoricalSet(0, feature_3, {"v1"}, quick_scorer_model);
  examples.SetCategoricalSet(1, feature_3, {"v2"}, quick_scorer_model);
  examples.SetCategoricalSet(2, feature_3, {"v3"}, quick_scorer_model);
  examples.SetCategoricalSet(3, feature_3, std::vector<std::string>{"v2", "v3"},
                             quick_scorer_model);
  examples.SetMissingCategoricalSet(4, feature_3, quick_scorer_model);

  examples.SetNumerical(0, feature_4, 0.00f, quick_scorer_model);
  examples.SetNumerical(1, feature_4, 0.05f, quick_scorer_model);
  examples.SetNumerical(2, feature_4, 0.10f, quick_scorer_model);
  examples.SetNumerical(3, feature_4, 0.20f, quick_scorer_model);
  examples.SetNumerical(4, feature_4, 0.30f, quick_scorer_model);

  std::vector<float> predictions;
  Predict(quick_scorer_model, examples, 5, &predictions);

  EXPECT_THAT(predictions, ElementsAre(1 + 4 + 40 + 400 + 1000 + 10000,
                                       1 + 3 + 20 + 300 + 2000 + 10000,
                                       1 + 3 + 20 + 400 + 3000 + 20000,
                                       1 + 2 + 30 + 400 + 3000 + 20000,
                                       1 + 1 + 10 + 300 + 2000 + 20000));
}

TEST(QuickScorer, ExceedStackBuffer) {
  const int duplicate_factor = 200;

  GradientBoostedTreesModel model;

  dataset::VerticalDataset dataset;
  BuildToyModelAndToyDataset(
      model::proto::Task::REGRESSION, /*use_num_feature=*/true,
      /*use_cat_feature=*/true, /*use_discnum_feature=*/true,
      /*use_catset_feature=*/true, &model, &dataset, duplicate_factor);
  GradientBoostedTreesRegressionQuickScorerExtended quick_scorer_model;
  ASSERT_OK(GenericToSpecializedModel(model, &quick_scorer_model));

  const auto model_description = DescribeQuickScorer(quick_scorer_model);
  LOG(INFO) << "Model:\n" << model_description;

  GradientBoostedTreesRegressionQuickScorerExtended::ExampleSet examples(
      5, quick_scorer_model);
  examples.FillMissing(quick_scorer_model);

  EXPECT_EQ(quick_scorer_model.features().input_features().size(), 4);

  const auto feature_1 =
      GradientBoostedTreesRegressionQuickScorerExtended::ExampleSet::
          GetNumericalFeatureId("num_2", quick_scorer_model)
              .value();
  const auto feature_2 =
      GradientBoostedTreesRegressionQuickScorerExtended::ExampleSet::
          GetCategoricalFeatureId("cat", quick_scorer_model)
              .value();
  const auto feature_3 =
      GradientBoostedTreesRegressionQuickScorerExtended::ExampleSet::
          GetCategoricalSetFeatureId("cat_set", quick_scorer_model)
              .value();

  const auto feature_4 =
      GradientBoostedTreesRegressionQuickScorerExtended::ExampleSet::
          GetNumericalFeatureId("disc_num", quick_scorer_model)
              .value();

  examples.SetNumerical(0, feature_1, 0.5f, quick_scorer_model);
  examples.SetNumerical(1, feature_1, 1.0f, quick_scorer_model);
  examples.SetNumerical(2, feature_1, 1.5f, quick_scorer_model);
  examples.SetNumerical(3, feature_1, 2.5f, quick_scorer_model);
  examples.SetNumerical(4, feature_1, 3.5f, quick_scorer_model);

  examples.SetCategorical(0, feature_2, 0, quick_scorer_model);
  examples.SetCategorical(1, feature_2, 1, quick_scorer_model);
  examples.SetCategorical(2, feature_2, 2, quick_scorer_model);
  examples.SetCategorical(3, feature_2, 0, quick_scorer_model);
  examples.SetCategorical(4, feature_2, 1, quick_scorer_model);

  examples.SetCategoricalSet(0, feature_3, {"v1"}, quick_scorer_model);
  examples.SetCategoricalSet(1, feature_3, {"v2"}, quick_scorer_model);
  examples.SetCategoricalSet(2, feature_3, {"v3"}, quick_scorer_model);
  examples.SetCategoricalSet(3, feature_3, std::vector<std::string>{"v2", "v3"},
                             quick_scorer_model);
  examples.SetMissingCategoricalSet(4, feature_3, quick_scorer_model);

  examples.SetNumerical(0, feature_4, 0.00f, quick_scorer_model);
  examples.SetNumerical(1, feature_4, 0.05f, quick_scorer_model);
  examples.SetNumerical(2, feature_4, 0.10f, quick_scorer_model);
  examples.SetNumerical(3, feature_4, 0.20f, quick_scorer_model);
  examples.SetNumerical(4, feature_4, 0.30f, quick_scorer_model);

  std::vector<float> predictions;
  Predict(quick_scorer_model, examples, 5, &predictions);

  EXPECT_THAT(
      predictions,
      ElementsAre((1 + 4 + 40 + 400 + 1000 + 10000) * duplicate_factor,
                  (1 + 3 + 20 + 300 + 2000 + 10000) * duplicate_factor,
                  (1 + 3 + 20 + 400 + 3000 + 20000) * duplicate_factor,
                  (1 + 2 + 30 + 400 + 3000 + 20000) * duplicate_factor,
                  (1 + 1 + 10 + 300 + 2000 + 20000) * duplicate_factor));
}

TEST(QuickScorer, FinalizeConditionItems) {
  std::vector<internal::QuickScorerExtendedModel::ConditionItem> items{
      {/*.tree_idx =*/2, /*.leaf_mask =*/0b0111},
      {/*.tree_idx =*/1, /*.leaf_mask =*/0b1011},
      {/*.tree_idx =*/2, /*.leaf_mask =*/0b1101},
      {/*.tree_idx =*/1, /*.leaf_mask =*/0b1110},
  };
  internal::FinalizeConditionItems(&items);
  EXPECT_EQ(items.size(), 2);
  EXPECT_EQ(items[0].tree_idx, 1);
  EXPECT_EQ(items[1].tree_idx, 2);
  EXPECT_EQ(items[0].leaf_mask, 0b1010);
  EXPECT_EQ(items[1].leaf_mask, 0b0101);
}

TEST(QuickScorer, FinalizeIsHigherConditionItems) {
  std::vector<internal::QuickScorerExtendedModel::IsHigherConditionItem> items{
      {/*.threshold =*/1.f, /*.tree_idx =*/2, /*.leaf_mask =*/0b0111},
      {/*.threshold =*/3.f, /*.tree_idx =*/1, /*.leaf_mask =*/0b1011},
      {/*.threshold =*/1.f, /*.tree_idx =*/2, /*.leaf_mask =*/0b1101},
      {/*.threshold =*/2.f, /*.tree_idx =*/1, /*.leaf_mask =*/0b1110},
  };
  internal::FinalizeIsHigherConditionItems(&items);
  EXPECT_EQ(items.size(), 3);

  EXPECT_EQ(items[0].tree_idx, 2);
  EXPECT_EQ(items[0].leaf_mask, 0b0101);
  EXPECT_EQ(items[0].threshold, 1);

  EXPECT_EQ(items[1].tree_idx, 1);
  EXPECT_EQ(items[1].leaf_mask, 0b1110);
  EXPECT_EQ(items[1].threshold, 2);

  EXPECT_EQ(items[2].tree_idx, 1);
  EXPECT_EQ(items[2].leaf_mask, 0b1011);
  EXPECT_EQ(items[2].threshold, 3);
}

}  // namespace
}  // namespace decision_forest
}  // namespace serving
}  // namespace yggdrasil_decision_forests
