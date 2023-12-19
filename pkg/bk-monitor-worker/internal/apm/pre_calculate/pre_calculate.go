// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package pre_calculate

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/bk-monitor-worker/config"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/bk-monitor-worker/internal/apm/pre_calculate/notifier"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/bk-monitor-worker/internal/apm/pre_calculate/storage"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/bk-monitor-worker/internal/apm/pre_calculate/window"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/utils/logger"
)

func Initial(parentCtx context.Context) (PreCalculateProcessor, error) {
	ctx, cancel := context.WithCancel(parentCtx)
	redisConfig := config.GlobalConfig.Store.RedisConfig

	// buffer size, default is 100000
	bufferSize, isExist := config.GlobalConfig.Task.ApmPreCalculate.Notifier["chanBufferSize"]
	if !isExist{
		bufferSize = 100000
	}

	windowConfig := config.GlobalConfig.Task.ApmPreCalculate.Window
	distributiveConfig :=windowConfig.Distributive

	// process config
	enabledTraceInfoCache, isExist := windowConfig.Processor["enabledTraceInfoCache"]
	if !isExist{
		enabledTraceInfoCache = 0
	}
	
	storageConfig := windowConfig.Storage
	autoClean, isExist:= storageConfig.Bloom.Normal["autoClean"]
	if isExist{
		autoClean = 24*time.Hour
	}

	overlapResetDuration, isExist := storageConfig.Bloom.NormalOverlap["resetDuration"]
	if isExist{
		overlapResetDuration = 2*time.Hour
	}

	bloomLayers, isExist := storageConfig.Bloom.LayersBloom["layers"]
	if isExist{
		bloomLayers = 5
	}

	cap, isExist := storageConfig.Bloom.DecreaseBloom["cap"]
	if isExist{
		cap = 100000000
	}

	decreaseLayers, isExist := storageConfig.Bloom.DecreaseBloom["layers"]
	if isExist{
		decreaseLayers = 10
	}
	divisor, isExist := storageConfig.Bloom.DecreaseBloom["divisor"]
	if isExist{
		divisor = 2
	}

	return NewPrecalculate().
		WithContext(ctx, cancel).
		WithNotifierConfig(
			notifier.BufferSize(bufferSize),
		).
		WithWindowRuntimeConfig(
			window.RuntimeConfigMaxSize(windowConfig.MaxSize),
			window.RuntimeConfigExpireInterval(windowConfig.ExpireInterval),
			window.RuntimeConfigMaxDuration(windowConfig.MaxDuration),
			window.ExpireIntervalIncrement(windowConfig.ExpireIntervalIncrement),
			window.NoDataMaxDuration(windowConfig.NoDataMaxDuration),
		).
		WithDistributiveWindowConfig(
			window.DistributiveWindowSubSize(distributiveConfig.SubSize),
			window.DistributiveWindowWatchExpiredInterval(distributiveConfig.WatchExpireInterval),
			window.ConcurrentProcessCount(distributiveConfig.ConcurrentCount),
			window.ConcurrentExpirationMaximum(distributiveConfig.ConcurrentExpirationMaximum),
		).
		WithProcessorConfig(
			window.EnabledTraceInfoCache(enabledTraceInfoCache != 0),
		).
		WithStorageConfig(
			storage.WorkerCount(storageConfig.WorkerCount),
			storage.SaveHoldMaxCount(storageConfig.SaveHoldMaxCount),
			storage.SaveHoldDuration(storageConfig.SaveHoldMaxDuration),
			storage.CacheBackend(storage.CacheTypeRedis),
			storage.CacheRedisConfig(
				storage.RedisCacheMode(redisConfig.Mode),
				storage.RedisCacheHost(redisConfig.Standalone.Host),
				storage.RedisCachePort(redisConfig.Standalone.Port),
				storage.RedisCacheSentinelAddress(redisConfig.Sentinel.Address...),
				storage.RedisCacheMasterName(redisConfig.Sentinel.MasterName),
				storage.RedisCacheSentinelPassword(redisConfig.Sentinel.Password),
				storage.RedisCachePassword(redisConfig.Standalone.Password),
				storage.RedisCacheDb(redisConfig.DB),
				storage.RedisCacheDialTimeout(redisConfig.DialTimeout),
				storage.RedisCacheReadTimeout(redisConfig.ReadTimeout),
			),
			storage.BloomConfig(
				storage.BloomFpRate(storageConfig.Bloom.FpRate),
				storage.NormalMemoryBloomConfig(
					storage.MemoryBloomAutoClean(autoClean),
				),
				storage.NormalOverlapMemoryBloomConfig(
					storage.OverlapBloomResetDuration(overlapResetDuration),
				),
				storage.LayerBloomConfig(storage.Layers(bloomLayers)),
				storage.LayerCapDecreaseBloomConfig(
					storage.CapDecreaseBloomCap(cap),
					storage.CapDecreaseBloomLayers(decreaseLayers),
					storage.CapDecreaseBloomDivisor(divisor),
				),
			),
			storage.SaveReqBufferSize(storageConfig.SaveRequestBufferSize),
		).
		WithMetricReport(
			EnabledMetricReport(config.MetricEnabled),
			MetricReportDataId(config.MetricReportDataId),
			MetricReportAccessToken(config.MetricReportAccessToken),
			MetricReportHost(config.MetricReportHost),
			ReportMetrics(
				SaveRequestChanCount,
				MessageReceiveChanCount,
				WindowMetric,
			),
			EnabledMetricReportInterval(config.MetricReportInterval),
			EnabledProfileReport(config.ProfileEnabled),
			ProfileAddress(config.ProfileHost),
			ProfileAppIdx(config.ProfileAppIdx),
		).
		Build(), nil
}

var apmLogger = logger.With(zap.String("package", "apm_precalculate"))
