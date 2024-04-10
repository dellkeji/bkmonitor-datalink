// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package metadata

import (
	"context"
	"fmt"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/unify-query/featureFlag"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/unify-query/trace"
)

// GetMustVmQueryFeatureFlag 判断该 TableID 是否强行指定为单指标单表
func GetMustVmQueryFeatureFlag(ctx context.Context, tableID string) bool {
	var (
		err  error
		user = GetUser(ctx)
	)

	ctx, span := trace.NewSpan(ctx, "check-must-vm-query-feature-flag")
	defer span.End(&err)

	span.Set("table-id", tableID)

	// 特性开关只有指定空间才启用 vm 查询
	ffUser := featureFlag.FFUser(span.TraceID(), map[string]interface{}{
		"name":     user.Name,
		"source":   user.Source,
		"spaceUid": user.SpaceUid,
		"tableID":  tableID,
	})

	status := featureFlag.BoolVariation(ctx, ffUser, "must-vm-query", false)
	span.Set("must-vm-query-flag", fmt.Sprintf("%v:%v", ffUser.GetCustom(), status))

	// 根据查询时间范围判断是否满足当前时间配置
	vmDataTime := featureFlag.IntVariation(ctx, ffUser, "range-vm-query", 0)

	if vmDataTime > 0 {
		queryParams := GetQueryParams(ctx)
		status = int64(vmDataTime) < queryParams.Start
		span.Set("vm-data-time", vmDataTime)
		span.Set("range-vm-query-flag", fmt.Sprintf("%v:%v", ffUser.GetCustom(), status))
	}

	return status
}

func GetVMQueryFeatureFlag(ctx context.Context) bool {
	var (
		err  error
		user = GetUser(ctx)
	)

	ctx, span := trace.NewSpan(ctx, "check-vm-query-feature-flag")
	defer span.End(&err)

	// 增加配置的特性开关
	if GetQueryRouter().CheckVmQuery(ctx, user.SpaceUid) {
		span.Set("vm-query-space-uid", "true")
		return true
	}

	// 特性开关只有指定空间才启用 vm 查询
	ffUser := featureFlag.FFUser(span.TraceID(), map[string]interface{}{
		"name":     user.Name,
		"source":   user.Source,
		"spaceUid": user.SpaceUid,
	})

	status := featureFlag.BoolVariation(ctx, ffUser, "vm-query", false)
	span.Set("vm-query-feature-flag", fmt.Sprintf("%v:%v", ffUser.GetCustom(), status))

	return status
}

func GetIsK8sFeatureFlag(ctx context.Context) bool {
	var (
		err  error
		user = GetUser(ctx)
	)

	ctx, span := trace.NewSpan(ctx, "check-is-k8s-feature-flag")
	defer span.End(&err)

	// 特性开关只有指定空间才启用 vm 查询
	ffUser := featureFlag.FFUser(span.TraceID(), map[string]interface{}{
		"name":     user.Name,
		"source":   user.Source,
		"spaceUid": user.SpaceUid,
	})

	status := featureFlag.BoolVariation(ctx, ffUser, "is-k8s", false)
	span.Set("is-k8s", fmt.Sprintf("%v:%v", ffUser.GetCustom(), status))

	return status
}
