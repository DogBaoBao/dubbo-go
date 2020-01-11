/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"time"
)

import (
	"github.com/apache/dubbo-go/common/constant"
)

const (
	defaultMaxSubCategoryCount       = 20
	defaultGlobalInterval            = 60 * time.Second
	defaultMaxMetricCountPerRegistry = 5000
	defaultMaxCompassErrorCodeCount  = 100
	defaultMaxCompassAddonCount      = 20
)

type MetricConfig struct {
	/**
	 * the MetricManager's name. You can use 'default' to use the default implementation.
	 */
	Manager string `yaml:"manager" json:"manager,omitempty"`
	/**
	 * the max sub category count, it's same with com.alibaba.metrics.maxSubCategoryCount
	 */
	MaxSubCategoryCount int `yaml:"max_subcategory_count" json:"max_subcategory_count,omitempty"`

	/**
	 * the interval of collecting data, or report data, and so on...
	 * the unit is second
	 * see Interval
	 * default value is 60s
	 * it should >= 1s
	 */
	GlobalInterval time.Duration `yaml:"global_interval" json:"global_interval,omitempty"`

	/**
	 * MetricLevel -> interval
	 * we will use this map to find out the interval of the MetricLevel.
	 * it should >= 1s
	 */
	LevelInterval map[int]time.Duration `yaml:"level_interval" json:"level_interval,omitempty"`

	/**
	 * The max metric count per registry.
	 * the default value is 5000
	 * com.alibaba.metrics.maxMetricCountPerRegistry
	 */
	MaxMetricCountPerRegistry int `yaml:"max_metric_count_per_registry" json:"max_metric_count_per_registry,omitempty"`

	/**
	 * the max count of error code recorded by Compass. The default value is 100
	 */
	MaxCompassErrorCodeCount int `yaml:"max_compass_error_code_count" json:"max_compass_error_code_count,omitempty"`
	MaxCompassAddonCount     int `yaml:"max_metric_count_per_registry" json:"max_metric_count_per_registry,omitempty"`
}

func (mc *MetricConfig) GetMaxCompassAddonCount() int {
	if mc.MaxCompassAddonCount <= 0 {
		return defaultMaxCompassAddonCount
	}
	return mc.MaxCompassAddonCount
}

func (mc *MetricConfig) GetMaxCompassErrorCodeCount() int {
	if mc.MaxCompassErrorCodeCount <= 0 {
		return defaultMaxCompassErrorCodeCount
	}
	return mc.MaxCompassErrorCodeCount
}

func (mc *MetricConfig) GetMaxMetricCountPerRegistry() int {
	if mc.MaxMetricCountPerRegistry <= 0 {
		return defaultMaxMetricCountPerRegistry
	}
	return mc.MaxMetricCountPerRegistry
}

// if the user configures the value for this metric level and the value >= 1s, the configured value will be returned.s
func (mc *MetricConfig) GetLevelInterval(metricLevel int) time.Duration {
	if mc.LevelInterval == nil {
		return mc.GetGlobalInterval()
	}
	result, found := mc.LevelInterval[metricLevel]
	if found && result >= time.Second {
		return result
	}
	return mc.GetGlobalInterval()
}

func (mc *MetricConfig) GetGlobalInterval() time.Duration {
	if mc.GlobalInterval <= time.Second {
		return defaultGlobalInterval
	}
	return mc.GlobalInterval
}

func (mc *MetricConfig) GetMetricManagerName() string {
	if len(mc.Manager) <= 0 {
		return constant.DEFAULT_KEY
	}
	return mc.Manager
}

func (mc *MetricConfig) GetMaxSubCategoryCount() int {
	if mc.MaxSubCategoryCount <= 0 {
		return defaultMaxSubCategoryCount
	}
	return mc.MaxSubCategoryCount
}

/**
 * If the application is both consumer and provider, the provider's metric configuration will be used.
 * If and only if the application is just consumer, consumer's metric configuration wll be used.
 * Never return nil
 */
func GetMetricConfig() *MetricConfig {
	result := GetProviderConfig().MetricConfig
	if result == nil {
		result = GetConsumerConfig().MetricConfig
	}

	if result == nil {
		result = &MetricConfig{}
	}
	return result
}