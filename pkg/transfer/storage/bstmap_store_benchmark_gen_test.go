// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

//go:build BSTMap
// +build BSTMap

package storage_test

import "testing"

// BenchmarkStoreSet_BSTMap :
func BenchmarkStoreSet_BSTMap(b *testing.B) {
	withClosingStore(benchmarkStoreSet, b, newBSTMap())
}

// BenchmarkStoreUpdate_BSTMap :
func BenchmarkStoreUpdate_BSTMap(b *testing.B) {
	withClosingStore(benchmarkStoreUpdate, b, newBSTMap())
}

// BenchmarkStoreGet_BSTMap :
func BenchmarkStoreGet_BSTMap(b *testing.B) {
	withClosingStore(benchmarkStoreGet, b, newBSTMap())
}

// BenchmarkStoreGetHotPot_BSTMap :
func BenchmarkStoreGetHotPot_BSTMap(b *testing.B) {
	withClosingStore(benchmarkStoreGetHotPot, b, newBSTMap())
}

// benchmarkStoreExistsMissing_BSTMap :
func BenchmarkStoreExistsMissing_BSTMap(b *testing.B) {
	withClosingStore(benchmarkStoreExistsMissing, b, newBSTMap())
}

// BenchmarkStoreExists_BSTMap :
func BenchmarkStoreExists_BSTMap(b *testing.B) {
	withClosingStore(benchmarkStoreExists, b, newBSTMap())
}

// BenchmarkStoreDelete_BSTMap :
func BenchmarkStoreDelete_BSTMap(b *testing.B) {
	withClosingStore(benchmarkStoreDelete, b, newBSTMap())
}

// BenchmarkStoreScan_BSTMap :
func BenchmarkStoreScan_BSTMap(b *testing.B) {
	withClosingStore(benchmarkStoreScan, b, newBSTMap())
}

// BenchmarkStoreCommit_BSTMap :
func BenchmarkStoreCommit_BSTMap(b *testing.B) {
	withClosingStore(benchmarkStoreCommit, b, newBSTMap())
}
