// Copyright 2022 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !condmetric_kvm_profile
// +build !condmetric_kvm_profile

package metric

import (
	pb "gvisor.dev/gvisor/pkg/metric/metric_go_proto"
)

// KvmProfileUint64Metric is a metric type that is registered and used only when
// the "condmetric_kvm_profile" go tag is specified when building runsc.
//
// Otherwise it is exactly like a Uint64Metric.
type KvmProfileUint64Metric = FakeUint64Metric

// KvmProfileDistributionMetric is a metric type that is registered and used only
// when the "condmetric_kvm_profile" go tag is specified when building runsc.
//
// Otherwise it is exactly like a DistributionMetric.
type KvmProfileDistributionMetric = FakeDistributionMetric

// KvmProfileTimerMetric is a metric type that is registered and used only when
// the "condmetric_kvm_profile" go tag is specified when building runsc.
//
// Otherwise it is exactly like a TimerMetric.
type KvmProfileTimerMetric = FakeTimerMetric

// NewKvmProfileUint64Metric is equivalent to NewUint64Metric except it creates a
// KvmProfileUint64Metric
func NewKvmProfileUint64Metric(name string, sync bool, units pb.MetricMetadata_Units, description string, fields ...Field) (*KvmProfileUint64Metric, error) {
	return &FakeUint64Metric{}, nil
}

// MustCreateNewKvmProfileUint64Metric is equivalent to MustCreateNewUint64Metric
// except it creates a KvmProfileUint64Metric.
func MustCreateNewKvmProfileUint64Metric(name string, sync bool, description string, fields ...Field) *KvmProfileUint64Metric {
	return &FakeUint64Metric{}
}

// NewKvmProfileDistributionMetric is equivalent to NewDistributionMetric except
// it creates a KvmProfileDistributionMetric.
func NewKvmProfileDistributionMetric(name string, sync bool, bucketer Bucketer, unit pb.MetricMetadata_Units, description string, fields ...Field) (*KvmProfileDistributionMetric, error) {
	return &FakeDistributionMetric{}, nil
}

// MustCreateNewKvmProfileDistributionMetric is equivalent to
// MustCreateNewDistributionMetric except it creates a
// KvmProfileDistributionMetric.
func MustCreateNewKvmProfileDistributionMetric(name string, sync bool, bucketer Bucketer, unit pb.MetricMetadata_Units, description string, fields ...Field) *KvmProfileDistributionMetric {
	return &FakeDistributionMetric{}
}

// NewKvmProfileTimerMetric is equivalent to NewTimerMetric except it creates a
// KvmProfileTimerMetric.
func NewKvmProfileTimerMetric(name string, nanoBucketer Bucketer, description string, fields ...Field) (*KvmProfileTimerMetric, error) {
	return &FakeTimerMetric{}, nil
}

// MustCreateNewKvmProfileTimerMetric is equivalent to MustCreateNewTimerMetric
// except it creates a KvmProfileTimerMetric.
func MustCreateNewKvmProfileTimerMetric(name string, nanoBucketer Bucketer, description string, fields ...Field) *KvmProfileTimerMetric {
	return &FakeTimerMetric{}
}
