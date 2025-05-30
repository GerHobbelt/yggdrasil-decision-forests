//
// Copyright 2022 Google LLC.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: yggdrasil_decision_forests/model/gradient_boosted_trees/gradient_boosted_trees.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Loss int32

const (
	// Selects the most adapted loss according to the nature of the task and the
	// statistics of the label.
	// - Binary classification -> BINOMIAL_LOG_LIKELIHOOD.
	Loss_DEFAULT Loss = 0
	// Binomial log likelihood. Only valid for binary classification.
	Loss_BINOMIAL_LOG_LIKELIHOOD Loss = 1
	// Least square loss. Only valid for regression.
	Loss_SQUARED_ERROR Loss = 2
	// Multinomial log likelihood i.e. cross-entropy.
	Loss_MULTINOMIAL_LOG_LIKELIHOOD Loss = 3
	// Deprecated: Use Loss_LAMBDA_MART_NDCG. LambdaMART with NDCG5
	Loss_LAMBDA_MART_NDCG5 Loss = 4
	// XE_NDCG_MART [arxiv.org/abs/1911.09798]
	Loss_XE_NDCG_MART Loss = 5
	// EXPERIMENTAl. Focal loss. Only valid for binary classification.
	// [https://arxiv.org/pdf/1708.02002.pdf]
	Loss_BINARY_FOCAL_LOSS Loss = 6
	// Poisson loss. Only valid for regression.
	Loss_POISSON Loss = 7
	// LambdaMART with NDCG@5.
	Loss_LAMBDA_MART_NDCG Loss = 8
)

// Enum value maps for Loss.
var (
	Loss_name = map[int32]string{
		0: "DEFAULT",
		1: "BINOMIAL_LOG_LIKELIHOOD",
		2: "SQUARED_ERROR",
		3: "MULTINOMIAL_LOG_LIKELIHOOD",
		4: "LAMBDA_MART_NDCG5",
		5: "XE_NDCG_MART",
		6: "BINARY_FOCAL_LOSS",
		7: "POISSON",
		8: "LAMBDA_MART_NDCG",
	}
	Loss_value = map[string]int32{
		"DEFAULT":                    0,
		"BINOMIAL_LOG_LIKELIHOOD":    1,
		"SQUARED_ERROR":              2,
		"MULTINOMIAL_LOG_LIKELIHOOD": 3,
		"LAMBDA_MART_NDCG5":          4,
		"XE_NDCG_MART":               5,
		"BINARY_FOCAL_LOSS":          6,
		"POISSON":                    7,
		"LAMBDA_MART_NDCG":           8,
	}
)

func (x Loss) Enum() *Loss {
	p := new(Loss)
	*p = x
	return p
}

func (x Loss) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Loss) Descriptor() protoreflect.EnumDescriptor {
	return file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_enumTypes[0].Descriptor()
}

func (Loss) Type() protoreflect.EnumType {
	return &file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_enumTypes[0]
}

func (x Loss) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Loss) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Loss(num)
	return nil
}

// Deprecated: Use Loss.Descriptor instead.
func (Loss) EnumDescriptor() ([]byte, []int) {
	return file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_rawDescGZIP(), []int{0}
}

// Header for the gradient boosted trees model.
type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of shards used to store the nodes.
	NumNodeShards *int32 `protobuf:"varint,1,opt,name=num_node_shards,json=numNodeShards" json:"num_node_shards,omitempty"`
	// Number of trees.
	NumTrees *int64 `protobuf:"varint,2,opt,name=num_trees,json=numTrees" json:"num_trees,omitempty"`
	// Loss used to train the model.
	Loss *Loss `protobuf:"varint,3,opt,name=loss,enum=yggdrasil_decision_forests.model.gradient_boosted_trees.proto.Loss" json:"loss,omitempty"`
	// Initial predictions of the model (before any tree is applied). The semantic
	// of the prediction depend on the "loss".
	InitialPredictions []float32 `protobuf:"fixed32,4,rep,name=initial_predictions,json=initialPredictions" json:"initial_predictions,omitempty"`
	// Number of trees extracted at each gradient boosting operation.
	NumTreesPerIter *int32 `protobuf:"varint,5,opt,name=num_trees_per_iter,json=numTreesPerIter,def=1" json:"num_trees_per_iter,omitempty"`
	// Loss evaluated on the validation dataset. Only available is a validation
	// dataset was provided during training.
	ValidationLoss *float32 `protobuf:"fixed32,6,opt,name=validation_loss,json=validationLoss" json:"validation_loss,omitempty"`
	// Container used to store the trees' nodes.
	NodeFormat *string `protobuf:"bytes,7,opt,name=node_format,json=nodeFormat,def=TFE_RECORDIO" json:"node_format,omitempty"`
	// Evaluation metrics and other meta-data computed during training.
	TrainingLogs *TrainingLogs `protobuf:"bytes,8,opt,name=training_logs,json=trainingLogs" json:"training_logs,omitempty"`
	// If true, call to predict methods return logits (e.g. instead of probability
	// in the case of classification).
	OutputLogits *bool `protobuf:"varint,9,opt,name=output_logits,json=outputLogits,def=0" json:"output_logits,omitempty"`
}

// Default values for Header fields.
const (
	Default_Header_NumTreesPerIter = int32(1)
	Default_Header_NodeFormat      = string("TFE_RECORDIO")
	Default_Header_OutputLogits    = bool(false)
)

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetNumNodeShards() int32 {
	if x != nil && x.NumNodeShards != nil {
		return *x.NumNodeShards
	}
	return 0
}

func (x *Header) GetNumTrees() int64 {
	if x != nil && x.NumTrees != nil {
		return *x.NumTrees
	}
	return 0
}

func (x *Header) GetLoss() Loss {
	if x != nil && x.Loss != nil {
		return *x.Loss
	}
	return Loss_DEFAULT
}

func (x *Header) GetInitialPredictions() []float32 {
	if x != nil {
		return x.InitialPredictions
	}
	return nil
}

func (x *Header) GetNumTreesPerIter() int32 {
	if x != nil && x.NumTreesPerIter != nil {
		return *x.NumTreesPerIter
	}
	return Default_Header_NumTreesPerIter
}

func (x *Header) GetValidationLoss() float32 {
	if x != nil && x.ValidationLoss != nil {
		return *x.ValidationLoss
	}
	return 0
}

func (x *Header) GetNodeFormat() string {
	if x != nil && x.NodeFormat != nil {
		return *x.NodeFormat
	}
	return Default_Header_NodeFormat
}

func (x *Header) GetTrainingLogs() *TrainingLogs {
	if x != nil {
		return x.TrainingLogs
	}
	return nil
}

func (x *Header) GetOutputLogits() bool {
	if x != nil && x.OutputLogits != nil {
		return *x.OutputLogits
	}
	return Default_Header_OutputLogits
}

// Log of the training. This proto is generated during the training of the
// model and optionally exported (as a plot) in the training logs directory.
type TrainingLogs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Measurements the model size and performances during the training.
	Entries []*TrainingLogs_Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
	// Names of the metrics stored in "secondary_metrics" field. The secondary
	// metrics depends on the task (e.g. classification) and is accessible with
	// "SecondaryMetricNames()". The i-th metric name of "secondary_metric_names"
	// correspond to the i-th metric value in "training_secondary_metrics" and
	// "validation_secondary_metrics".
	SecondaryMetricNames []string `protobuf:"bytes,2,rep,name=secondary_metric_names,json=secondaryMetricNames" json:"secondary_metric_names,omitempty"`
	// Number of trees in the final model. Without early stopping,
	// "number_of_trees_in_final_model" is equal to the "number_of_trees" of the
	// last "entries".
	NumberOfTreesInFinalModel *int32 `protobuf:"varint,3,opt,name=number_of_trees_in_final_model,json=numberOfTreesInFinalModel" json:"number_of_trees_in_final_model,omitempty"`
}

func (x *TrainingLogs) Reset() {
	*x = TrainingLogs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrainingLogs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrainingLogs) ProtoMessage() {}

func (x *TrainingLogs) ProtoReflect() protoreflect.Message {
	mi := &file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrainingLogs.ProtoReflect.Descriptor instead.
func (*TrainingLogs) Descriptor() ([]byte, []int) {
	return file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_rawDescGZIP(), []int{1}
}

func (x *TrainingLogs) GetEntries() []*TrainingLogs_Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *TrainingLogs) GetSecondaryMetricNames() []string {
	if x != nil {
		return x.SecondaryMetricNames
	}
	return nil
}

func (x *TrainingLogs) GetNumberOfTreesInFinalModel() int32 {
	if x != nil && x.NumberOfTreesInFinalModel != nil {
		return *x.NumberOfTreesInFinalModel
	}
	return 0
}

type TrainingLogs_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of trees. In the case of multi-dimensional gradients,
	// "number_of_trees" is the number of training step.
	NumberOfTrees *int32 `protobuf:"varint,1,opt,name=number_of_trees,json=numberOfTrees" json:"number_of_trees,omitempty"`
	// Performance of the model on the training dataset.
	TrainingLoss             *float32  `protobuf:"fixed32,2,opt,name=training_loss,json=trainingLoss" json:"training_loss,omitempty"`
	TrainingSecondaryMetrics []float32 `protobuf:"fixed32,3,rep,name=training_secondary_metrics,json=trainingSecondaryMetrics" json:"training_secondary_metrics,omitempty"`
	// Performance of the model on the validation dataset.
	ValidationLoss             *float32  `protobuf:"fixed32,4,opt,name=validation_loss,json=validationLoss" json:"validation_loss,omitempty"`
	ValidationSecondaryMetrics []float32 `protobuf:"fixed32,5,rep,name=validation_secondary_metrics,json=validationSecondaryMetrics" json:"validation_secondary_metrics,omitempty"`
	// Average of the absolute value of the new tree predictions estimated on
	// the training dataset. See Dart paper
	// (http://proceedings.mlr.press/v38/korlakaivinayak15.pdf) for details on
	// how to interpret it.
	MeanAbsPrediction *float64 `protobuf:"fixed64,6,opt,name=mean_abs_prediction,json=meanAbsPrediction" json:"mean_abs_prediction,omitempty"`
	// Sub-sampling factor applied during training on top of the "sampling"
	// hyper-parameter. Currently, the "subsample_factor" is only controlled by
	// the "adapt_subsample_for_maximum_training_duration" field i.e. the
	// "subsample_factor" factor (default to 1) is reduced so the training
	// finishes in "maximum_training_duration".
	SubsampleFactor *float32 `protobuf:"fixed32,7,opt,name=subsample_factor,json=subsampleFactor,def=1" json:"subsample_factor,omitempty"`
}

// Default values for TrainingLogs_Entry fields.
const (
	Default_TrainingLogs_Entry_SubsampleFactor = float32(1)
)

func (x *TrainingLogs_Entry) Reset() {
	*x = TrainingLogs_Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrainingLogs_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrainingLogs_Entry) ProtoMessage() {}

func (x *TrainingLogs_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrainingLogs_Entry.ProtoReflect.Descriptor instead.
func (*TrainingLogs_Entry) Descriptor() ([]byte, []int) {
	return file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_rawDescGZIP(), []int{1, 0}
}

func (x *TrainingLogs_Entry) GetNumberOfTrees() int32 {
	if x != nil && x.NumberOfTrees != nil {
		return *x.NumberOfTrees
	}
	return 0
}

func (x *TrainingLogs_Entry) GetTrainingLoss() float32 {
	if x != nil && x.TrainingLoss != nil {
		return *x.TrainingLoss
	}
	return 0
}

func (x *TrainingLogs_Entry) GetTrainingSecondaryMetrics() []float32 {
	if x != nil {
		return x.TrainingSecondaryMetrics
	}
	return nil
}

func (x *TrainingLogs_Entry) GetValidationLoss() float32 {
	if x != nil && x.ValidationLoss != nil {
		return *x.ValidationLoss
	}
	return 0
}

func (x *TrainingLogs_Entry) GetValidationSecondaryMetrics() []float32 {
	if x != nil {
		return x.ValidationSecondaryMetrics
	}
	return nil
}

func (x *TrainingLogs_Entry) GetMeanAbsPrediction() float64 {
	if x != nil && x.MeanAbsPrediction != nil {
		return *x.MeanAbsPrediction
	}
	return 0
}

func (x *TrainingLogs_Entry) GetSubsampleFactor() float32 {
	if x != nil && x.SubsampleFactor != nil {
		return *x.SubsampleFactor
	}
	return Default_TrainingLogs_Entry_SubsampleFactor
}

var File_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto protoreflect.FileDescriptor

var file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_rawDesc = []byte{
	0x0a, 0x54, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x5f, 0x64, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x73,
	0x74, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3d, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69,
	0x6c, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x73,
	0x74, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x03, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x68, 0x61,
	0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x4e, 0x6f,
	0x64, 0x65, 0x53, 0x68, 0x61, 0x72, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f,
	0x74, 0x72, 0x65, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x75, 0x6d,
	0x54, 0x72, 0x65, 0x65, 0x73, 0x12, 0x57, 0x0a, 0x04, 0x6c, 0x6f, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x43, 0x2e, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x5f,
	0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x73, 0x73, 0x52, 0x04, 0x6c, 0x6f, 0x73, 0x73, 0x12, 0x2f,
	0x0a, 0x13, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x02, 0x52, 0x12, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x2e, 0x0a, 0x12, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72,
	0x5f, 0x69, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x31, 0x52, 0x0f,
	0x6e, 0x75, 0x6d, 0x54, 0x72, 0x65, 0x65, 0x73, 0x50, 0x65, 0x72, 0x49, 0x74, 0x65, 0x72, 0x12,
	0x27, 0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f,
	0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x73, 0x73, 0x12, 0x2d, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x0c, 0x54,
	0x46, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x49, 0x4f, 0x52, 0x0a, 0x6e, 0x6f, 0x64,
	0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x70, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4b,
	0x2e, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61, 0x73, 0x69, 0x6c, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x73, 0x74,
	0x65, 0x64, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x0c, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x2a, 0x0a, 0x0d, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4c,
	0x6f, 0x67, 0x69, 0x74, 0x73, 0x22, 0xd2, 0x04, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x6b, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x79, 0x67, 0x67, 0x64, 0x72, 0x61,
	0x73, 0x69, 0x6c, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72,
	0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x72, 0x61, 0x64, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x65, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x4c, 0x6f, 0x67, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79,
	0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x14, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x1e, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x5f,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x19, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x54, 0x72, 0x65, 0x65, 0x73,
	0x49, 0x6e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0xdb, 0x02, 0x0a,
	0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x6f, 0x66, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x54, 0x72, 0x65, 0x65, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x6f, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x4c,
	0x6f, 0x73, 0x73, 0x12, 0x3c, 0x0a, 0x1a, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x02, 0x52, 0x18, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x12, 0x27, 0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6c, 0x6f, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x73, 0x73, 0x12, 0x40, 0x0a, 0x1c, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61,
	0x72, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x02,
	0x52, 0x1a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x2e, 0x0a, 0x13,
	0x6d, 0x65, 0x61, 0x6e, 0x5f, 0x61, 0x62, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x6d, 0x65, 0x61, 0x6e, 0x41,
	0x62, 0x73, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x10,
	0x73, 0x75, 0x62, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x3a, 0x01, 0x31, 0x52, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2a, 0xa3, 0x01, 0x0a, 0x04, 0x4c,
	0x6f, 0x73, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00,
	0x12, 0x1b, 0x0a, 0x17, 0x42, 0x49, 0x4e, 0x4f, 0x4d, 0x49, 0x41, 0x4c, 0x5f, 0x4c, 0x4f, 0x47,
	0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x4c, 0x49, 0x48, 0x4f, 0x4f, 0x44, 0x10, 0x01, 0x12, 0x11, 0x0a,
	0x0d, 0x53, 0x51, 0x55, 0x41, 0x52, 0x45, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02,
	0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x4e, 0x4f, 0x4d, 0x49, 0x41, 0x4c, 0x5f,
	0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x4c, 0x49, 0x48, 0x4f, 0x4f, 0x44, 0x10, 0x03,
	0x12, 0x15, 0x0a, 0x11, 0x4c, 0x41, 0x4d, 0x42, 0x44, 0x41, 0x5f, 0x4d, 0x41, 0x52, 0x54, 0x5f,
	0x4e, 0x44, 0x43, 0x47, 0x35, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x58, 0x45, 0x5f, 0x4e, 0x44,
	0x43, 0x47, 0x5f, 0x4d, 0x41, 0x52, 0x54, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x42, 0x49, 0x4e,
	0x41, 0x52, 0x59, 0x5f, 0x46, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x4c, 0x4f, 0x53, 0x53, 0x10, 0x06,
}

var (
	file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_rawDescOnce sync.Once
	file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_rawDescData = file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_rawDesc
)

func file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_rawDescGZIP() []byte {
	file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_rawDescOnce.Do(func() {
		file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_rawDescData = protoimpl.X.CompressGZIP(file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_rawDescData)
	})
	return file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_rawDescData
}

var file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_goTypes = []interface{}{
	(Loss)(0),                  // 0: yggdrasil_decision_forests.model.gradient_boosted_trees.proto.Loss
	(*Header)(nil),             // 1: yggdrasil_decision_forests.model.gradient_boosted_trees.proto.Header
	(*TrainingLogs)(nil),       // 2: yggdrasil_decision_forests.model.gradient_boosted_trees.proto.TrainingLogs
	(*TrainingLogs_Entry)(nil), // 3: yggdrasil_decision_forests.model.gradient_boosted_trees.proto.TrainingLogs.Entry
}
var file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_depIdxs = []int32{
	0, // 0: yggdrasil_decision_forests.model.gradient_boosted_trees.proto.Header.loss:type_name -> yggdrasil_decision_forests.model.gradient_boosted_trees.proto.Loss
	2, // 1: yggdrasil_decision_forests.model.gradient_boosted_trees.proto.Header.training_logs:type_name -> yggdrasil_decision_forests.model.gradient_boosted_trees.proto.TrainingLogs
	3, // 2: yggdrasil_decision_forests.model.gradient_boosted_trees.proto.TrainingLogs.entries:type_name -> yggdrasil_decision_forests.model.gradient_boosted_trees.proto.TrainingLogs.Entry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_init()
}
func file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_init() {
	if File_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrainingLogs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrainingLogs_Entry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_goTypes,
		DependencyIndexes: file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_depIdxs,
		EnumInfos:         file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_enumTypes,
		MessageInfos:      file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_msgTypes,
	}.Build()
	File_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto = out.File
	file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_rawDesc = nil
	file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_goTypes = nil
	file_yggdrasil_decision_forests_model_gradient_boosted_trees_gradient_boosted_trees_proto_depIdxs = nil
}
